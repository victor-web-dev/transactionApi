package util

import "testing"

func TestDateTime(t *testing.T) {
	time := DateTime()

	if time.Year() != 2026 {
		t.Errorf("Year should be 2026, got %d", time.Year())
	}
}
