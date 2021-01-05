package logit

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func Test_Logger_Parse(t *testing.T) {
	fakeNow := time.Date(2012, 11, 10, 9, 8, 7, 0, time.UTC)
	testCases := []struct {
		desc string
		line string
		exp  logrus.Entry
		err  error
	}{
		{
			desc: "simple json",
			line: `{"level":"fatal","msg":"A huge walrus appears","time":"2020-12-29 15:09:21"}`,
			exp: logrus.Entry{
				Message: "A huge walrus appears",
				Time:    time.Date(2020, 12, 29, 15, 9, 21, 0, time.UTC),
				Level:   logrus.FatalLevel,
				Data:    logrus.Fields{},
			},
		},
		{
			desc: "json with fields",
			line: `{"animal":"walrus","level":"fatal","msg":"A huge walrus appears","time":"2020-12-29 15:09:21"}`,
			exp: logrus.Entry{
				Message: "A huge walrus appears",
				Time:    time.Date(2020, 12, 29, 15, 9, 21, 0, time.UTC),
				Level:   logrus.FatalLevel,
				Data:    logrus.Fields{"animal": "walrus"},
			},
		},
		{
			desc: "plain text",
			line: `oh hi mark`,
			exp: logrus.Entry{
				Message: "oh hi mark",
				Time:    fakeNow,
				Level:   logrus.InfoLevel,
				Data:    logrus.Fields{},
			},
		},
		{
			desc: "invalid json",
			line: `{oh: no}`,
			err:  errors.New(`invalid character 'o' looking for beginning of object key string`),
		},
	}
	config := `
		[[handler]]
		format = "logfmt"
	`

	for _, tcase := range testCases {
		t.Run(tcase.desc, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)
			log, err := MakeLogger(config)
			is.Nil(err)
			log.now = func() time.Time { return fakeNow }
			e, err := log.Parse(tcase.line)
			if tcase.err != nil {
				is.Equal(err.Error(), tcase.err.Error())
				return
			}
			is.Nil(err)
			is.Equal(*e, tcase.exp)
		})
	}
}

func Test_Logger_Log(t *testing.T) {
	testCases := []struct {
		desc string
		line string
		exp  string
	}{
		{
			desc: "simple json",
			line: `{"level":"fatal","msg":"A huge walrus appears","time":"2020-12-29 15:09:21"}`,
			exp:  `time="2020-12-29 15:09:21" level=fatal msg="A huge walrus appears"`,
		},
		{
			desc: "json with field",
			line: `{"animal":"walrus","level":"fatal","msg":"A huge walrus appears","time":"2020-12-29 15:09:21"}`,
			exp:  `time="2020-12-29 15:09:21" level=fatal msg="A huge walrus appears" animal=walrus`,
		},
	}
	config := `
		[[handler]]
		format = "logfmt"
	`

	for _, tcase := range testCases {
		t.Run(tcase.desc, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)
			log, err := MakeLogger(config)
			is.Nil(err)
			var b bytes.Buffer
			log.SetStream(&b)
			err = log.Log(log.SafeParse(tcase.line))
			is.Nil(err)
			actual := b.String()
			is.Equal(strings.TrimSpace(strings.TrimSuffix(actual, "\n")), tcase.exp)
		})
	}
}
