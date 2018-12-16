package main

import (
	"fmt"
)

// BitBoard encodes the positions of a side's moves
// BitBoard.first  = 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0 | 0  0  0
//					 Zeroth sub-board          | First sub-board           | Second sub-board          |unused| sub-b wins (zeroth, first, second)
// BitBoard.second = 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0 | 0  0  0
//					 Third sub-board           | Fourth sub-board          | Fifth sub-board           |unused| sub-b wins (third, fourth, fifth)
// BitBoard.third  = 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0 | 0  0  0
//					 Sixth sub-board           | Seventh sub-board         | Eighth sub-board          |unused| sub-b wins (sixth, seventh, eighth)

// BitBoard encodes the positions of a side's moves
type BitBoard struct {
	zeroth uint32
	first  uint32
	second uint32
}

// WinConditions are all the win conditions for a sub board in the 9 most significant bits
var WinConditions = [8]uint32{
	3758096384, // Top Row
	469762048,  // Middle Row
	58720256,   // Bottom Row
	2449473536, // Left Column
	1224736768, // Middle Column
	612368384,  // Right Column
	2290089984, // Top L Bot R Diagonal
	704643072,  // Bot L Top R Diagonal
}

var subboardMasks = [3]uint32{
	4286578688,
	8372224,
	16352,
}

// NewBitBoard returns a bitboard given a mapping form Square to Bitboard.
func NewBitBoard(m map[Square]bool) BitBoard {
	bb := BitBoard{0, 0, 0}
	for sq := 0; sq < numSquaresInBoard; sq++ {
		if m[Square(sq)] {
			subboard := sq / numSquaresInSubBoardRow
			switch subboard {
			case 0:
				bb.zeroth ^= 1 << uint(31-sq)
			case 1:
				bb.first ^= 1 << uint(31-(sq%numSquaresInSubBoardRow))
			case 2:
				bb.second ^= 1 << uint(31-(sq%numSquaresInSubBoardRow))
			}
		}
	}
	bb.UpdateSubBoardWins()
	return bb
}

// Mapping returns a map from Square to bool, with true for any occupied squares.
func (b *BitBoard) Mapping() map[Square]bool {
	return map[Square]bool{}
}

// UpdateSubBoardWins returns a pointer to a BitBoard with the subboard wins updated
func (b *BitBoard) UpdateSubBoardWins() {
	subBoardRows := [3]uint32{b.zeroth, b.first, b.second}
	// looping through subboardrows
	for rowNo, subBoardRow := range []uint32{b.zeroth, b.first, b.second} {
		// looping through subboard
		for n := 0; n < numSubBoardsInSubBoardRow; n++ {
			// check if subboard already won -> no need to check win again
			if (subBoardRow>>uint(numSubBoardsInSubBoardRow-1-n))&1 == 1 {
			} else {
				// subboard may be won - need to check every win condition on that
				for _, wc := range WinConditions {
					// checks if win condition wc is on the nth subboard in subboardRow
					if (subBoardRow & subboardMasks[n] & (wc >> uint(n*numSquaresInSubBoard))) != 0 {
						// then updates the corresponding sub-b win bit and breaks
						subBoardRows[rowNo] = subBoardRow + (1 << uint(numSubBoardsInSubBoardRow-1-n))
						break
					}
				}
			}
		}
	}
	// update the subboardrows on each row
	b.zeroth = subBoardRows[0]
	b.first = subBoardRows[1]
	b.second = subBoardRows[2]
}

// Win checks if the bitboard is won
func (b *BitBoard) Win() bool {
	// update subboard wins
	b.UpdateSubBoardWins()

	// generate bigBoard - as if each subboard was a square
	var bigBoard uint32
	bigBoard += (7 & b.zeroth) << 6
	bigBoard += (7 & b.first) << 3
	bigBoard += 7 & b.second
	bigBoard <<= 23

	// check win on bigBoard
	for _, wc := range WinConditions {
		if bigBoard&wc != 0 {
			return true
		}
	}
	return false
}

// SubBoardWin returns bool for whether the SubBoard is won or not.
// NoSubBoard always returns false
func (b *BitBoard) SubBoardWin(sb SubBoard) bool {
	b.UpdateSubBoardWins()
	switch sb {
	case SubBoard0:
		if b.zeroth&(1<<2) != 0 {
			return true
		}
	case SubBoard1:
		if b.zeroth&(1<<1) != 0 {
			return true
		}
	case SubBoard2:
		if b.zeroth&(1<<0) != 0 {
			return true
		}
	case SubBoard3:
		if b.first&(1<<2) != 0 {
			return true
		}
	case SubBoard4:
		if b.first&(1<<1) != 0 {
			return true
		}
	case SubBoard5:
		if b.first&(1<<0) != 0 {
			return true
		}
	case SubBoard6:
		if b.second&(1<<2) != 0 {
			return true
		}
	case SubBoard7:
		if b.second&(1<<1) != 0 {
			return true
		}
	case SubBoard8:
		if b.second&(1<<0) != 0 {
			return true
		}
	}
	return false
}

// PrintRowsAsBits prints a bit representation of the various rows to stdout
func (b *BitBoard) PrintRowsAsBits() {
	fmt.Printf("%032b\n", b.zeroth)
	fmt.Printf("%032b\n", b.first)
	fmt.Printf("%032b\n", b.second)
}

// Move makes a move on the board, i.e. sets the bit corresponding to sq to 1;
// returns true if valid move, returns false if invalid (does not make invalid move).
func (b *BitBoard) Move(sq Square) bool {
	sqint := int(sq)
	sbrow := sqint / numSquaresInSubBoardRow
	if b.Occupied(sq) {
		return false
	}
	switch sbrow {
	case 0:
		b.zeroth ^= 1 << uint(31-sqint)
		return true
	case 1:
		b.first ^= 1 << uint(31-(sqint%numSquaresInSubBoardRow))
		return true
	case 2:
		b.second ^= 1 << uint(31-(sqint%numSquaresInSubBoardRow))
		return true
	}
	return false // unreachable code xD
}

// Occupied returns true if square is occupied, otherwise false.
func (b *BitBoard) Occupied(sq Square) bool {
	sqint := int(sq)
	sbrow := sqint / numSquaresInSubBoardRow
	switch sbrow {
	case 0:
		if b.zeroth&(1<<uint(31-sqint)) != 0 {
			return true
		}
	case 1:
		if b.first&(1<<uint(31-(sqint%numSquaresInSubBoardRow))) != 0 {
			return true
		}
	case 2:
		if b.second&(1<<uint(31-(sqint%numSquaresInSubBoardRow))) != 0 {
			return true
		}
	}
	return false
}

// MoveCount returns number of occupied squares of a SubBoard
func (b *BitBoard) MoveCount(sb SubBoard) uint8 {
	var counter uint8
	for _, sq := range sb.Squares() {
		if b.Occupied(sq) {
			counter++
		}
	}
	return counter
}

// OccupiedChar return "1" if sq is occupied, otherwise "0".
func (b *BitBoard) OccupiedChar(sq Square) string {
	if b.Occupied(sq) {
		return "1"
	}
	return "0"
}

// Display prints a board representation to stdout.
func (b *BitBoard) Display() {
	b.funcDisplay(b.OccupiedChar)
}

func (b *BitBoard) funcDisplay(f func(Square) string) {

	fmt.Printf(
		"%s %s %s | %s %s %s | %s %s %s \n"+
			"%s %s %s | %s %s %s | %s %s %s \n"+
			"%s %s %s | %s %s %s | %s %s %s \n"+
			"---------------------\n"+
			"%s %s %s | %s %s %s | %s %s %s \n"+
			"%s %s %s | %s %s %s | %s %s %s \n"+
			"%s %s %s | %s %s %s | %s %s %s \n"+
			"---------------------\n"+
			"%s %s %s | %s %s %s | %s %s %s \n"+
			"%s %s %s | %s %s %s | %s %s %s \n"+
			"%s %s %s | %s %s %s | %s %s %s \n",
		f(B0S0), f(B0S1), f(B0S2), f(B1S0), f(B1S1), f(B1S2), f(B2S0), f(B2S1), f(B2S2),
		f(B0S3), f(B0S4), f(B0S5), f(B1S3), f(B1S4), f(B1S5), f(B2S3), f(B2S4), f(B2S5),
		f(B0S6), f(B0S7), f(B0S8), f(B1S6), f(B1S7), f(B1S8), f(B2S6), f(B2S7), f(B2S8),
		// ----------------------------------------------------------------------------
		f(B3S0), f(B3S1), f(B3S2), f(B4S0), f(B4S1), f(B4S2), f(B5S0), f(B5S1), f(B5S2),
		f(B3S3), f(B3S4), f(B3S5), f(B4S3), f(B4S4), f(B4S5), f(B5S3), f(B5S4), f(B5S5),
		f(B3S6), f(B3S7), f(B3S8), f(B4S6), f(B4S7), f(B4S8), f(B5S6), f(B5S7), f(B5S8),
		// ----------------------------------------------------------------------------
		f(B6S0), f(B6S1), f(B6S2), f(B7S0), f(B7S1), f(B7S2), f(B8S0), f(B8S1), f(B8S2),
		f(B6S3), f(B6S4), f(B6S5), f(B7S3), f(B7S4), f(B7S5), f(B8S3), f(B8S4), f(B8S5),
		f(B6S6), f(B6S7), f(B6S8), f(B7S6), f(B7S7), f(B7S8), f(B8S6), f(B8S7), f(B8S8),
	)
}
