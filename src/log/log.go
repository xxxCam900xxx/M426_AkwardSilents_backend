package log

import (
	"log"
	"os"
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func init() {

	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime)
	warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs
func Info(message string) {
	infoLogger.Println(message)
}

// Warning logs
func Warning(message string) {
	warningLogger.Println(message)
}

// Error logs
func Error(message string) {
	errorLogger.Println(message)
}
