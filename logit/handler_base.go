package logit

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type BaseHandler struct {
	LevelFrom string `toml:"level_from"`
	LevelTo   string `toml:"level_to"`
	File      string
	Mode      os.FileMode
	Workers   int
}

func NewBaseHandler() BaseHandler {
	return BaseHandler{
		LevelFrom: "trace",
		LevelTo:   "panic",
		File:      "stdout",
		Workers:   1,
	}
}

func (config BaseHandler) Parse() (Handler, error) {
	lfrom, err := logrus.ParseLevel(config.LevelFrom)
	if err != nil {
		return nil, err
	}
	lto, err := logrus.ParseLevel(config.LevelTo)
	if err != nil {
		return nil, err
	}

	var stream io.Writer
	switch strings.ToLower(config.File) {
	case "stdout":
		stream = os.Stdout
	case "stderr":
		stream = os.Stderr
	default:
		stream, err = os.OpenFile(
			config.File,
			os.O_WRONLY|os.O_APPEND|os.O_CREATE,
			config.Mode,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot open file: %v", err)
		}
	}

	hSync := HandlerSync{
		stream:    stream,
		levelFrom: lfrom,
		levelTo:   lto,
	}

	if config.Workers > 1 {
		hAsync := HandlerAsync{
			workers: config.Workers,
			handler: hSync,
		}
		return &hAsync, nil
	}

	return &hSync, nil
}
