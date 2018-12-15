package main

// Color encodes which player moves first and second.
type Color bool

// Enumerating the sides.
const (
	White Color = true  // side which moves first
	Black Color = false // side which moves second
)

// Other returns the other color to the given.
func (c Color) Other() Color {
	return Color(!bool(White))
}
