package tests

import (
	"bufio"
	"os"
	"reflect"
	"sfox/v1/users-coordinates-api/files"
	"sfox/v1/users-coordinates-api/models"
	"testing"
)

func TestProcessUploadedFile(t *testing.T) {

	// Create a test file
	testFileName := "test.txt"
	testFile, err := os.Create(testFileName)
	if err != nil {
		t.Fatal(err)
	}

	// Write test data to the file
	testData := `{"user_id": 2, "name": "John Doe", "latitude": "42.123", "longitude": "-71.456"}`
	_, err = testFile.WriteString(testData)
	if err != nil {
		t.Fatal(err)
	}

	// Close the file
	testFile.Close()

	//Open the file
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	//Save the file content into a scanner
	scanner := bufio.NewScanner(file)

	// Get the users coordinates from a flat file
	locations, err := files.GetUsersCoordinatesFromFile(scanner)

	if err != nil {
		t.Errorf("Error processing uploaded file: %v", err)
	}

	// Check the locations slice for expected values
	expected := []models.UserCoordinates{
		{Id: 2, Name: "John Doe", Latitude: 42.123, Longitude: -71.456},
	}
	if !reflect.DeepEqual(locations, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, locations)
	}
}
