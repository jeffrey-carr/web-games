package utils

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

// NewUUIDString creates a new UUID in string form
func NewUUIDString() string {
	return uuid.New().String()
}

// ReadJSONFile reads in a JSON file and unmarshals it as a type
func ReadJSONFile[T any](filepath string) (T, error) {
	var result T

	data, err := os.ReadFile(filepath)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
