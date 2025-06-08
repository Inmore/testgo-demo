package math

import (
	"testing"
)

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

func TestLookup(t *testing.T) {
	cases := []struct {
		name     string
		a, b, ok int
	}{
		{"first", 1, 1, 2},
		{"second", 2, 5, 7},
		{"third", 3, 5, 8},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			if got := Add(c.a, c.b); got != c.ok {
				t.Fatalf("%s: got %d, want %d", c.name, got, c.ok)
			}
		})
	}
}

func TestSolveQuadraticEquation(t *testing.T) {
	cases := []struct {
		name, m         string
		a, b, c, x1, x2 float64
	}{
		{
			name: "has solutions",
			a:    2,
			b:    1,
			c:    -6,
			x1:   1.5,
			x2:   -2,
			m:    "",
		},
		{
			name: "hasn't solutions",
			a:    2,
			b:    -3,
			c:    7,
			x1:   0,
			x2:   0,
			m:    "решений нет",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got_x1, got_x2, got_m := SolveQuadraticEquation(c.a, c.b, c.c)
			if got_x1 != c.x1 || got_x2 != c.x2 || got_m != c.m {
				t.Fatalf("got x1=%f, x2=%f, m=%s, but want x1=%f, x2=%f, m=%s", got_x1, got_x2, got_m, c.x1, c.x2, c.m)
			}
		})
	}
}

func TestGreatestCommonDivisor(t *testing.T) {
	cases := []struct {
		name      string
		x, y, res uint
	}{
		{
			name: "regular case",
			x:    35,
			y:    21,
			res:  7,
		},
		{
			name: "0 case",
			x:    0,
			y:    0,
			res:  0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := GreatestCommonDivisor(c.x, c.y)
			if got != c.res {
				t.Fatalf("got %d, but want %d", got, c.res)
			}
		})
	}
}

func TestRocketLunch(t *testing.T) {
	cases := []struct {
		name, res string
		v         float64
	}{
		{
			name: "Скорость меньше 0",
			v:    -1,
			res:  "Скорость не может быть меньше 0",
		},
		{
			name: "Падение на Землю",
			v:    1,
			res:  "Ракета упадет на Землю",
		},
		{
			name: "Спутник Земли",
			v:    10,
			res:  "Ракета станет спутником Земли",
		},
		{
			name: "Спутник Солнца",
			v:    14,
			res:  "Ракета станет спутником Солнца",
		},
		{
			name: "Выход за пределы Солнечной системы",
			v:    20,
			res:  "Ракета покинет Солнечную систему",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := RocketLunch(c.v)
			if got != c.res {
				t.Fatalf("got %s, but expected %s", got, c.res)
			}
		})
	}
}

func TestPointInsideFigure(t *testing.T) {
	cases := []struct {
		name string
		x, y float64
		res  bool
	}{
		{
			name: "first case",
			x:    -0.4,
			y:    0.95,
			res:  false,
		},
		{
			name: "second case",
			x:    -0.25,
			y:    0.5,
			res:  true,
		},
		{
			name: "third case",
			x:    -0.5,
			y:    0.5,
			res:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := PointInsideFigure(c.x, c.y)
			if got != c.res {
				t.Fatalf("got %v, want %v", got, c.res)
			}
		})
	}
}

func TestMulTable(t *testing.T) {
	got := MulTable()
	want := 64
	if got[8][8] != 64 {
		t.Fatalf("got %d, want %d", got[8][8], want)
	}
}
