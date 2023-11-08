package services

import (
	"encoding/json"
	"log"
)

// Check an error and show the log in the console
func CheckError(err error) bool {
	if err != nil {
		log.Printf("Error: %v\n", err)
		return true
	}
	return false
}

// Convert any data to json
func ToJSON(obj any) []byte {
	data, err := json.Marshal(obj)

	CheckError(err)

	return data
}
