package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type CLevels struct {
	Default string `toml:"default"`
	Error   string `toml:"error"`
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
	Levels      CLevels
	Fields      CFields
	HandlersRaw []toml.Primitive `toml:"handler"`
}

type Levels struct {
	Default logrus.Level
	Error   logrus.Level
}

// RawConfig represents the config how it is used internally.
// It is generated from RawConfig.
type Config struct {
	Levels   Levels
	Handlers []Handler
	Fields   CFields
}

func ReadConfig(cpath string) (*Config, error) {
	raw := RawConfig{
		Levels: CLevels{
			Default: "INFO",
			Error:   "ERROR",
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
		handler, err := ParseHandler(meta, primitive)
		if err != nil {
			return nil, fmt.Errorf("cannot parse handler: %v", err)
		}
		config.Handlers[i] = *handler
	}

	undecoded := meta.Undecoded()
	if len(undecoded) != 0 {
		return nil, fmt.Errorf("unknown fields: %v", undecoded)
	}

	config.Levels.Default, err = logrus.ParseLevel(raw.Levels.Default)
	if err != nil {
		return nil, fmt.Errorf("cannot parse levels.default: %v", err)
	}
	config.Levels.Error, err = logrus.ParseLevel(raw.Levels.Error)
	if err != nil {
		return nil, fmt.Errorf("cannot parse levels.error: %v", err)
	}

	config.Fields = raw.Fields
	return &config, nil
}
