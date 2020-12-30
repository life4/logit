package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type CMain struct {
	Level     string
	Formatter string
}

// type cLevels struct {
// 	Panic   string
// 	Fatal   string
// 	Warning string
// 	Info    string
// 	Debug   string
// 	Trace   string
// }

type CHandler struct {
	Formatter string
}

type Config struct {
	Main        CMain
	HandlersRaw []toml.Primitive `toml:"handler"`
	formatters  []logrus.Formatter
}

func NewConfig() Config {
	c := Config{
		Main: CMain{
			Level:     "trace",
			Formatter: "text",
		},
	}
	return c
}

func ReadConfig(cpath string) (*Config, error) {
	config := NewConfig()
	if cpath == "" {
		return &config, nil
	}

	meta, err := toml.DecodeFile(cpath, &config)
	if err != nil {
		return nil, fmt.Errorf("cannot read config: %v", err)
	}

	config.formatters = make([]logrus.Formatter, len(config.HandlersRaw))
	for i, primitive := range config.HandlersRaw {
		f, err := parseFormatter(meta, primitive)
		if err != nil {
			return nil, fmt.Errorf("cannot parse handler: %v", err)
		}
		config.formatters[i] = f
	}

	// fmt.Println(config.Handlers[0])
	// undecoded := meta.Undecoded()
	// if len(undecoded) != 0 {
	// 	return nil, fmt.Errorf("unknown fields: %v", undecoded)
	// }
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

func (c *Config) Formatter() (logrus.Formatter, error) {
	return c.formatters[0], nil
}

func (c *Config) Level() (logrus.Level, error) {
	return logrus.ParseLevel(c.Main.Level)
}
