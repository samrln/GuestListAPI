package logger

import (
    "log"
    "os"
)

var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

func Info(format string, args ...interface{}) {
    logger.Printf(format, args...)
}

func Error(format string, args ...interface{}) {
    logger.Printf("ERROR: "+format, args...)
}
