package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/orsinium-labs/logit/logit"
	"github.com/sirupsen/logrus"
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

	log.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
