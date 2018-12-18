package game

// Result represents a result for a side.
type Result uint8

// Enumerating the types of results (for a side)
const (
	Tie Result = iota
	Win
	Loss
)

func (r Result) String() string {
	switch r {
	case Win:
		return "Win"
	case Loss:
		return "Loss"
	}
	return "Tie"
}

// Other returns opposite result
func (r Result) Other() Result {
	switch r {
	case Win:
		return Loss
	case Loss:
		return Win
	}
	return Tie
}
