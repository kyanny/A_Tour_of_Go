package sum

import "testing"

func TestSum(t *testing.T) {
	var n int
	n = sum(3, 5)
	if n != 8 {
		t.Error("Expected 8, got ", n)
	}
}