package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	switch {
	case x == 0:
		return 0
	case x < 0:
		return math.NaN()
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
}
