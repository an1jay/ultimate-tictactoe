package game

// SubBoard represents one of the 9 3*3 boards which comprise an Ultimate Tic Tac Toe board.
type SubBoard uint8

// Enumerating the possible subboards
const (
	NoSubBoard SubBoard = iota
	SubBoard0
	SubBoard1
	SubBoard2
	SubBoard3
	SubBoard4
	SubBoard5
	SubBoard6
	SubBoard7
	SubBoard8
)

func (sb SubBoard) String() string {
	switch sb {
	case SubBoard0:
		return "SubBoard 0"
	case SubBoard1:
		return "SubBoard 1"
	case SubBoard2:
		return "SubBoard 2"
	case SubBoard3:
		return "SubBoard 3"
	case SubBoard4:
		return "SubBoard 4"
	case SubBoard5:
		return "SubBoard 5"
	case SubBoard6:
		return "SubBoard 6"
	case SubBoard7:
		return "SubBoard 7"
	case SubBoard8:
		return "SubBoard 8"
	}
	return "No SubBoard"
}

func (sb SubBoard) SubSquareSubBoard() SubBoard {
	sbint := int(sb)
	return SubBoard()
}

// Squares returns a slice of all squares in SubBoard
func (sb SubBoard) Squares() []Square {

	switch sb {
	case SubBoard0:
		return []Square{B0S0, B0S1, B0S2, B0S3, B0S4, B0S5, B0S6, B0S7, B0S8}
	case SubBoard1:
		return []Square{B1S0, B1S1, B1S2, B1S3, B1S4, B1S5, B1S6, B1S7, B1S8}
	case SubBoard2:
		return []Square{B2S0, B2S1, B2S2, B2S3, B2S4, B2S5, B2S6, B2S7, B2S8}
	case SubBoard3:
		return []Square{B3S0, B3S1, B3S2, B3S3, B3S4, B3S5, B3S6, B3S7, B3S8}
	case SubBoard4:
		return []Square{B4S0, B4S1, B4S2, B4S3, B4S4, B4S5, B4S6, B4S7, B4S8}
	case SubBoard5:
		return []Square{B5S0, B5S1, B5S2, B5S3, B5S4, B5S5, B5S6, B5S7, B5S8}
	case SubBoard6:
		return []Square{B6S0, B6S1, B6S2, B6S3, B6S4, B6S5, B6S6, B6S7, B6S8}
	case SubBoard7:
		return []Square{B7S0, B7S1, B7S2, B7S3, B7S4, B7S5, B7S6, B7S7, B7S8}
	case SubBoard8:
		return []Square{B8S0, B8S1, B8S2, B8S3, B8S4, B8S5, B8S6, B8S7, B8S8}
	}
	return []Square{}
}

// AllSubBoards returns a list SubBoards 0 - 8
func AllSubBoards() [9]SubBoard {
	return [9]SubBoard{SubBoard0, SubBoard1, SubBoard2, SubBoard3, SubBoard4, SubBoard5, SubBoard6, SubBoard7, SubBoard8}
}
