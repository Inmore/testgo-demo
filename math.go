package math

import "time"

func Add(a, b int) int {
	time.Sleep(time.Second * 10)
	return a + b
}
