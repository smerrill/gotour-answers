package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	y, z, count := 1.0, x, 1
	for y >= 0.000001 {
		y = ((math.Pow(z, 2) - x) / (2 * z))
		z = z - y
		count++
	}
	return z, count
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
