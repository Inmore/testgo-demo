package math

import "testing"

func TestAdd(t *testing.T) {
	got := Add(4, 11)
	if got != 15 {
		t.Fatalf("want 15, got %d", got)
	}
}
