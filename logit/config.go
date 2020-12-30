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
	Formatter string
}

// RawConfig represents the struct to parse TOML config into.
type RawConfig struct {
	Main        CMain
	HandlersRaw []toml.Primitive `toml:"handler"`
}

// RawConfig represents the config how it is used internally.
// It is generated from RawConfig.
type Config struct {
	DefaultLevel logrus.Level
	Formatters   []logrus.Formatter
}

func ReadConfig(cpath string) (*Config, error) {
	raw := RawConfig{
		Main: CMain{
			DefaultLevel: "ERROR",
		},
	}

	meta, err := toml.DecodeFile(cpath, &raw)
	if err != nil {
		return nil, fmt.Errorf("cannot read config: %v", err)
	}

	config := Config{
		Formatters: make([]logrus.Formatter, len(raw.HandlersRaw)),
	}
	for i, primitive := range raw.HandlersRaw {
		f, err := parseFormatter(meta, primitive)
		if err != nil {
			return nil, fmt.Errorf("cannot parse handler: %v", err)
		}
		config.Formatters[i] = f
	}

	undecoded := meta.Undecoded()
	if len(undecoded) != 0 {
		return nil, fmt.Errorf("unknown fields: %v", undecoded)
	}

	config.DefaultLevel, err = logrus.ParseLevel(raw.Main.DefaultLevel)
	if err != nil {
		return nil, fmt.Errorf("cannot parse default_level: %v", err)
	}

	return &config, nil
}

func parseFormatter(meta toml.MetaData, primitive toml.Primitive) (logrus.Formatter, error) {
	var h CHandler
	err := meta.PrimitiveDecode(primitive, &h)
	if err != nil {
		return nil, err
	}

	switch h.Formatter {
	case "text":
		fconf := NewFText()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, err
		}
		return FTextParse(fconf)
	case "json":
		fconf := NewFJSON()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, err
		}
		return FJSONParse(fconf)
	default:
		return nil, fmt.Errorf("unknown formatter: %s", h.Formatter)
	}
}
