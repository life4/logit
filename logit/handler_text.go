package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type TextHandler struct {
	BaseHandler
	Timestamp string
	Sort      bool
}

func NewTextHandler() TextHandler {
	return TextHandler{
		Timestamp:   "YYYY-MM-dd HH:mm:ss",
		Sort:        true,
		BaseHandler: NewBaseHandler(),
	}
}

func (config TextHandler) Parse() (*Handler, error) {
	f := logrus.TextFormatter{
		ForceColors:     true,
		DisableSorting:  !config.Sort,
		FullTimestamp:   true,
		TimestampFormat: convertDateFormat(config.Timestamp),
	}

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.formatter = &f
	return h, nil
}

func textParser(meta toml.MetaData, primitive toml.Primitive) (*Handler, error) {
	fconf := NewTextHandler()
	err := meta.PrimitiveDecode(primitive, &fconf)
	if err != nil {
		return nil, fmt.Errorf("parse: %v", err)
	}
	handler, err := fconf.Parse()
	if err != nil {
		return nil, fmt.Errorf("init: %v", err)
	}
	return handler, nil
}

func init() {
	RegisterParser("text", textParser)
}
