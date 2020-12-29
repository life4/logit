package logit

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type cMain struct {
	Level     string
	Formatter string
}

type cLevels struct {
	Text FText
}

type cFormatters struct {
	Text FText
}

type Config struct {
	Main       cMain
	Levels     cLevels
	Formatters cFormatters
}

func NewConfig() Config {
	c := Config{
		Main: cMain{
			Level:     "trace",
			Formatter: "text",
		},
		Formatters: cFormatters{
			Text: NewFText(),
		},
	}
	return c
}

func (c *Config) Formatter() (logrus.Formatter, error) {
	switch c.Main.Formatter {
	case "text":
		return FTextParse(c.Formatters.Text)
	default:
		return nil, fmt.Errorf("unknown formatter: %s", c.Main.Formatter)
	}
}

func (c *Config) Level() (logrus.Level, error) {
	return logrus.ParseLevel(c.Main.Level)
}
