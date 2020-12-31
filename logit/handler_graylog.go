package logit

import graylog "github.com/gemnasium/logrus-graylog-hook"

type GraylogHandler struct {
	BaseHandler
	Address string
}

func NewGraylogHandler() GraylogHandler {
	return GraylogHandler{
		BaseHandler: NewBaseHandler(),
	}
}

func (config GraylogHandler) Parse() (*Handler, error) {
	hook := graylog.NewGraylogHook(config.Address, map[string]interface{}{})
	h, err := config.BaseHandler.Parse()
	if err != nil {
		return nil, err
	}
	h.hook = hook
	return h, nil
}
