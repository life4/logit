package logit

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Config(t *testing.T) {
	fakeNow := time.Date(2012, 11, 10, 9, 8, 7, 0, time.UTC)
	testCases := []struct {
		desc string
		conf string
		line string
		exp  string
	}{
		{
			desc: "levels",
			conf: `
				[levels]
				default = "error"

				[[handler]]
				format = "logfmt"
			`,
			line: `oh hi mark`,
			exp:  `time="2012-11-10 09:08:07" level=error msg="oh hi mark"`,
		},
		{
			desc: "defaults plain ",
			conf: `
				[defaults]
				hello = "world"

				[[handler]]
				format = "logfmt"
			`,
			line: `oh hi mark`,
			exp:  `time="2012-11-10 09:08:07" level=info msg="oh hi mark" hello=world`,
		},
		{
			desc: "defaults json",
			conf: `
				[defaults]
				hello = "world"

				[[handler]]
				format = "logfmt"
			`,
			line: `{"animal":"walrus","level":"info","msg":"oh hi mark","time":"2020-12-29 15:09:21"}`,
			exp:  `time="2020-12-29 15:09:21" level=info msg="oh hi mark" animal=walrus hello=world`,
		},
	}

	for _, tcase := range testCases {
		t.Run(tcase.desc, func(t *testing.T) {
			is := require.New(t)
			log, err := MakeLogger(tcase.conf)
			is.Nil(err)
			var b bytes.Buffer
			log.SetStream(&b)
			log.now = func() time.Time { return fakeNow }
			err = log.Log(log.SafeParse(tcase.line))
			is.Nil(err)
			actual := b.String()
			is.Equal(strings.TrimSuffix(actual, "\n"), tcase.exp)
		})
	}
}
