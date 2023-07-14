package ocsval

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

const configSchema = `
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "title": "Config",
  "type": "object",
  "properties": {
    "fileMetadata": {
      "type": "object",
      "properties": {
        "delimiter": {"type": "string", "minLength": 1},
        "encoding": {"type": "string", "minLength": 1},
        "hasHeader": {"type": "boolean"}
      },
      "required": ["delimiter", "encoding", "hasHeader"]
    },
    "columns": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "name": {"type": "string", "minLength": 1},
          "constraints": {
            "type": "object",
            "properties": {
              "type": {"type": "string", "enum": ["string", "int", "float"]},
              "unique": {"type": "boolean"},
              "required": {"type": "boolean"},
              "min": {"type": "integer"},
              "max": {"type": "integer"},
              "pattern": {"type": "string"},
              "maxLength": {"type": "integer"},
              "is": {"type": "string"}
            },
            "required": ["type"]
          }
        },
        "required": ["name", "constraints"]
      }
    },
    "fileConstraints": {
      "type": "object",
      "properties": {
        "maxRows": {"type": "integer", "minimum": 0},
        "maxSize": {"type": "integer", "minimum": 0}
      },
      "required": ["maxRows", "maxSize"]
    }
  },
  "required": ["fileMetadata", "columns", "fileConstraints"]
}
`

// LoadConfig reads the JSON configuration file, validates it against the schema,
// and returns a Config struct. It takes the path to the configuration file as input.
func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	if err := validateConfig(configPath); err != nil {
		return nil, err
	}

	return &config, nil
}

func validateConfig(configPath string) error {
	schemaLoader := gojsonschema.NewStringLoader(configSchema)
	configLoader := gojsonschema.NewReferenceLoader("file://" + configPath)

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
