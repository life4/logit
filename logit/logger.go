package logit

import (
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type Logger struct {
	loggers []*logrus.Logger
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

func (log Logger) Log(entry *logrus.Entry) {
	for _, sublog := range log.loggers {
		entry.Logger = sublog
		entry.Log(entry.Level, entry.Message)
	}
}

func (log Logger) LogError(err error, msg string) {
	entry := logrus.NewEntry(nil)
	entry = entry.WithError(err)
	for _, sublog := range log.loggers {
		entry.Logger = sublog
		entry.Log(logrus.ErrorLevel, msg)
	}
}

func NewLogger(config *Config) (Logger, error) {
	log := Logger{
		loggers: make([]*logrus.Logger, len(config.Formatters)),
	}
	for i, f := range config.Formatters {
		sublog := logrus.New()
		sublog.SetFormatter(f)
		log.loggers[i] = sublog
	}

	// level, err := config.Level()
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot parse level: %v", err)
	// }
	// log.SetLevel(level)

	return log, nil
}
