package logit

import "github.com/weekface/mgorus"

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
