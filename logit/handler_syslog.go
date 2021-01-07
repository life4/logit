// +build h_syslog !h_clean

package logit

import (
	"fmt"
	"log/syslog"
	"strings"

	"github.com/BurntSushi/toml"
	syslogrus "github.com/sirupsen/logrus/hooks/syslog"
)

type SysLogHandler struct {
	BaseHandler
	Network  string
	Address  string
	Tag      string
	Priority string
}

func NewSysLogHandler() SysLogHandler {
	return SysLogHandler{
		BaseHandler: NewBaseHandler(),
		Tag:         "logit",
	}
}

func (config SysLogHandler) Parse() (Handler, error) {
	priority, err := config.parsedPriority()
	if err != nil {
		return nil, err
	}
	hook, err := syslogrus.NewSyslogHook(
		config.Network,
		config.Address,
		priority,
		config.Tag,
	)
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

func (config SysLogHandler) parsedPriority() (syslog.Priority, error) {
	switch strings.ToLower(config.Priority) {
	case "emerg":
		return syslog.LOG_EMERG, nil
	case "alert":
		return syslog.LOG_ALERT, nil
	case "crit":
		return syslog.LOG_CRIT, nil
	case "err":
		return syslog.LOG_ERR, nil
	case "warning":
		return syslog.LOG_WARNING, nil
	case "notice":
		return syslog.LOG_NOTICE, nil
	case "info":
		return syslog.LOG_INFO, nil
	case "debug":
		return syslog.LOG_DEBUG, nil
	}
	return syslog.Priority(0), fmt.Errorf("unknown priority: %s", config.Priority)
}

func init() {
	RegisterParser("syslog", func(
		meta toml.MetaData,
		primitive toml.Primitive,
	) (Handler, error) {
		fconf := NewSysLogHandler()
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
