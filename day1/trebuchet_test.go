package day1

import (
	"os"
	"testing"
)

func TestTrebuchet(t *testing.T) {
	tempFile, err := os.CreateTemp("", "tesFile.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(tempFile.Name())

	// Write test data to the temporary file
	_, err = tempFile.WriteString("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet")
	if err != nil {
		t.Fatal(err)
	}
	err = tempFile.Close()
	if err != nil {
		return
	}

	result := Trebuchet(tempFile.Name())

	// Define the expected result based on the test data
	expectedResult := 142

	// Check if the result matches the expected result
	if result != expectedResult {
		t.Errorf("Trebuchet returned %d, expected %d", result, expectedResult)
	}
}
