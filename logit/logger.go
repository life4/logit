package logit

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	config   Config
	handlers []Handler
}

func (log Logger) Parse(line string) (*logrus.Entry, error) {
	// If non-JSON passed, use it as message, set the default level
	if line == "" || line[0] != '{' {
		entry := logrus.NewEntry(nil)
		entry.Level = log.config.Levels.Default
		entry.Message = line
		return entry, nil
	}

	fields := make(logrus.Fields)
	err := json.Unmarshal([]byte(line), &fields)
	if err != nil {
		return nil, err
	}

	// extract message
	msgRaw, ok := fields[log.config.Fields.Message]
	if !ok {
		return nil, errors.New("cannot find message field")
	}
	msgStr, ok := msgRaw.(string)
	if !ok {
		return nil, errors.New("message is not a string")
	}
	delete(fields, log.config.Fields.Message)

	// extract level
	lvlRaw, ok := fields[log.config.Fields.Level]
	if !ok {
		return nil, errors.New("cannot find level field")
	}
	lvlStr, ok := lvlRaw.(string)
	if !ok {
		return nil, errors.New("level is not a string")
	}
	lvl, err := logrus.ParseLevel(lvlStr)
	if err != nil {
		return nil, err
	}
	delete(fields, log.config.Fields.Level)

	// extract time
	timeRaw, ok := fields[log.config.Fields.Time]
	if !ok {
		return nil, errors.New("cannot find time field")
	}
	timeStr, ok := timeRaw.(string)
	if !ok {
		return nil, errors.New("time is not a string")
	}
	time, err := dateparse.ParseAny(timeStr)
	if err != nil {
		return nil, err
	}
	delete(fields, log.config.Fields.Time)

	e := logrus.NewEntry(nil)
	e = e.WithFields(fields)

	e.Message = msgStr
	e.Level = lvl
	e.Time = time
	return e, nil
}

func (log Logger) Wait() {
	for _, handler := range log.handlers {
		handler.Wait()
	}
}

func (log Logger) Log(entry *logrus.Entry) error {
	for _, handler := range log.handlers {
		// run in background
		if handler.Async {
			go func() {
				err := handler.Log(entry)
				if err != nil {
					fmt.Printf("cannot write log entry: %v", err)
				}
			}()
			continue
		}

		// run synchronously
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
	entry.Level = log.config.Levels.Error
	entry.Message = msg
	return log.Log(entry)
}

func NewLogger(config Config) (Logger, error) {
	log := Logger{
		handlers: make([]Handler, len(config.Handlers)),
	}
	for i, handler := range config.Handlers {
		log.handlers[i] = handler
	}
	log.config = config
	return log, nil
}
