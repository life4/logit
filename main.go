package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/orsinium-labs/logit/logit"
)

func main() {
	var cpath string
	flag.StringVar(&cpath, "", "logit.toml", "")
	flag.Parse()

	config, err := logit.ReadConfig(cpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log, err := logit.NewLogger(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := log.Parse(line)
		if err != nil {
			err = fmt.Errorf("cannot parse entry: %v", err)
			fmt.Println(err)
			// log.WithError(err).Error(err)
		}
		log.Log(entry)
	}

}
