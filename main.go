package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/life4/logit/logit"
	"github.com/spf13/pflag"
)

func run(args []string, stream io.Reader) error {
	parser := pflag.NewFlagSet("logit", pflag.ContinueOnError)
	var cpath string
	parser.StringVarP(&cpath, "config", "c", "logit.toml", "path to the config file")
	err := parser.Parse(args)
	if err != nil {
		if err == pflag.ErrHelp {
			return nil
		}
		return fmt.Errorf("cli error: %v", err)
	}

	log, err := logit.ReadLogger(cpath)
	if err != nil {
		return fmt.Errorf("config error: %v", err)
	}

	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		err = log.Log(log.SafeParse(scanner.Text()))
		if err != nil {
			fmt.Println(err)
		}
	}

	log.Wait()
	return nil
}

func main() {
	err := run(os.Args[1:], os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
