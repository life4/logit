// +build h_rolling !h_clean

package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

type RollingHandler struct {
	BaseHandler
	SubHandler toml.Primitive `toml:"handler"`

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

func (config RollingHandler) Parse() (Handler, error) {
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
	h.SetFormatter(&logrus.TextFormatter{})
	h.SetStream(writer)
	return h, nil
}

func init() {
	RegisterParser("rolling", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (Handler, error) {
		fconf := NewRollingHandler()
		err := meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("parse: %v", err)
		}
		handler, err := fconf.Parse()
		if err != nil {
			return nil, fmt.Errorf("init: %v", err)
		}
		subhandler, err := ParseHandler(meta, fconf.SubHandler)
		if err != nil {
			return nil, err
		}

		switch h := subhandler.(type) {
		case *HandlerAsync:
			handler.SetFormatter(h.handler.formatter)
		case *HandlerSync:
			handler.SetFormatter(h.formatter)
		default:
			panic("unreachable")
		}

		return handler, nil
	})
}
