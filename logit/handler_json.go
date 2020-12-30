package logit

import (
	"github.com/sirupsen/logrus"
)

type JSONHandler struct {
	BaseHandler
	DataKey   string
	Timestamp string
}

func NewJSONHandler() JSONHandler {
	return JSONHandler{
		DataKey:     "",
		Timestamp:   "YYYY-MM-dd HH:mm:ss",
		BaseHandler: NewBaseHandler(),
	}
}

func (config JSONHandler) Parse() (*Handler, error) {
	f := logrus.JSONFormatter{
		DataKey:         config.DataKey,
		TimestampFormat: convertDateFormat(config.Timestamp),
	}
	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.formatter = &f
	return h, nil
}
