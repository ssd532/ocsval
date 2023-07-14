# CSV Validator 

This library validates CSV files based on a defined schema. It ensures the data meets specified rules.

## Features

- Validate CSV columns and overall file.
- Support for common constraints and rules.

## Supported Rules

- **required**: Ensures the field is not empty.
- **type**: Validates the data type (integer, string).
- **unique**: Ensures all values in the field are unique.
- **min**: Validates that numeric values are above a minimum.
- **max**: Validates that numeric values are below a maximum.
- **pattern**: Validates strings against a regex pattern.
- **is**: Ensures the field matches an exact value.

## Schema Example

```json
{
  "fileMetadata": {
    "delimiter": ",",
    "encoding": "UTF-8",
    "hasHeader": true
  },
  "columns": [
    {
      "name": "id",
      "constraints": {
        "type": "integer",
        "unique": true,
        "required": true
      }
    },
    {
      "name": "name",
      "constraints": {
        "type": "string",
        "required": true,
        "maxLength": 255
      }
    },
    {
      "name": "age",
      "constraints": {
        "type": "integer",
        "min": 0,
        "max": 120
      }
    },
    {
      "name": "email",
      "constraints": {
        "type": "string",
        "pattern": "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$",
        "unique": true
      }
    }
  ],
  "fileConstraints": {
    "maxRows": 1000,
    "maxSize": 1048576
  }
}
```

### JSON Schema Sections

- **fileMetadata**: Contains basic properties about the CSV file like delimiter, encoding, and whether it has a header.
- **columns**: Defines the columns of the CSV file, including their names and constraints like type, uniqueness, and validation patterns.
- **fileConstraints**: Specifies overall file limits, such as the maximum number of rows and the maximum file size in bytes.
