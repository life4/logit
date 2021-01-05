package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/orsinium-labs/logit/logit"
	"github.com/spf13/pflag"
)

// handle parses the given text line and logs it.
// Error returned if the record cannot be logged.
// Basically, it can happen only if something is wrong in handler upstream.
// So, the returned error shouldn't be logged by the logger itself.
func handle(log *logit.Logger, line string) error {
	// parse
	entry, err := log.Parse(line)
	if err != nil {
		err = fmt.Errorf("cannot parse entry: %v", err)
		return log.LogError(err, line)
	}
	// log
	return log.Log(entry)
}

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
		err = handle(log, scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
	}

	log.Wait()
}
