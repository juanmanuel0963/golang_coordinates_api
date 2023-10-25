package tests

import (
	"os"
	"sfox/v1/users-coordinates-api/logger"
	"testing"
	"time"
)

func TestLogError(t *testing.T) {
	// Define a temporary test directory
	testDir := "test_logs"
	defer os.RemoveAll(testDir)

	// Test case 1: Successful log
	errorMessage := "This is a test error message."
	logger.LogError(errorMessage)

	// Check if the log file was created
	logFilePath := "logs"
	_, err := os.Stat(logFilePath)
	if err != nil {
		t.Errorf("Expected log folder to be created, but got error: %v", err)
	}
	/*
		// Test case 2: Error creating log file
		// Create a read-only directory to simulate a permission error
		err = os.MkdirAll(testDir, 0444)
		if err != nil {
			t.Fatal(err)
		}
		errorMessage = "Error message with log file creation failure."
		LogError(errorMessage)

		// Check if an error occurred (since log file creation failed)
		_, err = os.Stat(logFilePath)
		if err == nil {
			t.Errorf("Expected an error creating the log file, but got no error")
		}
	*/

	// Test case 3: Error writing to log file
	// Create a file with read-only permission to simulate a write error
	err = os.MkdirAll(logFilePath, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chmod(logFilePath, 0444)
	if err != nil {
		t.Fatal(err)
	}
	errorMessage = "Error message with log file write failure."
	logger.LogError(errorMessage)

	// Check if an error occurred (since writing to the log file failed)
	logFile := logFilePath + "/" + getLogFileName()
	_, err = os.Stat(logFile)
	if err == nil {
		t.Errorf("Expected an error writing to the log file, but got no error")
	}

}

func getLogFileName() string {
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	return "logs/" + date + ".log"
}
