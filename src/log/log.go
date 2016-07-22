// Package "log" handles logging.
package log

import (
	"log"
)

type Logger struct {
	Verbose bool
}

// New instantiates and returns a new logger instance.
func New(verbose bool) *Logger {
	return &Logger {
		Verbose: verbose,
	}
}

// Printf outputs a log message if Verbose is enabled.
func (l Logger) Printf(s string, v ...interface{}) {
	if l.Verbose {
		l.ForcePrintf(s, v...)
	}
}

// ForcePrintf outputs a log message regardless of verbose logging.
func (l Logger) ForcePrintf(s string, v ...interface{}) {
	log.Printf(s, v...)
}

// Error outputs an error to the logs, regardless of verbose logging.
func (l Logger) Error(err error) {
	log.Fatal(err)
}