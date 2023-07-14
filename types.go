package ocsval

type Config struct {
	FileMetadata    FileMetadata    `json:"fileMetadata"`
	Columns         []Column        `json:"columns"`
	FileConstraints FileConstraints `json:"fileConstraints"`
}

type FileMetadata struct {
	Delimiter string `json:"delimiter"`
	Encoding  string `json:"encoding"`
	HasHeader bool   `json:"hasHeader"`
}

type Column struct {
	Name        string      `json:"name"`
	Constraints Constraints `json:"constraints"`
}

type Constraints struct {
	Type        string  `json:"type"`
	Unique      bool    `json:"unique,omitempty"`
	MustPresent bool    `json:"mustPresent,omitempty"`
	Min         *int    `json:"min,omitempty"`
	Max         *int    `json:"max,omitempty"`
	Pattern     *string `json:"pattern,omitempty"`
	MaxLength   *int    `json:"maxLength,omitempty"`
	Is          *string `json:"is,omitempty"`
}

type FileConstraints struct {
	MaxRows int `json:"maxRows"`
	MaxSize int `json:"maxSize"`
}

type ValidationError struct {
	Row     int
	Column  string
	Message string
}

type EmptyCSVErr struct{}

func (e *EmptyCSVErr) Error() string {
	return "CSV file is empty"
}

// Alias for a row in the CSV
type Row []string
