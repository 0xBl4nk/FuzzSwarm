package src

import (
    "log"
)

// LogFatal logs a fatal error and exits the program.
func LogFatal(format string, v ...interface{}) {
    log.Fatalf("[FATAL] "+format, v...)
}

// LogInfo logs informational messages.
func LogInfo(format string, v ...interface{}) {
    log.Printf("[INFO] "+format, v...)
}

// LogError logs error messages.
func LogError(format string, v ...interface{}) {
    log.Printf("[ERROR] "+format, v...)
}
