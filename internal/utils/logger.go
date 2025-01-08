package utils

import (
	"log"
	"os"
)

func InitLogger() *log.Logger {
	file, err := os.Create("goloadtest.log")
	if err != nil {
		log.Fatalf("Could not create log file: %v", err)
	}
	return log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
