package logit

import "github.com/evalphobia/logrus_fluent"

type FluentdHandler struct {
	BaseHandler
	Host     string
	Port     int
	MaxRetry int
}

func NewFluentdHandler() FluentdHandler {
	return FluentdHandler{
		BaseHandler: NewBaseHandler(),
		Host:        "localhost",
		Port:        24224,
	}
}

func (config FluentdHandler) Parse() (*Handler, error) {
	hook, err := logrus_fluent.NewWithConfig(logrus_fluent.Config{
		Host:     config.Host,
		Port:     config.Port,
		MaxRetry: config.MaxRetry,
	})
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
