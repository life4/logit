package logit

import (
	"github.com/kenshaw/sdhook"
)

type GCloudHandler struct {
	BaseHandler
	Credentials string
	Service     string
	LogName     string `toml:"app_name"`
	ProjectId   string `toml:"project_id"`
}

func NewGCloudHandler() GCloudHandler {
	return GCloudHandler{
		BaseHandler: NewBaseHandler(),
	}
}

func (config GCloudHandler) Parse() (*Handler, error) {
	options := []sdhook.Option{
		sdhook.GoogleServiceAccountCredentialsFile(config.Credentials),
	}
	if config.Service != "" {
		options = append(options, sdhook.ErrorReportingService(config.Service))
	}
	if config.LogName != "" {
		options = append(options, sdhook.ErrorReportingLogName(config.LogName))
	}
	if config.ProjectId != "" {
		options = append(options, sdhook.ProjectID(config.ProjectId))
	}

	hook, err := sdhook.New(options...)
	if err != nil {
		return nil, err
	}

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.hook = hook
	return h, nil
}
