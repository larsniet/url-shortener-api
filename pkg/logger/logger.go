package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

func InitLogger() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(msg string, args ...interface{}) {
	infoLogger.Output(2, fmt.Sprintf(msg, args...))
}

func Warn(msg string, args ...interface{}) {
	warnLogger.Output(2, fmt.Sprintf(msg, args...))
}

func Error(msg string, args ...interface{}) {
	errorLogger.Output(2, fmt.Sprintf(msg, args...))
}
