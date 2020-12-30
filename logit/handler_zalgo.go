package logit

import "github.com/aybabtme/logzalgo"

type ZalgoHandler struct {
	BaseHandler
}

func NewZalgoHandler() ZalgoHandler {
	return ZalgoHandler{
		BaseHandler: NewBaseHandler(),
	}
}

func (config ZalgoHandler) Parse() (*Handler, error) {
	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.formatter = logzalgo.NewZalgoFormatterrrrrr()
	return h, nil
}
