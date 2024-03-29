package main

import (
	"fmt"
)

// Color represents an RGB color.
type Color struct {
	R float64
	G float64
	B float64
}

// NewColor returns a new color instance.
func NewColor(r, g, b float64) Color {
	return Color{R: r, G: g, B: b}
}
func (c Color) RGB() string {
	return fmt.Sprintf("%d %d %d", int(255.999*c.R), int(255.999*c.G), int(255.999*c.B))
}

// Lerp stand for linear interpolation.
func (c Color) Lerp(end Color, factor float64) Color {
	oneMinusFactor := 1 - factor

	return Color{
		oneMinusFactor*c.R + factor*end.R,
		oneMinusFactor*c.G + factor*end.G,
		oneMinusFactor*c.B + factor*end.B,
	}
}
