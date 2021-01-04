package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type TextHandler struct {
	BaseHandler
	Timestamp     string
	Sort          bool
	TruncateLevel bool `toml:"truncate_level"`
}

func NewTextHandler() TextHandler {
	return TextHandler{
		BaseHandler:   NewBaseHandler(),
		Timestamp:     "YYYY-MM-dd HH:mm:ss",
		Sort:          true,
		TruncateLevel: false,
	}
}

func (config TextHandler) Parse() (*Handler, error) {
	f := logrus.TextFormatter{
		ForceColors:            true,
		DisableSorting:         !config.Sort,
		FullTimestamp:          true,
		TimestampFormat:        convertDateFormat(config.Timestamp),
		DisableLevelTruncation: !config.TruncateLevel,
		PadLevelText:           !config.TruncateLevel,
	}

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.formatter = &f
	return h, nil
}

func init() {
	RegisterParser("text", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (*Handler, error) {
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
	})
}
