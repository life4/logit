// +build h_logstash h_all

package logit

import (
	"fmt"
	"net"

	"github.com/BurntSushi/toml"
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

func init() {
	RegisterParser("logstash", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (*Handler, error) {
		fconf := NewLogstashHandler()
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
