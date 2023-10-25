package tests

import (
	"fmt"
	"os"
	"sfox/v1/users-coordinates-api/config"
	"testing"
)

// TestLoadValidConfig tests loading a valid configuration file.
func TestLoadValidConfig(t *testing.T) {

	expectedConfig := &config.Configuration{
		Tolerance:               0.01,
		EarthRadius:             6371,
		Port:                    8080,
		MaxDistance:             100,
		ReferencePointLatitude:  53.339428,
		ReferencePointLongitude: -6.257664,
	}

	config, err := config.LoadConfig("../config/config.yaml")

	if err != nil {
		t.Errorf("Error loading valid config: %v", err)
	}

	if config == nil {
		t.Error("Expected non-nil configuration, but got nil")
	}

	if config == nil {
		t.Error("Expected non-nil configuration, but got nil")
	} else if *config != *expectedConfig {
		t.Errorf("Expected config does not match loaded config.\nExpected: %v\nActual: %v", expectedConfig, config)
	}

	if *config != *expectedConfig {
		t.Errorf("Expected config does not match loaded config.\nExpected: %v\nActual: %v", expectedConfig, config)
	}
}

// TestLoadNonExistentConfigFile tests loading a non-existent configuration file.
func TestLoadNonExistentConfigFile(t *testing.T) {

	_, err := config.LoadConfig("nonexistent.yaml")

	if err == nil {
		t.Error("Expected an error when loading a non-existent file, but got nil")
	}
}

// TestLoadInvalidConfigFile tests loading an invalid configuration file (e.g., missing required fields).
func TestLoadInvalidConfigFile(t *testing.T) {

	// Create an invalid config.yaml with strings instead of int and float values
	invalidConfigYAML :=
		`tolerance: "0.01"
earthRadius: "6371"
port: "8080"
maxDistance: "100"
referencePointLatitude: "53.339428"
referencePointLongitude: "-6.257664"`

	err := createTempConfigFile(invalidConfigYAML, "invalid_config.yaml")
	if err != nil {
		t.Errorf("Error creating temporary invalid config file: %v", err)
		return
	}

	config, loadErr := config.LoadConfig("invalid_config.yaml")

	fmt.Println(config)

	if loadErr == nil {
		t.Error("Expected an error when loading an invalid file, but got nil")
	}
}

func createTempConfigFile(content, filename string) error {

	// Create a temporary config file for testing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, writeErr := file.WriteString(content)
	if writeErr != nil {
		return writeErr
	}
	return nil
}
