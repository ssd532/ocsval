package main

import (
	"fmt"
	"os"

	"ocsval"
)

func main() {
	// Read the configuration schema from config_schema.json
	configContent, err := os.ReadFile("config_schema.json")
	if err != nil {
		fmt.Println("Error reading config schema:", err)
		return
	}

	// Read the CSV content from sample.csv
	csvFile, err := os.Open("sample.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Validate the CSV content
	validationErrors, err := ocsval.Validate(csvFile, string(configContent))
	if err != nil {
		fmt.Println("Error validating CSV:", err)
		return
	}

	// Print validation errors if any
	if len(validationErrors) > 0 {
		fmt.Println("Validation Errors:")
		for _, ve := range validationErrors {
			fmt.Printf("Row %d, Column %s: %s\n", ve.Row, ve.Column, ve.Message)
		}
	} else {
		fmt.Println("CSV is valid.")
	}
}
