package logit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/araddon/dateparse"
	"github.com/sirupsen/logrus"
)

// Levels represent configuration for default levels
type Levels struct {
	Default logrus.Level
	Error   logrus.Level
}

type Logger struct {
	Levels   Levels
	Handlers []Handler
	Fields   CFields
	now      func() time.Time
}

func (log *Logger) SetStream(stream io.Writer) {
	for _, h := range log.Handlers {
		h.SetStream(stream)
	}
}

func (log Logger) SafeParse(line string) *logrus.Entry {
	entry, err := log.Parse(line)
	if err != nil {
		err = fmt.Errorf("cannot parse entry: %v", err)
		entry = logrus.NewEntry(nil)
		entry = entry.WithField("error", err.Error())
		entry.Level = log.Levels.Error
		entry.Message = line
		entry.Time = log.now()
	}
	return entry
}

func (log Logger) Parse(line string) (*logrus.Entry, error) {
	// If non-JSON passed, use it as message, set the default level
	if line == "" || line[0] != '{' {
		entry := logrus.NewEntry(nil)
		entry.Level = log.Levels.Default
		entry.Message = line
		entry.Time = log.now()
		return entry, nil
	}

	e := logrus.NewEntry(nil)

	err := json.Unmarshal([]byte(line), &e.Data)
	if err != nil {
		return nil, err
	}

	// extract message
	msgRaw, ok := e.Data[log.Fields.Message]
	if !ok {
		return nil, errors.New("cannot find message field")
	}
	e.Message, ok = msgRaw.(string)
	if !ok {
		return nil, errors.New("message is not a string")
	}
	delete(e.Data, log.Fields.Message)

	// extract level
	lvlRaw, ok := e.Data[log.Fields.Level]
	if !ok {
		return nil, errors.New("cannot find level field")
	}
	lvlStr, ok := lvlRaw.(string)
	if !ok {
		return nil, errors.New("level is not a string")
	}
	e.Level, err = logrus.ParseLevel(lvlStr)
	if err != nil {
		return nil, err
	}
	delete(e.Data, log.Fields.Level)

	// extract time
	timeRaw, ok := e.Data[log.Fields.Time]
	if ok {
		timeStr, ok := timeRaw.(string)
		if !ok {
			return nil, errors.New("time is not a string")
		}
		e.Time, err = dateparse.ParseAny(timeStr)
		if err != nil {
			return nil, err
		}
		delete(e.Data, log.Fields.Time)
	} else {
		e.Time = log.now()
	}

	return e, nil
}

func (log Logger) Wait() {
	for _, handler := range log.Handlers {
		handler.Wait()
	}
}

func (log Logger) Log(entry *logrus.Entry) error {
	for _, handler := range log.Handlers {
		err := handler.Log(entry)
		if err != nil {
			return fmt.Errorf("cannot write log entry: %v", err)
		}
	}
	return nil
}
