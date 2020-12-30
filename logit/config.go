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
	LevelFrom string `toml:"level_from"`
	LevelTo   string `toml:"level_to"`
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
	Handlers     []Handler
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

	return &config, nil
}

func parseFormatter(meta toml.MetaData, primitive toml.Primitive) (*Handler, error) {
	var config CHandler
	err := meta.PrimitiveDecode(primitive, &config)
	if err != nil {
		return nil, err
	}

	var f logrus.Formatter
	switch config.Formatter {
	case "text":
		fconf := NewFText()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, err
		}
		f, err = fconf.Parse()
	case "json":
		fconf := NewFJSON()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, err
		}
		f, err = fconf.Parse()
	default:
		return nil, fmt.Errorf("unknown formatter: %s", config.Formatter)
	}
	if err != nil {
		return nil, err
	}

	handler := Handler{formatter: f}
	handler.levelFrom, err = logrus.ParseLevel(config.LevelFrom)
	if err != nil {
		return nil, fmt.Errorf("cannot parse level_from: %v", err)
	}
	handler.levelTo, err = logrus.ParseLevel(config.LevelTo)
	if err != nil {
		return nil, fmt.Errorf("cannot parse level_to: %v", err)
	}
	return &handler, nil
}
