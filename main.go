package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/orsinium-labs/logit/logit"
	"github.com/spf13/pflag"
)

func main() {
	var cpath string
	pflag.StringVarP(&cpath, "config", "c", "logit.toml", "path to the config file")
	pflag.Parse()

	log, err := logit.ReadLogger(cpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err = log.Log(log.SafeParse(scanner.Text()))
		if err != nil {
			fmt.Println(err)
		}
	}

	log.Wait()
}
