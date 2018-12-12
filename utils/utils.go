package utils

import (
	"flag"
	"log"

	"github.com/sirupsen/logrus"
)

// Debug is a function that is able to enable debug logging
func Debug() {
	debug := flag.Bool("d", false, "Enable debug mode.")

	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// ErrorLineNumber is a function that adds the linenumber to the error message
func ErrorLineNumber() {
	// https://stackoverflow.com/a/24809859/2777965
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Fatal is a function that returns the error message and an exit code of 1
func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}