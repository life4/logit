package logit

import (
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func ParseEntry(log *logrus.Logger, line string) (*logrus.Entry, error) {
	parsed := gjson.Parse(line)
	fields := make(logrus.Fields)
	for k, v := range parsed.Map() {
		fields[k] = v.String()
	}
	e := log.WithFields(fields)
	e.Level = logrus.InfoLevel
	e.Message = parsed.Get("msg").String()
	return e, nil
}
