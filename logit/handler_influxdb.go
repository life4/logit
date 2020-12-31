package logit

import (
	"github.com/Abramovic/logrus_influxdb"
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

func (config InfluxDBHandler) Parse() (*Handler, error) {
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
	h.hook = hook
	return h, nil
}
