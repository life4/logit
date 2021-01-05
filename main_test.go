package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_run_help(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	err := run([]string{"--help"}, nil)
	is.Nil(err)
}

func Test_run_NoConfig(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	err := run([]string{"--config", "non-existent"}, nil)
	is.Error(err)
	exp := "config error: cannot read config file: open non-existent: no such file or directory"
	is.Equal(err.Error(), exp)
}

func Test_run_BadFlag(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	err := run([]string{"--something"}, nil)
	is.Error(err)
	exp := "cli error: unknown flag: --something"
	is.Equal(err.Error(), exp)
}
