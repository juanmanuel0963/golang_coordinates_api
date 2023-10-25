package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"sfox/v1/users-coordinates-api/logger"
	"sfox/v1/users-coordinates-api/models"

	"github.com/gin-gonic/gin"
)

// Struct to define error messages
type FuncCreateErrorMessages struct {
	ErrorBindingUserCoordinates string
}

// Variable to hold error messages
var funcCreateErrorMessages = FuncCreateErrorMessages{
	ErrorBindingUserCoordinates: "PK_CONTROLLERS_FUNC_CREATE_ERROR_BINDING_USERCOORDINATES",
}

// Mock Create UserCoordinates endpoint
func CreateUserCoordinates(c *gin.Context) {

	var bodyUserCoordinates models.UserCoordinates

	// Bind incoming JSON to user struct
	if err := c.ShouldBindJSON(&bodyUserCoordinates); err != nil {
		//Log the error in "logs/today.log" for further review
		errMsg := fmt.Sprintf("%s. %s.", funcCreateErrorMessages.ErrorBindingUserCoordinates, err.Error())
		logger.LogError(errMsg)
		//Send error response to client
		c.JSON(http.StatusBadRequest, gin.H{funcCreateErrorMessages.ErrorBindingUserCoordinates: err.Error()})
		return
	}

	//Assign random ID
	bodyUserCoordinates.Id = int32(rand.Intn(1000))

	// Return
	c.JSON(http.StatusCreated, gin.H{
		"status": statusMessages.Created,
		"item":   bodyUserCoordinates,
	})
}
