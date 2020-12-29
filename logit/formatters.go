package logit

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vjeantet/jodaTime"
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
	date := time.Date(2006, time.January, 2, 15, 4, 5, 999999999, time.UTC)
	tformat := jodaTime.Format(config.Timestamp, date)

	f := logrus.TextFormatter{
		ForceColors:     config.Colors,
		DisableColors:   !config.Colors,
		DisableSorting:  !config.Sort,
		FullTimestamp:   true,
		TimestampFormat: tformat,
	}
	return &f, nil
}
