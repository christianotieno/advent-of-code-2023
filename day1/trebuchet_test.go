package day1

import (
	"os"
	"testing"
)

func TestTrebuchet(t *testing.T) {
	t.Run("Scenario 1", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "testFile1.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer func(name string) {
			err := os.Remove(name)
			if err != nil {
				t.Fatal(err)
			}
		}(tempFile.Name())

		_, err = tempFile.WriteString("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet")
		if err != nil {
			t.Fatal(err)
		}
		err = tempFile.Close()
		if err != nil {
			t.Fatal(err)
		}

		p1Result, p2Result, err := Trebuchet(tempFile.Name())

		if err != nil {
			t.Fatal(err)
		}

		expectedP1Result := 142
		expectedP2Result := 142

		if p1Result != expectedP1Result {
			t.Errorf("Trebuchet returned p1: %d, expected p1: %d", p1Result, expectedP1Result)
		}

		if p2Result != expectedP2Result {
			t.Errorf("Trebuchet returned p2: %d, expected p2: %d", p2Result, expectedP2Result)
		}
	})

	t.Run("Scenario 2", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "testFile2.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer func(name string) {
			err := os.Remove(name)
			if err != nil {
				t.Fatal(err)
			}
		}(tempFile.Name())

		_, err = tempFile.WriteString("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")
		if err != nil {
			t.Fatal(err)
		}
		err = tempFile.Close()
		if err != nil {
			t.Fatal(err)
		}

		p1Result, p2Result, err := Trebuchet(tempFile.Name())

		if err != nil {
			t.Fatal(err)
		}

		expectedP1Result := 209
		expectedP2Result := 281

		if p1Result != expectedP1Result {
			t.Errorf("Trebuchet returned p1: %d, expected p1: %d", p1Result, expectedP1Result)
		}

		if p2Result != expectedP2Result {
			t.Errorf("Trebuchet returned p2: %d, expected p2: %d", p2Result, expectedP2Result)
		}
	})
}
