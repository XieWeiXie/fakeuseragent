package parse

import (
	"testing"
)

func TestStatistics(t *testing.T) {
	if len(Statistics()) == 0 {
		t.Errorf("Statistics() should not be nil")
	}
}
