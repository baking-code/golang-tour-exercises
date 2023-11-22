package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

// At implements image.Image.
func (Image) At(x int, y int) color.Color {
	return color.RGBA{uint8((x ^ y) / 2), uint8((x ^ y) / 2), 255, 255}
}

// Bounds implements image.Image.
func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 500, 500)
}

// ColorModel implements image.Image.
func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
