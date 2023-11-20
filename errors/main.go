package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

/*
calculate an approximation for sqrt
*/
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	// initial guess - I don't think it really matters what this is
	z := x / 2
	for n := 0; n < 10; n += 1 {
		prev := z
		// our approximation - take difference between previous attempt and the derivative
		z -= (z*z - x) / (2 * z)
		fmt.Printf("iteration %d -> result %g\n", n, z)
		// terminate if we're close enough
		if math.Abs(z-prev) < 1e-9 {
			return z, nil
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
