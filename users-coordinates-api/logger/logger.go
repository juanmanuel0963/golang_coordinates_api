// Package for logging error messages on external files
package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Save error messages to a log file
func LogError(errorMessage string) {

	// Create a "logs" subfolder if it doesn't exist
	logsFolder := "logs"
	if _, err := os.Stat(logsFolder); os.IsNotExist(err) {
		os.Mkdir(logsFolder, os.ModePerm)
	}

	// Get the current date and time
	currentTime := time.Now()
	// Format the date in YYYY-MM-DD format
	date := currentTime.Format("2006-01-02")

	// Create or open the log file in the "logs" subfolder with the current date as the filename
	logFileName := fmt.Sprintf("%s/%s.log", logsFolder, date)
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// If there is an error opening the log file, print the error and exit
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer logFile.Close()

	// Create a new logger that writes to the log file
	logger := log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Log the error message with the current date and time
	logger.Printf("%s - %s", currentTime.Format("2006-01-02 15:04:05"), errorMessage)
}
