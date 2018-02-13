package main

import (
	"flag"
	"os"

	"github.com/davegarred/repeater/grpcserver"
	"github.com/davegarred/repeater/log"
	"github.com/davegarred/repeater/persist"
	"github.com/davegarred/repeater/web"
)

var defaultPort = ":8080"

func main() {
	logfileName := flag.String("log", "", "Location of the log file to use")
	diskStorage := flag.Bool("disk", false, "Store objects on disk (default is in-memory)")
	flag.Parse()
	if *logfileName != "" {
		if logfile, e := os.Create(*logfileName); e == nil {
			log.SetLogFile(logfile)
			log.Log("Logger attached to log at %v", *logfileName)
			defer logfile.Close()
		}
	}
	var store web.Storer
	if *diskStorage {
		home := os.Getenv("HOME")
		store = persist.NewLocalStore(home + "/repeater")
	} else {
		store = persist.NewMemStore()
	}
	go server.StartGRPCServer(store, defaultPort)

	web.Start(store)
}
