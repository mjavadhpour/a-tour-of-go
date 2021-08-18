package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	result := float64(0)
	for {
		z -= (z*z - x) / (2*z)
		if math.Abs(result - z) < 1e-15 {
			break
		}
		result = z
	}
	return result
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}