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
			return nil, fmt.Errorf("text config: %v", err)
		}
		return fconf.Parse()
	case "logfmt":
		fconf := NewLogFmtHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("logfmt config: %v", err)
		}
		return fconf.Parse()
	case "json":
		fconf := NewJSONHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("json config: %v", err)
		}
		return fconf.Parse()
	case "zalgo":
		fconf := NewZalgoHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("zalgo config: %v", err)
		}
		return fconf.Parse()
	case "syslog":
		fconf := NewSysLogHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("syslog config: %v", err)
		}
		return fconf.Parse()
	case "sentry":
		fconf := NewSentryHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("sentry config: %v", err)
		}
		return fconf.Parse()
	case "logstash":
		fconf := NewLogstashHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("logstash config: %v", err)
		}
		return fconf.Parse()
	case "elastic":
		fconf := NewElasticHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("elastic config: %v", err)
		}
		return fconf.Parse()
	case "slack":
		fconf := NewSlackHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("slack config: %v", err)
		}
		return fconf.Parse()
	case "gcloud":
		fconf := NewGCloudHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("gcloud config: %v", err)
		}
		return fconf.Parse()
	case "graylog":
		fconf := NewGraylogHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("graylog config: %v", err)
		}
		return fconf.Parse()
	case "fluentd":
		fconf := NewFluentdHandler()
		err = meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("fluentd config: %v", err)
		}
		return fconf.Parse()
	default:
		return nil, fmt.Errorf("unknown format: %s", config.Format)
	}
}
