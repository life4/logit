package logit

import (
	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
)

type SentryHandler struct {
	BaseHandler
	DSN string
}

func NewSentryHandler() SentryHandler {
	return SentryHandler{
		BaseHandler: NewBaseHandler(),
	}
}

func (config SentryHandler) Parse() (*Handler, error) {
	allLevels := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
	hook, err := logrus_sentry.NewSentryHook(config.DSN, allLevels)
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
