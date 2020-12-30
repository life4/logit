package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type CMain struct {
	DefaultLevel string `toml:"default_level"`
}

type CHandler struct {
	Format string
}

type CFields struct {
	Message string
	Level   string
	Time    string
}

// RawConfig represents the struct to parse TOML config into.
type RawConfig struct {
	Main        CMain
	Fields      CFields
	HandlersRaw []toml.Primitive `toml:"handler"`
}

// RawConfig represents the config how it is used internally.
// It is generated from RawConfig.
type Config struct {
	DefaultLevel logrus.Level
	Handlers     []Handler
	Fields       CFields
}

func ReadConfig(cpath string) (*Config, error) {
	raw := RawConfig{
		Main: CMain{
			DefaultLevel: "ERROR",
		},
		Fields: CFields{
			Message: "msg",
			Level:   "level",
			Time:    "time",
		},
	}

	meta, err := toml.DecodeFile(cpath, &raw)
	if err != nil {
		return nil, fmt.Errorf("cannot read config: %v", err)
	}

	config := Config{
		Handlers: make([]Handler, len(raw.HandlersRaw)),
	}
	for i, primitive := range raw.HandlersRaw {
		h, err := parseFormatter(meta, primitive)
		if err != nil {
			return nil, fmt.Errorf("cannot parse handler: %v", err)
		}
		config.Handlers[i] = *h
	}

	undecoded := meta.Undecoded()
	if len(undecoded) != 0 {
		return nil, fmt.Errorf("unknown fields: %v", undecoded)
	}

	config.DefaultLevel, err = logrus.ParseLevel(raw.Main.DefaultLevel)
	if err != nil {
		return nil, fmt.Errorf("cannot parse default_level: %v", err)
	}

	config.Fields = raw.Fields
	return &config, nil
}

func parseFormatter(meta toml.MetaData, primitive toml.Primitive) (*Handler, error) {
	var config CHandler
	err := meta.PrimitiveDecode(primitive, &config)
	if err != nil {
		return nil, err
	}

	switch config.Format {
	case "text":
		fconf := NewTextHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, err
		}
		return fconf.Parse()
	case "logfmt":
		fconf := NewLogFmtHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, err
		}
		return fconf.Parse()
	case "json":
		fconf := NewJSONHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, err
		}
		return fconf.Parse()
	default:
		return nil, fmt.Errorf("unknown format: %s", config.Format)
	}
}
