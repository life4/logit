package logit

import (
	"fmt"
	"io"
	"sync"

	"github.com/sirupsen/logrus"
)

type Handler interface {
	Wait()
	Log(*logrus.Entry) error
	SetFormatter(logrus.Formatter)
	SetHook(logrus.Hook)
	SetStream(io.Writer)
}

type HandlerSync struct {
	hook      logrus.Hook
	formatter logrus.Formatter
	stream    io.Writer
	levelFrom logrus.Level
	levelTo   logrus.Level
	wait      func()
}

func (handler *HandlerSync) SetFormatter(f logrus.Formatter) {
	handler.formatter = f
}

func (handler *HandlerSync) SetHook(hook logrus.Hook) {
	handler.hook = hook
}
func (handler *HandlerSync) SetStream(stream io.Writer) {
	handler.stream = stream
}

func (handler HandlerSync) Wait() {
	if handler.wait != nil {
		handler.wait()
	}
}

func (handler HandlerSync) Log(entry *logrus.Entry) error {
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

type HandlerAsync struct {
	workers int
	handler HandlerSync

	started bool
	wg      *sync.WaitGroup
	entries chan *logrus.Entry
	errors  chan error
}

func (handler *HandlerAsync) SetFormatter(f logrus.Formatter) {
	handler.handler.SetFormatter(f)
}

func (handler *HandlerAsync) SetHook(hook logrus.Hook) {
	handler.handler.SetHook(hook)
}
func (handler *HandlerAsync) SetStream(stream io.Writer) {
	handler.handler.SetStream(stream)
}

func (h HandlerAsync) worker() {
	for e := range h.entries {
		err := h.handler.Log(e)
		if err != nil {
			h.errors <- err
		}
	}
	h.wg.Done()
}

func (h *HandlerAsync) start() {
	h.started = true
	h.entries = make(chan *logrus.Entry)
	h.errors = make(chan error, h.workers)
	h.wg = &sync.WaitGroup{}
	for i := 0; i < h.workers; i++ {
		h.wg.Add(1)
		go h.worker()
	}
}

func (h *HandlerAsync) Wait() {
	close(h.entries)
	h.wg.Wait()
}

func (h *HandlerAsync) Log(entry *logrus.Entry) error {
	if !h.started {
		h.start()
	}
	h.entries <- entry

	select {
	case err := <-h.errors:
		return err
	default:
		return nil
	}
}
