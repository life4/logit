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
	log, err := logit.NewLogger(cpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := logit.ParseEntry(log, line)
		if err != nil {
			err = fmt.Errorf("cannot parse entry: %v", err)
			log.WithError(err).Error(err)
		}
		entry.Log(entry.Level, entry.Message)
	}

}
