package controllers

import (
	"bufio"
	"fmt"
	"mime/multipart"
	"net/http"
	"sfox/v1/users-coordinates-api/config"
	"sfox/v1/users-coordinates-api/coordinates"
	"sfox/v1/users-coordinates-api/files"
	"sfox/v1/users-coordinates-api/logger"
	"sfox/v1/users-coordinates-api/models"
	"strconv"

	"sort"

	"github.com/gin-gonic/gin"
)

// Struct to define error messages
type FuncFilterErrorMessages struct {
	ErrorBindingMaxDistance                   string
	ErrorCastingMaxDistance                   string
	ErrorBidingInputFile                      string
	ErrorOpeningInputFile                     string
	ErrorGettingUsersCoordinatesFromInputFile string
}

// Variable to hold error messages
var funcFilterErrorMessages = FuncFilterErrorMessages{
	ErrorBindingMaxDistance:                   "PK_CONTROLLERS_FUNC_FILTER_ERROR_BINDING_MAXDISTANCE",
	ErrorCastingMaxDistance:                   "PK_CONTROLLERS_FUNC_FILTER_ERROR_CASTING_MAXDISTANCE",
	ErrorBidingInputFile:                      "PK_CONTROLLERS_FUNC_FILTER_ERROR_BINDING_INPUT_FILE",
	ErrorOpeningInputFile:                     "PK_CONTROLLERS_FUNC_FILTER_ERROR_OPENING_INPUT_FILE",
	ErrorGettingUsersCoordinatesFromInputFile: "PK_CONTROLLERS_FUNC_FILTER_ERROR_GETTING_USERSCOORDINATES_FROM_INPUT_FILE",
}

// Package level struct to define success status messages
type StatusMessages struct {
	OK      string
	Created string
}

// Package level variable to hold success status messages
var statusMessages = StatusMessages{
	OK:      "OK",
	Created: "CREATED",
}

// Struct for binding the input file
type txtUploadInput struct {
	TxtFile *multipart.FileHeader `form:"file" binding:"required"`
}

// Struct for binding the input max distance
type maxDistanceInput struct {
	TextValue string `form:"maxdistance" binding:"required"`
}

// Filter a file of users' coordinates and return those within a maximum distance relative to a reference point
func FilterUsersCoordinates(c *gin.Context, config *config.Configuration) {

	// Step 1.---------------Get Input Distance---------------

	var inputMaxDistance maxDistanceInput
	if err := c.ShouldBind(&inputMaxDistance); err != nil {
		//Log the error in "logs/today.log" for further review
		errMsg := fmt.Sprintf("%s. %s.", funcFilterErrorMessages.ErrorBindingMaxDistance, err.Error())
		logger.LogError(errMsg)
		//Send error response to client
		c.JSON(http.StatusBadRequest, gin.H{funcFilterErrorMessages.ErrorBindingMaxDistance: err.Error()})
		return
	}

	//Convert the string containing the max distance to float64
	maxDistance, err := strconv.ParseFloat(inputMaxDistance.TextValue, 64)
	if err != nil {
		//Log the error in "logs/today.log" for further review
		errMsg := fmt.Sprintf("%s. %s.", funcFilterErrorMessages.ErrorCastingMaxDistance, err.Error())
		logger.LogError(errMsg)
		//Send error response to client
		c.JSON(http.StatusBadRequest, gin.H{funcFilterErrorMessages.ErrorCastingMaxDistance: err.Error()})
		return
	}

	// Step 2.---------------Get Input File---------------

	var inputFile txtUploadInput
	if err := c.ShouldBind(&inputFile); err != nil {
		//Log the error in "logs/today.log" for further review
		errMsg := fmt.Sprintf("%s. %s.", funcFilterErrorMessages.ErrorBidingInputFile, err.Error())
		logger.LogError(errMsg)
		//Send error response to client
		c.JSON(http.StatusBadRequest, gin.H{funcFilterErrorMessages.ErrorBidingInputFile: err.Error()})
		return
	}

	//Open de file
	file, err := inputFile.TxtFile.Open()

	// If there is an error opening the flat file
	if err != nil {
		//Log the error in "logs/today.log" for further review
		errMsg := fmt.Sprintf("%s. %s.", funcFilterErrorMessages.ErrorOpeningInputFile, err.Error())
		logger.LogError(errMsg)
		//Send error response to client
		c.JSON(http.StatusBadRequest, gin.H{funcFilterErrorMessages.ErrorOpeningInputFile: err.Error()})
		return
	}

	defer file.Close()

	//Save the file content into a scanner
	scanner := bufio.NewScanner(file)

	// Step 3.---------------Get the users coordinates from flat file---------------

	listUserCoordinates, err := files.GetUsersCoordinatesFromFile(scanner)

	// If there is an error getting the users coordinates from the flat file
	if err != nil {
		//Log the error in "logs/today.log" for further review
		errMsg := fmt.Sprintf("%s. %s.", funcFilterErrorMessages.ErrorGettingUsersCoordinatesFromInputFile, err.Error())
		logger.LogError(errMsg)
		//Send error response to client
		c.JSON(http.StatusInternalServerError, gin.H{funcFilterErrorMessages.ErrorGettingUsersCoordinatesFromInputFile: err.Error()})
		return
	}

	// Step 4.---------------Sort coordinates by UserID---------------

	// Sort locations by UserID in ascending order
	sort.Sort(models.ByUserID(listUserCoordinates))

	// Step 5.---------------Filter coordinates by distance to a reference point---------------

	//Get the reference point
	referencePoint := models.UserCoordinates{Latitude: config.ReferencePointLatitude, Longitude: config.ReferencePointLongitude}

	// Filter coordinates by distance to a reference point
	outputUserCoordinates := coordinates.FilterCoordinatesByDistance(referencePoint, listUserCoordinates, maxDistance, config)

	//Step 6.------------------ Respond with the filtered locations as JSON---------------

	c.JSON(http.StatusOK, gin.H{
		"status": statusMessages.OK,
		"count":  len(outputUserCoordinates),
		"items":  outputUserCoordinates,
	})

}
