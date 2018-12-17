package game

import "fmt"

// Square represents one of the 81 squares on the board.
type Square uint8

const (
	numSquaresInSubBoard    int = 9
	numSquaresInBoard       int = 81
	numSquaresInSubBoardRow int = 27

	numSubBoardsInBoard       int = 9
	numSubBoardsInSubBoardRow int = 3
	numSubBoardRowsInBoard    int = 3
)

// String returns a string representation of a square.
// DOES NOT (investigate?) Implements the fmt.Stringer() interface.
func (sq Square) String() string {
	//fmt.Printf("%d", int(sq)/numSubBoardsInBoard)
	//fmt.Printf("%d", int(sq)/numSquaresInSubBoard)
	return fmt.Sprintf("B%dS%d", int(sq)/numSubBoardsInBoard, int(sq)%numSquaresInSubBoard)
}

// SubBoard returns the subboard that a square is on
func (sq Square) SubBoard() SubBoard {
	sb := int(sq) / numSquaresInSubBoard
	switch sb {
	case 0:
		return SubBoard0
	case 1:
		return SubBoard1
	case 2:
		return SubBoard2
	case 3:
		return SubBoard3
	case 4:
		return SubBoard4
	case 5:
		return SubBoard5
	case 6:
		return SubBoard6
	case 7:
		return SubBoard7
	case 8:
		return SubBoard8
	}
	return NoSubBoard
}

// SubSquareSubBoard returns the corresponding SubBoard to SubSquare that the square is on
func (sq Square) SubSquareSubBoard() SubBoard {
	sqint := int(sq)
	return SubBoard(sqint%numSquaresInSubBoard + 1)
}

// NewSquare constructs a square from subboard (b) and subsquare (s)
func NewSquare(b, s int) Square {
	if b >= 0 && b < 9 && s >= 0 && s < 9 {
		return Square(b*numSquaresInSubBoard + s)
	}
	panic(fmt.Sprintf("Invalid values for subboard (%d) and square (%d", b, s))
}

// Enumerating through the Squares on the board.
// e.g. B0S0 is the zeroth square of the zeroth subboard.
const (
	B0S0 Square = iota
	B0S1
	B0S2
	B0S3
	B0S4
	B0S5
	B0S6
	B0S7
	B0S8

	B1S0
	B1S1
	B1S2
	B1S3
	B1S4
	B1S5
	B1S6
	B1S7
	B1S8

	B2S0
	B2S1
	B2S2
	B2S3
	B2S4
	B2S5
	B2S6
	B2S7
	B2S8

	B3S0
	B3S1
	B3S2
	B3S3
	B3S4
	B3S5
	B3S6
	B3S7
	B3S8

	B4S0
	B4S1
	B4S2
	B4S3
	B4S4
	B4S5
	B4S6
	B4S7
	B4S8

	B5S0
	B5S1
	B5S2
	B5S3
	B5S4
	B5S5
	B5S6
	B5S7
	B5S8

	B6S0
	B6S1
	B6S2
	B6S3
	B6S4
	B6S5
	B6S6
	B6S7
	B6S8

	B7S0
	B7S1
	B7S2
	B7S3
	B7S4
	B7S5
	B7S6
	B7S7
	B7S8

	B8S0
	B8S1
	B8S2
	B8S3
	B8S4
	B8S5
	B8S6
	B8S7
	B8S8
)
