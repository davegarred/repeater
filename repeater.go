package main

import (
	"flag"
	"os"

	"github.com/davegarred/repeater/persist"
	"github.com/davegarred/repeater/util"
	"github.com/davegarred/repeater/web"
)

func main() {
	logfileName := flag.String("log", "", "Location of the log file to use")
	flag.Parse()
	if *logfileName != "" {
		if logfile, e := os.Create(*logfileName); e == nil {
			util.SetLogFile(logfile)
			util.Log("Logger attached to log at %v", *logfileName)
			defer logfile.Close()
		}
	}
	store := persist.NewMemStore()
	web.Start(store)
}
