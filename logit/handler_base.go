package logit

import (
	"os"

	"github.com/sirupsen/logrus"
)

type BaseHandler struct {
	LevelFrom string `toml:"level_from"`
	LevelTo   string `toml:"level_to"`
}

func NewBaseHandler() BaseHandler {
	return BaseHandler{
		LevelFrom: "trace",
		LevelTo:   "panic",
	}
}

func (config BaseHandler) Parse() (*Handler, error) {
	lfrom, err := logrus.ParseLevel(config.LevelFrom)
	if err != nil {
		return nil, err
	}
	lto, err := logrus.ParseLevel(config.LevelTo)
	if err != nil {
		return nil, err
	}
	h := Handler{
		stream:    os.Stdout,
		levelFrom: lfrom,
		levelTo:   lto,
	}
	return &h, nil
}
