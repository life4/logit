// +build h_discord !h_clean

package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/kz/discordrus"
	"github.com/sirupsen/logrus"
)

type DiscordHandler struct {
	BaseHandler
	URL       string
	Username  string
	Author    string
	Inline    bool
	Timestamp string
}

func NewDiscordHandler() DiscordHandler {
	return DiscordHandler{
		BaseHandler: NewBaseHandler(),
		Username:    "logit",
		Inline:      true,
		Timestamp:   "YYYY-MM-dd HH:mm:ss",
	}
}

func (config DiscordHandler) Parse() (*Handler, error) {
	hook := discordrus.NewHook(
		config.URL,
		logrus.TraceLevel,
		&discordrus.Opts{
			Username:            config.Username,
			Author:              config.Author,
			DisableTimestamp:    config.Timestamp == "",
			TimestampFormat:     convertDateFormat(config.Timestamp),
			DisableInlineFields: !config.Inline,
		},
	)

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.hook = hook
	return h, nil
}

func init() {
	RegisterParser("discord", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (*Handler, error) {
		fconf := NewDiscordHandler()
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
