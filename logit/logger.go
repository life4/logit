package logit

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type Logger struct {
	handlers []Handler
}

func (Logger) Parse(line string) (*logrus.Entry, error) {
	parsed := gjson.Parse(line)
	fields := make(logrus.Fields)
	for k, v := range parsed.Map() {
		fields[k] = v.String()
	}
	e := logrus.NewEntry(nil)
	e = e.WithFields(fields)
	e.Level = logrus.InfoLevel
	e.Message = parsed.Get("msg").String()
	return e, nil
}

func (log Logger) Log(entry *logrus.Entry) error {
	for _, handler := range log.handlers {
		err := handler.Log(entry)
		if err != nil {
			return fmt.Errorf("cannot write log entry: %v", err)
		}
	}
	return nil
}

func (log Logger) LogError(err error, msg string) error {
	entry := logrus.NewEntry(nil)
	entry = entry.WithError(err)
	return log.Log(entry)
}

func NewLogger(config *Config) (Logger, error) {
	log := Logger{
		handlers: make([]Handler, len(config.Handlers)),
	}
	for i, handler := range config.Handlers {
		log.handlers[i] = handler
	}

	// level, err := config.Level()
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot parse level: %v", err)
	// }
	// log.SetLevel(level)

	return log, nil
}
