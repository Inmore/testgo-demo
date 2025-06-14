package math

import (
	"fmt"
	"math"
	"time"
)

func Add(a, b int) int {
	time.Sleep(time.Second * 1)
	return a + b
}

func SolveQuadraticEquation(a, b, c float64) (x1, x2 float64, m string) {
	d := b*b - 4*a*c
	if d < 0 {
		return 0, 0, "решений нет"
	}
	x1 = (-b + math.Sqrt(d)) / (2 * a)
	x2 = (-b - math.Sqrt(d)) / (2 * a)
	return x1, x2, ""
}

// Алгоритм Евклида для нахождения наибольшего общего делителя
func GreatestCommonDivisor(x, y uint) uint {
	if x == 0 || y == 0 {
		return 0
	}
	for x != y {
		if x > y {
			x = x - y
		} else {
			y = y - x
		}
	}
	return x
}

func RocketLunch(v float64) string {
	res := ""
	if v < 0 {
		res = "Скорость не может быть меньше 0"
	} else if v < 7.8 {
		res = "Ракета упадет на Землю"
	} else if v < 11.2 {
		res = "Ракета станет спутником Земли"
	} else if v < 16.4 {
		res = "Ракета станет спутником Солнца"
	} else {
		res = "Ракета покинет Солнечную систему"
	}
	return res
}

func PointInsideFigure(x, y float64) bool {
	if y >= 0 && (math.Abs(x)+math.Abs(y)) <= 1 {
		return true
	} else {
		return false
	}
}

func MulTable() [10][10]int {
	var table [10][10]int
	for i := 1; i < 10; i++ {
		table[i][0] = i
		table[0][i] = i
	}
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			table[i][j] = i * j
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 0 && j == 0 {
				fmt.Printf("    ")
			} else {
				fmt.Printf("%4d", table[i][j])
			}
		}
		fmt.Println()
	}

	return table
}
