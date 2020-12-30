package logit

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	formatter logrus.Formatter
	stream    io.Writer
	levelFrom logrus.Level
	levelTo   logrus.Level
}

func (h Handler) Log(entry *logrus.Entry) error {
	if entry.Level < h.levelTo {
		return nil
	}
	if entry.Level > h.levelFrom {
		return nil
	}

	serialized, err := h.formatter.Format(entry)
	if err != nil {
		return fmt.Errorf("cannot format: %v", err)
	}
	_, err = h.stream.Write(serialized)
	if err != nil {
		return fmt.Errorf("cannot write to log: %v", err)
	}
	return nil
}
