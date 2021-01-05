package logit

import (
	"fmt"
	"io/ioutil"
	"time"

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

// ReadLogger reads configuration file by the given path and returns Logger instance.
func ReadLogger(cpath string) (*Logger, error) {
	bytes, err := ioutil.ReadFile(cpath)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %v", err)
	}
	return MakeLogger(string(bytes))
}

// MakeLogger parses the given configuration file and returns Logger instance.
func MakeLogger(content string) (*Logger, error) {
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

	meta, err := toml.Decode(content, &raw)
	if err != nil {
		return nil, fmt.Errorf("cannot read config: %v", err)
	}

	logger := Logger{
		Handlers: make([]*Handler, len(raw.HandlersRaw)),
	}
	for i, primitive := range raw.HandlersRaw {
		handler, err := ParseHandler(meta, primitive)
		if err != nil {
			return nil, fmt.Errorf("cannot parse handler: %v", err)
		}
		logger.Handlers[i] = handler
	}

	undecoded := meta.Undecoded()
	if len(undecoded) != 0 {
		return nil, fmt.Errorf("unknown fields: %v", undecoded)
	}

	logger.Levels.Default, err = logrus.ParseLevel(raw.Levels.Default)
	if err != nil {
		return nil, fmt.Errorf("cannot parse levels.default: %v", err)
	}
	logger.Levels.Error, err = logrus.ParseLevel(raw.Levels.Error)
	if err != nil {
		return nil, fmt.Errorf("cannot parse levels.error: %v", err)
	}

	logger.Fields = raw.Fields
	logger.now = time.Now
	return &logger, nil
}
