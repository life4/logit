package logit

import (
	"log/syslog"

	syslogrus "github.com/sirupsen/logrus/hooks/syslog"
)

type SysLogHandler struct {
	BaseHandler
	Network string
	Address string
	Tag     string
}

func NewSysLogHandler() SysLogHandler {
	return SysLogHandler{
		BaseHandler: NewBaseHandler(),
		Tag:         "logit",
	}
}

func (config SysLogHandler) Parse() (*Handler, error) {
	hook, err := syslogrus.NewSyslogHook(
		config.Network,
		config.Address,
		syslog.LOG_INFO,
		config.Tag,
	)
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
