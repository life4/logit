package logit

import (
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
