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

type CFormatters struct {
	Text FText
	JSON FJSON
}

type Config struct {
	Main       CMain
	Formatters CFormatters
}

func NewConfig() Config {
	c := Config{
		Main: CMain{
			Level:     "trace",
			Formatter: "text",
		},
		Formatters: CFormatters{
			Text: NewFText(),
			JSON: NewFJSON(),
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
	undecoded := meta.Undecoded()
	if len(undecoded) != 0 {
		return nil, fmt.Errorf("unknown fields: %v", undecoded)
	}
	return &config, nil
}

func (c *Config) Formatter() (logrus.Formatter, error) {
	switch c.Main.Formatter {
	case "text":
		return FTextParse(c.Formatters.Text)
	case "json":
		return FJSONParse(c.Formatters.JSON)
	default:
		return nil, fmt.Errorf("unknown formatter: %s", c.Main.Formatter)
	}
}

func (c *Config) Level() (logrus.Level, error) {
	return logrus.ParseLevel(c.Main.Level)
}
