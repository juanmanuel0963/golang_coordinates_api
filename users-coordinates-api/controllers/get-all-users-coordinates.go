package controllers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sfox/v1/users-coordinates-api/files"
	"sfox/v1/users-coordinates-api/logger"

	"github.com/gin-gonic/gin"
)

// Struct to define error messages
type GetAllErrorMessages struct {
	ErrorOpenningInputFile                    string
	ErrorGettingUsersCoordinatesFromInputFile string
}

// Variable to hold error messages
var getAllErrorMessages = GetAllErrorMessages{
	ErrorOpenningInputFile:                    "PK_CONTROLLERS_FUNC_GETALL_ERROR_OPENING_INPUT_FILE",
	ErrorGettingUsersCoordinatesFromInputFile: "PK_CONTROLLERS_FUNC_GETALL_ERROR_GETTING_USERSCOORDINATES_FROM_INPUT_FILE",
}

// Mock Get All UserCoordinates endpoint
func GetAllUsersCoordinates(c *gin.Context) {

	//Load UserCoordinates from flat file
	dataFolder := "data"
	dataFileName := "customers.txt"
	dataFilePath := fmt.Sprintf("%s/%s", dataFolder, dataFileName)

	//Open the data file
	file, err := os.Open(dataFilePath)
	if err != nil {
		//Log the error in "logs/today.log" for further review
		errMsg := fmt.Sprintf("%s. %s.", getAllErrorMessages.ErrorOpenningInputFile, err.Error())
		logger.LogError(errMsg)
		//Send error response to client
		c.JSON(http.StatusBadRequest, gin.H{getAllErrorMessages.ErrorOpenningInputFile: err.Error()})
		return
	}

	defer file.Close()

	//Save the file content into a scanner
	scanner := bufio.NewScanner(file)

	// Get the users coordinates from a flat file
	outputUserCoordinates, err := files.GetUsersCoordinatesFromFile(scanner)

	// If there is an error getting the users coordinates from the flat file
	if err != nil {
		//Log the error in "logs/today.log" for further review
		errMsg := fmt.Sprintf("%s. %s.", getAllErrorMessages.ErrorGettingUsersCoordinatesFromInputFile, err.Error())
		logger.LogError(errMsg)
		//Send error response to client
		c.JSON(http.StatusInternalServerError, gin.H{getAllErrorMessages.ErrorGettingUsersCoordinatesFromInputFile: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": statusMessages.OK,
		"count":  len(outputUserCoordinates),
		"items":  outputUserCoordinates,
	})
}
