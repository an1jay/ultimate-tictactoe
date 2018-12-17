package game

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

// String returns a string reprsentation of c.
func (c Color) String() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	}
	return "NoColor"
}

// EvaluationCoefficient return 1 for White, -1 for Black, 0 for NoColor
func (c Color) EvaluationCoefficient() float32 {
	switch c {
	case White:
		return 1
	case Black:
		return -1
	}
	return 0
}
