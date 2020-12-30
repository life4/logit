package logit

import (
	"github.com/sirupsen/logrus"
)

type LogFmtHandler struct {
	BaseHandler
	Timestamp string
	Sort      bool
}

func NewLogFmtHandler() LogFmtHandler {
	return LogFmtHandler{
		Timestamp:   "YYYY-MM-dd HH:mm:ss",
		Sort:        true,
		BaseHandler: NewBaseHandler(),
	}
}

func (config LogFmtHandler) Parse() (*Handler, error) {
	f := logrus.TextFormatter{
		DisableColors:   true,
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
