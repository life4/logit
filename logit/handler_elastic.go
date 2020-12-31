package logit

import (
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
		logrus.DebugLevel,
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
