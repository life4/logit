package logit

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Async bool

	hook      logrus.Hook
	formatter logrus.Formatter
	stream    io.Writer
	levelFrom logrus.Level
	levelTo   logrus.Level
	wait      func()
}

func (handler Handler) Wait() {
	if handler.wait != nil {
		handler.wait()
	}
}

func (handler Handler) Log(entry *logrus.Entry) error {
	if entry.Level < handler.levelTo {
		return nil
	}
	if entry.Level > handler.levelFrom {
		return nil
	}

	if handler.hook != nil {
		return handler.hook.Fire(entry)
	}

	serialized, err := handler.formatter.Format(entry)
	if err != nil {
		return fmt.Errorf("cannot format: %v", err)
	}
	_, err = handler.stream.Write(serialized)
	if err != nil {
		return fmt.Errorf("cannot write to log: %v", err)
	}
	return nil
}
