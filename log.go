package main

import (
	"fmt"
)

const (
	LOG_INFO = 1 << iota
	LOG_ERROR
	LOG_DEBUG
	LOG_VERBOSE
)

var (
	loggingLevel = LOG_INFO | LOG_ERROR
)

func logVerbose(message string, args ...interface{}) {
	if 0 != LOG_VERBOSE&loggingLevel {
		log("LOG", message, args...)
	}
}

func logInfo(message string, args ...interface{}) {
	if 0 != LOG_INFO&loggingLevel {
		log("INFO", message, args...)
	}
}

func logDebug(message string, args ...interface{}) {
	if 0 != LOG_DEBUG&loggingLevel {
		log("DEBUG", message, args...)
	}
}

func logIgnoredError(err error) {
	if nil != err {
		logDebug("Ignoring Error %s", err)
	}
}

func logErrorf(message string, args ...interface{}) {
	log("ERROR:", message, args...)
}

func log(key, message string, args ...interface{}) {
	fmt.Printf("%s: %s\n", key, fmt.Sprintf(message, args...))
}
