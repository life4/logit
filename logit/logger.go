package logit

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func NewLogger(config *Config) (*logrus.Logger, error) {

	log := logrus.New()

	formatter, err := config.Formatter()
	if err != nil {
		return nil, fmt.Errorf("cannot init formatter: %v", err)
	}
	log.SetFormatter(formatter)

	level, err := config.Level()
	if err != nil {
		return nil, fmt.Errorf("cannot parse level: %v", err)
	}
	log.SetLevel(level)

	return log, nil
}
