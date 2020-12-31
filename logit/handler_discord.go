package logit

import (
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
