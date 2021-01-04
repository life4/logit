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

func Test_TextHandler(t *testing.T) {
	testCases := []struct {
		desc string
		conf string
		exp  string
	}{
		{
			desc: "simple",
			conf: `format = "text"`,
			exp:  "\x1b[31mERROR  \x1b[0m[2020-05-04 01:02:03] oh hi mark",
		},
		{
			desc: "truncate_level",
			conf: `
				format = "text"
				truncate_level = true
			`,
			exp: "\x1b[31mERRO\x1b[0m[2020-05-04 01:02:03] oh hi mark",
		},
		{
			desc: "timestamp",
			conf: `
				format = "text"
				timestamp = "dd.MM.YYYY HH.mm.ss"
			`,
			exp: "\x1b[31mERROR  \x1b[0m[04.05.2020 01.02.03] oh hi mark",
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
			h.stream = &b
			err = h.Log(&e)
			is.Nil(err)
			actual := b.String()
			is.Equal(strings.TrimSpace(strings.TrimSuffix(actual, "\n")), tcase.exp)
		})
	}
}
