package logit

import (
	"github.com/sebest/logrusly"
	"github.com/sirupsen/logrus"
)

type LogglyHandler struct {
	BaseHandler
	Token string
	Host  string
	Tags  []string
}

func NewLogglyHandler() LogglyHandler {
	return LogglyHandler{
		BaseHandler: NewBaseHandler(),
		Tags:        []string{},
	}
}

func (config LogglyHandler) Parse() (*Handler, error) {
	hook := logrusly.NewLogglyHook(
		config.Token,
		config.Host,
		logrus.TraceLevel,
		config.Tags...,
	)

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.hook = hook
	return h, nil
}
