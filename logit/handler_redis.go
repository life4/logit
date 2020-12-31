package logit

import (
	logredis "github.com/rogierlommers/logrus-redis-hook"
)

type RedisHandler struct {
	BaseHandler

	Host     string
	Port     int
	Password string

	Key        string
	Format     string
	App        string
	SourceHost string `toml:"source_host"`
	DB         int
	TTL        int
}

func NewRedisHandler() RedisHandler {
	return RedisHandler{
		BaseHandler: NewBaseHandler(),
		Host:        "localhost",
		Port:        6379,
		TTL:         3600,
		Key:         "logit",
	}
}

func (config RedisHandler) Parse() (*Handler, error) {
	hook, err := logredis.NewHook(logredis.HookConfig{
		Hostname: config.SourceHost,
		Port:     config.Port,
		Password: config.Password,

		Key:    config.Key,
		Format: config.Format,
		App:    config.App,
		Host:   config.Host,
		DB:     config.DB,
		TTL:    config.TTL,
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
