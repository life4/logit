package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/orsinium-labs/logit/logit"
	"github.com/stretchr/testify/require"
)

func Test_handle(t *testing.T) {
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
			log, err := logit.MakeLogger(config)
			is.Nil(err)
			var b bytes.Buffer
			log.SetStream(&b)
			err = handle(log, tcase.line)
			is.Nil(err)
			actual := b.String()
			is.Equal(strings.TrimSpace(strings.TrimSuffix(actual, "\n")), tcase.exp)
		})
	}
}
