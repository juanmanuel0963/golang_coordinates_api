package files

import (
	"bufio"
	"encoding/json"
	"fmt"
	"sfox/v1/users-coordinates-api/logger"
	"sfox/v1/users-coordinates-api/models"
)

// Struct to define error messages
type FuncGetErrorMessages struct {
	ErrorMarshalingUserCoordinates string
}

// Variable to hold error messages
var funcGetErrorMessages = FuncGetErrorMessages{
	ErrorMarshalingUserCoordinates: "PK_FILES_FUNC_GET_ERROR_MARSHALING_USERCOORDINATES",
}

// Get the users coordinates from a flat file
func GetUsersCoordinatesFromFile(scanner *bufio.Scanner) ([]models.UserCoordinates, error) {

	//Variable for storing whole user locations
	var locations []models.UserCoordinates

	//Iterate each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		//Marshalls the line to a UserLocation struct
		var location models.UserCoordinates
		if err := json.Unmarshal([]byte(line), &location); err != nil {
			//Log the error in "logs/today.log" for further review
			errMsg := fmt.Sprintf("%s. %s.", funcGetErrorMessages.ErrorMarshalingUserCoordinates, err.Error())
			logger.LogError(errMsg)

			continue
		}
		//Adds the UserLocation to a slice list
		locations = append(locations, location)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	//Return the list of UserLocations
	return locations, nil
}
