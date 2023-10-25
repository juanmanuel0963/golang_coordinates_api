package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Configuration is a struct that represents the configuration constants defined in the input "filename" variable
type Configuration struct {
	Tolerance               float64 `yaml:"tolerance"`
	EarthRadius             float64 `yaml:"earthRadius"`
	Port                    int     `yaml:"port"`
	ReferencePointLatitude  float64 `yaml:"referencePointLatitude"`
	ReferencePointLongitude float64 `yaml:"referencePointLongitude"`
	MaxDistance             float64 `yaml:"maxDistance"`
}

// LoadConfig loads the configuration from a YAML file
func LoadConfig(filename string) (*Configuration, error) {

	config := &Configuration{}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return config, nil
}
