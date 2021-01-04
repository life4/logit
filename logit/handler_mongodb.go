package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/weekface/mgorus"
)

type MongoDBHandler struct {
	BaseHandler
	URL        string
	DB         string
	Collection string
}

func NewMongoDBHandler() MongoDBHandler {
	return MongoDBHandler{
		BaseHandler: NewBaseHandler(),
		URL:         "localhost",
	}
}

func (config MongoDBHandler) Parse() (*Handler, error) {
	hook, err := mgorus.NewHooker(config.URL, config.DB, config.Collection)
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
	RegisterParser("mongodb", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (*Handler, error) {
		fconf := NewMongoDBHandler()
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
