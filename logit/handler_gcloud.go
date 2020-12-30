package logit

import (
	"fmt"

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

func (config GCloudHandler) Parse() (*Handler, error) {
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
	h.hook = hook
	h.wait = hook.Wait
	return h, nil
}
