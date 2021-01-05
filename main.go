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
		line := scanner.Text()

		// parse
		entry, err := log.Parse(line)
		if err != nil {
			err = fmt.Errorf("cannot parse entry: %v", err)
			err = log.LogError(err, line)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		// log
		err = log.Log(entry)
		if err != nil {
			fmt.Println(err)
		}
	}

	log.Wait()
}
