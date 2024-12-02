package day3

import (
	"os"
	"testing"
)

func TestGearRatios(t *testing.T) {
	testCases := []struct {
		name          string
		fileContent   string
		expectedP1    int
		expectedP2    int
		expectedError error
	}{
		{
			name: "ExampleTest",
			fileContent: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
			expectedP1:    4361,
			expectedP2:    467835,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a temporary file and write the content
			tmpFile, err := os.CreateTemp("", "gearRatios.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer func(name string) {
				err := os.Remove(name)
				if err != nil {

				}
			}(tmpFile.Name())

			_, err = tmpFile.WriteString(tc.fileContent)
			if err != nil {
				t.Fatal(err)
			}
			err = tmpFile.Close()
			if err != nil {
				return
			}

			// Call the function with the temporary file
			actualP1, actualP2 := GearRatios(tmpFile.Name())

			// Check the results
			if actualP1 != tc.expectedP1 {
				t.Errorf("Expected P1: %d, Got: %d", tc.expectedP1, actualP1)
			}

			if actualP2 != tc.expectedP2 {
				t.Errorf("Expected P2: %d, Got: %d", tc.expectedP2, actualP2)
			}

		})
	}
}
