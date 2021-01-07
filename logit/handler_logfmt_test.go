package logit

import (
	"bytes"
	"strings"
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
		{
			desc: "level_from exclude",
			conf: `
				format = "logfmt"
				level_from = "panic"
			`,
		},
		{
			desc: "level_from include",
			conf: `
				format = "logfmt"
				level_from = "error"
			`,
			exp: `time="2020-05-04 01:02:03" level=error msg="oh hi mark"`,
		},
		{
			desc: "level_from include #2",
			conf: `
				format = "logfmt"
				level_from = "info"
			`,
			exp: `time="2020-05-04 01:02:03" level=error msg="oh hi mark"`,
		},
		{
			desc: "level_to exclude",
			conf: `
				format = "logfmt"
				level_to = "warning"
			`,
		},
		{
			desc: "level_to include",
			conf: `
				format = "logfmt"
				level_to = "error"
			`,
			exp: `time="2020-05-04 01:02:03" level=error msg="oh hi mark"`,
		},
		{
			desc: "level_to include #2",
			conf: `
				format = "logfmt"
				level_to = "panic"
			`,
			exp: `time="2020-05-04 01:02:03" level=error msg="oh hi mark"`,
		},
	}

	e := logrus.Entry{
		Message: "oh hi mark",
		Level:   logrus.ErrorLevel,
		Time:    time.Date(2020, time.May, 4, 1, 2, 3, 4, time.UTC),
	}
	for _, tcase := range testCases {
		t.Run(tcase.desc, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)
			var p toml.Primitive
			m, err := toml.Decode(tcase.conf, &p)
			is.Nil(err)
			h, err := ParseHandler(m, p)
			is.Nil(err)
			var b bytes.Buffer
			h.SetStream(&b)
			err = h.Log(&e)
			is.Nil(err)
			actual := b.String()
			is.Equal(strings.TrimSuffix(actual, "\n"), tcase.exp)
		})
	}
}
