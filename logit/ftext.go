package logit

import (
	"github.com/sirupsen/logrus"
)

type FText struct {
	Colors    bool
	Timestamp string
	Sort      bool
}

func NewFText() FText {
	return FText{
		Colors:    true,
		Timestamp: "YYYY-MM-dd HH:mm:ss",
		Sort:      true,
	}
}

func FTextParse(config FText) (*logrus.TextFormatter, error) {
	f := logrus.TextFormatter{
		ForceColors:     config.Colors,
		DisableColors:   !config.Colors,
		DisableSorting:  !config.Sort,
		FullTimestamp:   true,
		TimestampFormat: convertDateFormat(config.Timestamp),
	}
	return &f, nil
}
