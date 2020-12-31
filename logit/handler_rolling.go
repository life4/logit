package logit

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

type RollingHandler struct {
	BaseHandler

	File       string `toml:"file"`
	MaxSize    int    `toml:"max_size"`
	MaxAge     int    `toml:"max_age"`
	MaxBackups int    `toml:"max_backups"`
	LocalTime  bool   `toml:"local_time"`
	Compress   bool   `toml:"compress"`
}

func NewRollingHandler() RollingHandler {
	return RollingHandler{
		BaseHandler: NewBaseHandler(),
		MaxSize:     500,
		MaxBackups:  3,
		MaxAge:      28,
		Compress:    false,
	}
}

func (config RollingHandler) Parse() (*Handler, error) {
	writer := &lumberjack.Logger{
		Filename:   config.File,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
		Compress:   config.Compress,
	}

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.formatter = &logrus.TextFormatter{}
	h.stream = writer
	return h, nil
}
