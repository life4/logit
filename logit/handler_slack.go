// +build h_slack h_all

package logit

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/johntdyer/slackrus"
	"github.com/sirupsen/logrus"
)

type SlackHandler struct {
	BaseHandler
	HookURL   string `toml:"hook_url"`
	IconURL   string `toml:"icon_url"`
	Channel   string
	IconEmoji string `toml:"icon_emoji"`
	Username  string
}

func NewSlackHandler() SlackHandler {
	return SlackHandler{
		BaseHandler: NewBaseHandler(),
	}
}

func (config SlackHandler) Parse() (*Handler, error) {
	hook := slackrus.SlackrusHook{
		HookURL:        config.HookURL,
		AcceptedLevels: slackrus.LevelThreshold(logrus.TraceLevel),
		Channel:        config.Channel,
		IconEmoji:      config.IconEmoji,
		Username:       config.Username,
	}

	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.hook = &hook
	return h, nil
}

func init() {
	RegisterParser("slack", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (*Handler, error) {
		fconf := NewSlackHandler()
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
