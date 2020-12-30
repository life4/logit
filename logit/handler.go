package logit

import (
	"github.com/sirupsen/logrus"
)

type Handler struct {
	formatter logrus.Formatter
	levelFrom logrus.Level
	levelTo   logrus.Level
}

func (h Handler) Format(e *logrus.Entry) ([]byte, error) {
	if e.Level < h.levelTo {
		return []byte{}, nil
	}
	if e.Level > h.levelFrom {
		return []byte{}, nil
	}
	return h.formatter.Format(e)
}
