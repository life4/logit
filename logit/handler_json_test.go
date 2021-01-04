package logit

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func Test_JSONHandler(t *testing.T) {
	testCases := []struct {
		desc   string
		conf   string
		exp    map[string]interface{}
		fields logrus.Fields
	}{
		{
			desc: "simple",
			conf: `format = "json"`,
			exp: map[string]interface{}{
				"level": "error",
				"msg":   "oh hi mark",
				"time":  "2020-05-04 01:02:03",
			},
		},
		{
			desc: "timestamp",
			conf: `
				format = "json"
				timestamp = "dd.MM.YYYY HH.mm.ss"
			`,
			exp: map[string]interface{}{
				"level": "error",
				"msg":   "oh hi mark",
				"time":  "04.05.2020 01.02.03",
			},
		},
		{
			desc: "additional data",
			conf: `format = "json"`,
			exp: map[string]interface{}{
				"level": "error",
				"msg":   "oh hi mark",
				"time":  "2020-05-04 01:02:03",
				"check": "it",
			},
			fields: logrus.Fields{"check": "it"},
		},
		{
			desc: "data_key",
			conf: `
				format = "json"
				data_key = "data"
			`,
			exp: map[string]interface{}{
				"level": "error",
				"msg":   "oh hi mark",
				"time":  "2020-05-04 01:02:03",
				"data": map[string]interface{}{
					"check": "it",
				},
			},
			fields: logrus.Fields{"check": "it"},
		},
	}

	for _, tcase := range testCases {
		t.Run(tcase.desc, func(t *testing.T) {
			t.Parallel()
			e := logrus.Entry{
				Message: "oh hi mark",
				Level:   logrus.ErrorLevel,
				Time:    time.Date(2020, time.May, 4, 1, 2, 3, 4, time.UTC),
				Data:    tcase.fields,
			}

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
			var actual map[string]interface{}
			err = json.Unmarshal(b.Bytes(), &actual)
			is.Nil(err)
			is.Equal(actual, tcase.exp)
		})
	}
}
