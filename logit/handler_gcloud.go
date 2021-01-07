// +build h_gcloud !h_clean

package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/kenshaw/sdhook"
)

type GCloudHandler struct {
	BaseHandler
	Credentials string
	Service     string
	LogName     string `toml:"log_name"`
	ProjectId   string `toml:"project_id"`
}

func NewGCloudHandler() GCloudHandler {
	return GCloudHandler{
		BaseHandler: NewBaseHandler(),
	}
}

func (config GCloudHandler) Parse() (Handler, error) {
	options := []sdhook.Option{
		sdhook.GoogleServiceAccountCredentialsFile(config.Credentials),
	}
	if config.Service != "" {
		options = append(options, sdhook.ErrorReportingService(config.Service))
	}
	if config.LogName != "" {
		options = append(options, sdhook.LogName(config.LogName))
	}
	if config.ProjectId != "" {
		options = append(options, sdhook.ProjectID(config.ProjectId))
	}

	hook, err := sdhook.New(options...)
	if err != nil {
		return nil, fmt.Errorf("cannot create gcloud hook: %v", err)
	}

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.SetHook(hook)

	// TODO: sdhook is always async. Can we force it to be sync?
	// Contribute? The project seems dead.
	switch handler := h.(type) {
	case *HandlerAsync:
		handler.handler.wait = hook.Wait
		return &handler.handler, nil
	case *HandlerSync:
		handler.wait = hook.Wait
		return handler, nil
	}
	panic("unreachable")
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
