package day6

import "testing"

func TestWaitForIt(t *testing.T) {
	got := WaitForIt("waitForItTestFile.txt")
	expected := 352

	if expected != got {
		t.Errorf("Expected: %d, got: %d", expected, got)
	}
}
