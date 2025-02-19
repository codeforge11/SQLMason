package app

import (
	"log"
	"os"
)

var (
	errorLog *os.File
	logger   *log.Logger
)

func init() {
	var err error
	errorLog, err = os.OpenFile("error_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Failed to open error log file: %s", err)
	}

	logger = log.New(errorLog, "", log.Ldate|log.Ltime)
}

func logError(err error) {
	logger.Printf("|-| %s", err)
}
