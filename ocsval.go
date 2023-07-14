package ocsval

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

// Validate function validates the given CSV based on the provided schema and returns a slice of ValidationError.
func Validate(reader io.Reader, configContent string) ([]ValidationError, error) {
	// Load and validate the configuration
	config, err := loadConfigFromContent(configContent)
	if err != nil {
		return nil, err
	}

	// Create a CSV reader
	csvReader := csv.NewReader(reader)
	csvReader.Comma = rune(config.FileMetadata.Delimiter[0])

	// Skip the header if present
	if config.FileMetadata.HasHeader {
		if _, err := csvReader.Read(); err != nil {
			return nil, err
		}
	}

	// Validate the CSV rows
	return validateCSVRows(csvReader, config.Columns)
}

// loadConfigFromContent parses the JSON configuration content and validates it against the schema.
func loadConfigFromContent(configContent string) (*Config, error) {
	var config Config
	if err := json.Unmarshal([]byte(configContent), &config); err != nil {
		return nil, err
	}

	if err := validateConfigContent(configContent); err != nil {
		return nil, err
	}

	return &config, nil
}

// validateConfigContent validates the configuration content against the JSON schema.
func validateConfigContent(configContent string) error {
	schemaLoader := gojsonschema.NewStringLoader(configSchema)
	configLoader := gojsonschema.NewStringLoader(configContent)

	result, err := gojsonschema.Validate(schemaLoader, configLoader)
	if err != nil {
		return err
	}

	if !result.Valid() {
		var validationErrors []string
		for _, err := range result.Errors() {
			validationErrors = append(validationErrors, err.String())
		}
		return errors.New("config validation failed: " + strings.Join(validationErrors, ", "))
	}

	return nil
}

// validateCSVRows reads and validates each row from the CSV reader.
func validateCSVRows(csvReader *csv.Reader, columns []Column) ([]ValidationError, error) {
	var validationErrors []ValidationError
	rowIndex := 0

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		rowIndex++
		rowErrors := validateRow(Row(row), columns, rowIndex)
		validationErrors = append(validationErrors, rowErrors...)
	}

	return validationErrors, nil
}

// validateRow validates a single row against the given columns and returns a slice of ValidationError.
func validateRow(row Row, columns []Column, rowIndex int) []ValidationError {
	var validationErrors []ValidationError

	for colIndex, col := range columns {
		if colIndex >= len(row) {
			validationErrors = append(validationErrors, ValidationError{
				Row:     rowIndex,
				Column:  col.Name,
				Message: "missing value",
			})
			continue
		}

		value := row[colIndex]
		if err := validateValue(value, col.Constraints); err != nil {
			validationErrors = append(validationErrors, ValidationError{
				Row:     rowIndex,
				Column:  col.Name,
				Message: err.Error(),
			})
		}
	}

	return validationErrors
}

// validateValue validates a single value against the given constraints.
func validateValue(value string, constraints Constraints) error {
	if constraints.MustPresent && value == "" {
		return errors.New("value is required")
	}

	switch constraints.Type {
	case "int":
		if _, err := strconv.Atoi(value); err != nil {
			return errors.New("value is not an integer")
		}
	case "string":
		// No specific validation for string type
	default:
		return errors.New("unsupported type")
	}

	if constraints.Unique {
		// TODO: Implement unique constraint validation
	}

	if constraints.Min != nil {
		val, err := strconv.Atoi(value)
		if err != nil {
			return errors.New("value is not an integer")
		}
		if val < *constraints.Min {
			return errors.New("value is below minimum")
		}
	}

	if constraints.Max != nil {
		val, err := strconv.Atoi(value)
		if err != nil {
			return errors.New("value is not an integer")
		}
		if val > *constraints.Max {
			return errors.New("value is above maximum")
		}
	}

	if constraints.Pattern != nil {
		matched, err := regexp.MatchString(*constraints.Pattern, value)
		if err != nil {
			return err
		}
		if !matched {
			return errors.New("value does not match pattern")
		}
	}

	if constraints.MaxLength != nil && len(value) > *constraints.MaxLength {
		return errors.New("value exceeds maximum length")
	}

	if constraints.Is != nil && value != *constraints.Is {
		return errors.New("value does not match the required value")
	}

	return nil
}
