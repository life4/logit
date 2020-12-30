package logit

import (
	"os"

	"github.com/sirupsen/logrus"
)

type JSONHandler struct {
	DataKey   string
	Timestamp string
	LevelFrom Level `toml:"level_from"`
	LevelTo   Level `toml:"level_to"`
}

func NewJSONHandler() JSONHandler {
	return JSONHandler{
		DataKey:   "",
		Timestamp: "YYYY-MM-dd HH:mm:ss",
		LevelFrom: TraceLevel,
		LevelTo:   PanicLevel,
	}
}

func (config JSONHandler) Parse() (*Handler, error) {
	f := logrus.JSONFormatter{
		DataKey:         config.DataKey,
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
