// +build h_gcloud !h_clean

package logit

import (
	"context"
	"fmt"

	"cloud.google.com/go/logging"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

var gCloudLevels = map[logrus.Level]logging.Severity{
	logrus.TraceLevel: logging.Debug,
	logrus.DebugLevel: logging.Debug,
	logrus.InfoLevel:  logging.Info,
	logrus.WarnLevel:  logging.Warning,
	logrus.ErrorLevel: logging.Error,
	logrus.FatalLevel: logging.Critical,
	logrus.PanicLevel: logging.Alert,
}

type GCloudHook struct {
	logger *logging.Logger
	labels map[string]string
}

func (*GCloudHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (h *GCloudHook) Fire(entry *logrus.Entry) error {
	ctx := context.Background()
	return h.logger.LogSync(ctx, logging.Entry{
		Timestamp: entry.Time,
		Severity:  gCloudLevels[entry.Level],
		Labels:    h.labels,
		Payload: map[string]interface{}{
			"message": entry.Message,
			"data":    entry.Data,
		},
	})
}

type GCloudHandler struct {
	BaseHandler
	Credentials string
	Endpoint    string
	Labels      map[string]string
	LogName     string `toml:"log_name"`
	ProjectId   string `toml:"project_id"`
}

func NewGCloudHandler() GCloudHandler {
	return GCloudHandler{
		BaseHandler: NewBaseHandler(),
		LogName:     "logit",
	}
}

func (config GCloudHandler) Parse() (Handler, error) {
	opts := make([]option.ClientOption, 0)
	if config.Credentials != "" {
		opts = append(opts, option.WithCredentialsFile(config.Credentials))
	}
	if config.Endpoint != "" {
		opts = append(opts, option.WithEndpoint(config.Endpoint))
	}

	ctx := context.Background()
	client, err := logging.NewClient(ctx, config.ProjectId, opts...)
	if err != nil {
		return nil, err
	}
	hook := &GCloudHook{
		logger: client.Logger(config.LogName),
		labels: config.Labels,
	}

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.SetHook(hook)
	return h, nil
}

func init() {
	RegisterParser("gcloud", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (Handler, error) {
		fconf := NewGCloudHandler()
		err := meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("parse: %v", err)
		}
		handler, err := fconf.Parse()
		if err != nil {
			return nil, fmt.Errorf("init: %v", err)
		}
		return handler, nil
	})
}
