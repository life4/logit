// +build h_elastic h_all

package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

type ElasticHandler struct {
	BaseHandler
	URLs  []string
	Host  string
	Index string
}

func NewElasticHandler() ElasticHandler {
	return ElasticHandler{
		BaseHandler: NewBaseHandler(),
		URLs:        []string{"http://localhost:9200"},
		Host:        "localhost",
		Index:       "logit",
	}
}

func (config ElasticHandler) Parse() (*Handler, error) {
	client, err := elastic.NewClient(elastic.SetURL(config.URLs...))
	if err != nil {
		return nil, err
	}
	hook, err := elogrus.NewElasticHook(
		client,
		config.Host,
		logrus.TraceLevel,
		config.Index,
	)
	if err != nil {
		return nil, err
	}

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
	RegisterParser("elastic", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (*Handler, error) {
		fconf := NewElasticHandler()
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
