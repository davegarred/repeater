// Package log provides simple logging shortcuts
package log

import (
	"log"
	"log/syslog"
	"os"
)

var logger *log.Logger

// Log provides a shortcut to the configured logger that maps to a Printf
func Log(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

// SetLogFile allows the user to specify the file where logs are stored
func SetLogFile(l *os.File) {
	logger = log.New(l, "repeater ", log.LstdFlags|log.Lshortfile)
}

func init() {
	// default to system logger (usually /var/log/syslog)
	var err error
	logger, err = syslog.NewLogger(syslog.LOG_LOCAL3|syslog.LOG_NOTICE, log.Lshortfile)
	if err != nil {
		panic(err)
	}
}
