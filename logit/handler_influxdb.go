// +build h_influxdb !h_clean

package logit

import (
	"fmt"

	"github.com/Abramovic/logrus_influxdb"
	"github.com/BurntSushi/toml"
)

type InfluxDBHandler struct {
	BaseHandler

	Host     string
	Port     int
	Username string
	Password string

	Database   string
	Precision  string
	UseHTTPS   bool `toml:"use_https"`
	BatchCount int  `toml:"batch_count"`
}

func NewInfluxDBHandler() InfluxDBHandler {
	return InfluxDBHandler{
		BaseHandler: NewBaseHandler(),
		Host:        "localhost",
		Port:        6379,
		Database:    "logit",
		Precision:   "ns",
	}
}

func (config InfluxDBHandler) Parse() (Handler, error) {
	hook, err := logrus_influxdb.NewInfluxDB(&logrus_influxdb.Config{
		Host:       config.Host,
		Port:       config.Port,
		Database:   config.Database,
		Username:   config.Username,
		Password:   config.Password,
		UseHTTPS:   config.UseHTTPS,
		Precision:  config.Precision,
		BatchCount: config.BatchCount,
	})
	if err != nil {
		return nil, err
	}

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.SetHook(hook)
	return h, nil
}

func init() {
	RegisterParser("influxdb", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (Handler, error) {
		fconf := NewInfluxDBHandler()
		err := meta.PrimitiveDecode(primitive, &fconf)
		if err != nil {
			return nil, fmt.Errorf("parse: %v", err)
		}
		handler, err := fconf.Parse()
		if err != nil {
			return nil, fmt.Errorf("init: %v", err)
		}
		return handler, nil
	})
}
