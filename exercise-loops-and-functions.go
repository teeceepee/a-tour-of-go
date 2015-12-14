package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; ; i++ {
		z = z - ((z * z - x) / (2.0 * z))

		delta := math.Abs(z * z - x)
		if delta < 0.000001 {
			fmt.Printf("Try count: %v.\n", i)
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
