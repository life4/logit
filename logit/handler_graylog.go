// +build h_graylog h_all

package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	graylog "github.com/gemnasium/logrus-graylog-hook"
)

type GraylogHandler struct {
	BaseHandler
	Address string
}

func NewGraylogHandler() GraylogHandler {
	return GraylogHandler{
		BaseHandler: NewBaseHandler(),
	}
}

func (config GraylogHandler) Parse() (*Handler, error) {
	hook := graylog.NewGraylogHook(config.Address, map[string]interface{}{})
	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.hook = hook
	return h, nil
}

func init() {
	RegisterParser("graylog", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (*Handler, error) {
		fconf := NewGraylogHandler()
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
