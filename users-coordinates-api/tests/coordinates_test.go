package tests

import (
	"log"
	"math"
	"sfox/v1/users-coordinates-api/config"
	"sfox/v1/users-coordinates-api/coordinates"
	"sfox/v1/users-coordinates-api/models"
	"testing"
)

// Declare a package-level variable to hold the configuration
var configObject *config.Configuration

func init() {

	// Load the configuration from the YAML file
	config, err := config.LoadConfig("../config/config.yaml")

	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Assign the loaded configuration to the package-level variable
	configObject = config
}

func TestDegreesToRadians(t *testing.T) {
	testCases := []struct {
		degrees         float64
		expectedRadians float64
	}{
		{0.0, 0.0},
		{180.0, math.Pi},
		{-90.0, -math.Pi / 2},
		{45.0, math.Pi / 4},
	}

	for _, tc := range testCases {
		actual := coordinates.DegreesToRadians(tc.degrees)
		if actual != tc.expectedRadians {
			t.Errorf("degreesToRadians(%f) expected: %f, got: %f", tc.degrees, tc.expectedRadians, actual)
		}
	}
}

func TestHaversineDistance(t *testing.T) {

	newYork := models.UserCoordinates{Latitude: 40.7128, Longitude: -74.0060}
	losAngeles := models.UserCoordinates{Latitude: 34.0522, Longitude: -118.2437}

	expectedDistance := 3935.74 // Expected distance between New York and Los Angeles in km

	actualDistance := coordinates.HaversineDistance(newYork, losAngeles, configObject)

	// Allow for a small tolerance (e.g., 0.01) due to float precision
	tolerance := 0.01
	if math.Abs(actualDistance-expectedDistance) > tolerance {
		t.Errorf("haversineDistance() expected: %f, got: %f", expectedDistance, actualDistance)
	}
}

func TestFilterCoordinatesByDistance(t *testing.T) {

	// Create a reference point
	referencePoint := models.UserCoordinates{
		Id:        1,
		Name:      "Reference",
		Latitude:  40.0,
		Longitude: -75.0,
	}

	// Create a set of user coordinates
	userCoordinates := []models.UserCoordinates{
		{Id: 2, Name: "User1", Latitude: 40.1, Longitude: -75.1},
		{Id: 3, Name: "User2", Latitude: 40.2, Longitude: -75.2},
		{Id: 4, Name: "User3", Latitude: 40.3, Longitude: -75.3},
		{Id: 5, Name: "User4", Latitude: 40.4, Longitude: -75.4},
	}

	// Define the maximum distance
	maxDistance := 30.0

	// Call the function to filter coordinates
	filteredCoordinates := coordinates.FilterCoordinatesByDistance(referencePoint, userCoordinates, maxDistance, configObject)

	// Verify the number of filtered coordinates
	if len(filteredCoordinates) != 2 {
		t.Errorf("Expected 2 filtered coordinates, but got %d", len(filteredCoordinates))
	}

	// Allow for a small tolerance (e.g., 0.01) due to float precision
	tolerance := 0.01

	// Verify the filtered coordinates' distances
	expectedDistances := []float64{coordinates.HaversineDistance(referencePoint, userCoordinates[0], configObject), coordinates.HaversineDistance(referencePoint, userCoordinates[1], configObject)}
	for i, coord := range filteredCoordinates {
		if math.Abs(coord.Distance-expectedDistances[i]) > tolerance {
			t.Errorf("Distance for %s is incorrect. Expected: %f, Got: %f", coord.Name, expectedDistances[i], coord.Distance)
		}
	}
}
