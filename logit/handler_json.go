package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type JSONHandler struct {
	BaseHandler
	DataKey   string `toml:"data_key"`
	Timestamp string
}

func NewJSONHandler() JSONHandler {
	return JSONHandler{
		DataKey:     "",
		Timestamp:   "YYYY-MM-dd HH:mm:ss",
		BaseHandler: NewBaseHandler(),
	}
}

func (config JSONHandler) Parse() (Handler, error) {
	f := logrus.JSONFormatter{
		DataKey:         config.DataKey,
		TimestampFormat: convertDateFormat(config.Timestamp),
	}
	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.SetFormatter(&f)
	return h, nil
}

func init() {
	RegisterParser("json", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (Handler, error) {
		fconf := NewJSONHandler()
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
