// +build h_loggly h_all

package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
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

func init() {
	RegisterParser("loggly", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (*Handler, error) {
		fconf := NewLogglyHandler()
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
