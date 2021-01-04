package logit

import (
	"bytes"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func Test_LogFmtHandler(t *testing.T) {
	testCases := []struct {
		desc string
		conf string
		exp  string
	}{
		{
			desc: "simple",
			conf: `format = "logfmt"`,
			exp:  `time="2020-05-04 01:02:03" level=error msg="oh hi mark"`,
		},
		{
			desc: "timestamp",
			conf: `
				format = "logfmt"
				timestamp = "dd.MM.YYYY HH.mm.ss"
			`,
			exp: `time="04.05.2020 01.02.03" level=error msg="oh hi mark"`,
		},
	}

	for _, tcase := range testCases {
		t.Run(tcase.desc, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)
			e := logrus.Entry{
				Message: "oh hi mark",
				Level:   logrus.ErrorLevel,
				Time:    time.Date(2020, time.May, 4, 1, 2, 3, 4, time.UTC),
			}

			var p toml.Primitive
			m, err := toml.Decode(tcase.conf, &p)
			is.Nil(err)
			h, err := ParseHandler(m, p)
			is.Nil(err)
			var b bytes.Buffer
			h.stream = &b
			err = h.Log(&e)
			is.Nil(err)
			actual := b.String()
			is.Equal(actual, tcase.exp+"\n")

		})
	}
}
