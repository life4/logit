package logit

import (
	"net"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

type LogstashHandler struct {
	BaseHandler
	Network string
	Address string
	Version string
	Type    string
}

func NewLogstashHandler() LogstashHandler {
	return LogstashHandler{
		BaseHandler: NewBaseHandler(),
		Version:     "1",
		Type:        "logit",
	}
}

func (config LogstashHandler) Parse() (*Handler, error) {
	conn, err := net.Dial(config.Network, config.Address)
	if err != nil {
		return nil, err
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{
		"@version": config.Version,
		"type":     config.Type,
	}))

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
