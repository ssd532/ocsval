package ocsval

import (
	"reflect"
	"testing"
)

func TestReadCSV(t *testing.T) {
	data, err := ReadCSV("test.csv")
	if err != nil {
		t.Errorf("Returned error: %s", err)
	}

	expected := [][]string{
		{"Name", "Age", "Email"},
		{"John Doe", "30", "johndoe@example.com"},
		// and so forth, for all the rows you expect in your test.csv
	}

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("ReadCSV = %v, want %v for given CSV file", data, expected)
	}
}
