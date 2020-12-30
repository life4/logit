package logit

import (
	"github.com/sirupsen/logrus"
)

type TextHandler struct {
	BaseHandler
	Colors    bool
	Timestamp string
	Sort      bool
}

func NewTextHandler() TextHandler {
	return TextHandler{
		Colors:      true,
		Timestamp:   "YYYY-MM-dd HH:mm:ss",
		Sort:        true,
		BaseHandler: NewBaseHandler(),
	}
}

func (config TextHandler) Parse() (*Handler, error) {
	f := logrus.TextFormatter{
		ForceColors:     config.Colors,
		DisableColors:   !config.Colors,
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
