package main

// Color encodes which player moves first and second.
type Color uint8

// Enumerating the sides.
const (
	NoColor Color = iota // neither side
	White                // side which moves first
	Black                // side which moves second
)

// Other returns the other color to the given.
func (c Color) Other() Color {
	switch c {
	case White:
		return Black
	case Black:
		return White
	}
	return NoColor
}
