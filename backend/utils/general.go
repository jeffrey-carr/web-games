package utils

import (
	"binoku/entities"
	"os"

	"gopkg.in/yaml.v3"
)

// IsZero returns true if the item is the zero-value of its type
func IsZero[T comparable](item T) bool {
	var zero T
	return item == zero
}

// ReadConfig reads the provided configuration file
func ReadConfig(filename string) (entities.Config, error) {
	// Load file
	file, err := os.ReadFile(filename)
	if err != nil {
		return entities.Config{}, err
	}

	// Read file into entity
	var config entities.Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return entities.Config{}, err
	}

	return config, nil
}
