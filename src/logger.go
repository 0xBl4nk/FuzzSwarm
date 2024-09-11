package src

import "log"

func LogFatal(format string, v ...interface{}) {
    log.Fatalf(format, v...)
}

