// +build h_fluentd h_all

package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/evalphobia/logrus_fluent"
)

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

func init() {
	RegisterParser("fluentd", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (*Handler, error) {
		fconf := NewFluentdHandler()
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
