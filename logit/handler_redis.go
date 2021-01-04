// +build h_redis !h_clean

package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
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
	Database   int
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
		DB:     config.Database,
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

func init() {
	RegisterParser("redis", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (*Handler, error) {
		fconf := NewRedisHandler()
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
