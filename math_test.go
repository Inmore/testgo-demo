package math

import "testing"

func TestAdd(t *testing.T) {
	got := Add(4, 11)
	if got != 15 {
		t.Fatalf("want 15, got %d", got)
	}
}

func TestAdd_Table(t *testing.T) {
	cases := []struct {
		name     string
		a, b, ok int
	}{
		{"zero", 0, 0, 0},
		{"positive", 4, 11, 15},
		{"negative", -1, -3, -4},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Add(c.a, c.b); got != c.ok {
				t.Fatalf("%s: want %d, got %d", c.name, c.a, c.b)
			}
		})
	}

}
