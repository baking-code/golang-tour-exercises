package main

import (
	"fmt"
	"math"
)

/*
calculate an approximation for sqrt
*/
func Sqrt(x float64) float64 {
	// initial guess - I don't think it really matters what this is
	z := x / 2
	for n := 0; n < 10; n += 1 {
		prev := z
		// our approximation - take difference between previous attempt and the derivative
		z -= (z*z - x) / (2 * z)
		fmt.Printf("iteration %d -> result %g\n", n, z)
		// terminate if we're close enough
		if math.Abs(z-prev) < 1e-9 {
			return z
		}
	}
	return z
}

func main() {
	var x = 2.0
	fmt.Printf("Real %g\n", math.Sqrt(x))
	fmt.Println(Sqrt(x))
}
