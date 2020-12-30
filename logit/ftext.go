package logit

import (
	"os"

	"github.com/sirupsen/logrus"
)

type TextHandler struct {
	Colors    bool
	Timestamp string
	Sort      bool
	LevelFrom Level `toml:"level_from"`
	LevelTo   Level `toml:"level_to"`
}

func NewTextHandler() TextHandler {
	return TextHandler{
		Colors:    true,
		Timestamp: "YYYY-MM-dd HH:mm:ss",
		Sort:      true,
		LevelFrom: TraceLevel,
		LevelTo:   PanicLevel,
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
	h := Handler{
		formatter: &f,
		stream:    os.Stdout,
		levelFrom: config.LevelFrom.Level,
		levelTo:   config.LevelTo.Level,
	}
	return &h, nil
}
