package logit

import (
	"github.com/sirupsen/logrus"
)

type FJSON struct {
	DataKey   string
	Timestamp string
}

func NewFJSON() FJSON {
	return FJSON{
		DataKey:   "",
		Timestamp: "YYYY-MM-dd HH:mm:ss",
	}
}

func FJSONParse(config FJSON) (*logrus.JSONFormatter, error) {
	f := logrus.JSONFormatter{
		DataKey:         config.DataKey,
		TimestampFormat: convertDateFormat(config.Timestamp),
	}
	return &f, nil
}
