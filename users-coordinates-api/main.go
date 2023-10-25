package main

import (
	"fmt"
	"log"
	"sfox/v1/users-coordinates-api/config"
	"sfox/v1/users-coordinates-api/controllers"
	"sfox/v1/users-coordinates-api/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Declare a package-level variable to hold the configuration
var configObject *config.Configuration

// ErrorMessages is a struct to define error messages
type MainErrorMessages struct {
	ErrorStartingGinServer string
}

// Global variable to hold error messages
var mainErrorMessages = MainErrorMessages{
	ErrorStartingGinServer: "PK_MAIN_FC_MAIN_ERROR_STARTING_GIN_SERVER",
}

func init() {

	// Load the configuration from the YAML file
	config, err := config.LoadConfig("config/config.yaml")

	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Assign the loaded configuration to the package-level variable
	configObject = config
}

func main() {

	//Create a new default Gin HTTP Server
	server := gin.Default()

	// Define CORS middleware options - Just for localhost testing
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}                                       // Change this to the specific origins allowed
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // Allow the HTTP methods need
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type"}

	// Use the CORS middleware
	server.Use(cors.New(corsConfig))

	// userscoordinates routes
	userslocation := server.Group("/userscoordinates")
	{
		// Create a closure to pass the 'configObject' to the FilterUsersCoordinates function
		filterUsersCoordinatesHandler := func(c *gin.Context) {
			controllers.FilterUsersCoordinates(c, configObject)
		}

		// Set up route for filtering the coordinates
		userslocation.POST("/filter", filterUsersCoordinatesHandler)

		// Other mocked endpoints
		userslocation.POST("/", controllers.CreateUserCoordinates)
		userslocation.GET("/", controllers.GetAllUsersCoordinates)
	}

	//Raise the Gin HTTP Server
	err := server.Run(fmt.Sprintf(":%v", configObject.Port))

	// If there is an error raising the http server
	if err != nil {
		//Log the error in "logs/today.log" for further review
		errMsg := fmt.Sprintf("%s. %s.", mainErrorMessages.ErrorStartingGinServer, err.Error())
		logger.LogError(errMsg)
	}
}
