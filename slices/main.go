package main

import "golang.org/x/tour/pic"

//import "math"

func Pic(dx, dy int) [][]uint8 {
	// create 2d slice, of length dy
	ss := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		// for each row, create slice for column length dx
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// for each cell, use coordinates to calculate some funky value - this will render differently depending on the equation used
			s[x] = uint8((x ^ 2 - y ^ 2) / 2)
			// s[x] = uint8((x + y) / 2)
			// s[x] = uint8((x ^ y) / 2)
		}
		ss[y] = s
	}
	return ss
}

func main() {
	pic.Show(Pic)
}
