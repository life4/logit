package logit

import "github.com/sirupsen/logrus"

type Level struct{ logrus.Level }

var (
	TraceLevel Level = Level{logrus.TraceLevel}
	PanicLevel Level = Level{logrus.PanicLevel}
)

func (level *Level) UnmarshalText(text []byte) error {
	parsed, err := logrus.ParseLevel(string(text))
	level.Level = parsed
	return err
}
