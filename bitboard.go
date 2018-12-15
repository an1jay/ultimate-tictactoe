package main

// BitBoard encodes the positions of a side's moves
// BitBoard.first  = 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0 | 0  0  0
//					 Zeroth sub-board          | First sub-board           | Second sub-board          |unused| sub-b wins
// BitBoard.second = 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0 | 0  0  0
//					 Third sub-board           | Fourth sub-board          | Fifth sub-board           |unused| sub-b wins
// BitBoard.third  = 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0  0  0  0  0  0  0  0 | 0  0 | 0  0  0
//					 Sixth sub-board           | Seventh sub-board         | Eighth sub-board          |unused| sub-b wins

// BitBoard encodes the positions of a side's moves
type BitBoard struct {
	zeroth uint32
	first  uint32
	second uint32
}

// WinConditions are all the win conditions for a sub board in the 9 most significant bits
var WinConditions = [8]uint32{
	3758096384,
	469762048,
	58720256,
	2449473536,
	1224736768,
	612368384,
	2290089984,
	704643072,
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
			var one uint32 = 1
			subboard := sq / numSquaresInSubBoardRow
			switch subboard {
			case 0:
				bb.zeroth ^= one << uint(31-sq)
			case 1:
				bb.first ^= one << uint(31-(sq%27))
			case 2:
				bb.second ^= one << uint(31-(sq%27))
			}
		}
	}
	bb.UpdateSubboardWins()
	return bb
}

// Mapping returns a map from Square to bool, with true for any occupied squares.
func (b *BitBoard) Mapping() map[Square]bool {
	return map[Square]bool{}
}

// UpdateSubboardWins returns a pointer to a BitBoard with the subboard wins updated
func (b *BitBoard) UpdateSubboardWins() *BitBoard {
	subboardWins := [3]bool{}
	for _, subboardRow := range []uint32{b.zeroth, b.first, b.second} {
		for n := 0; n < 3; n++ {
			subboardIsWon := false
			if 
			for _, wc := range WinConditions {
				subboardIsWon = false
				// checks if win condition wc is on the nth subboard in subboardRow
				if (subboardRow^(wc>>uint(n*numSquaresInSubBoard)))&subboardMasks[n] != 0 {
					subboardIsWon = true
				} else {

				}
			}
		}
	}
	return &BitBoard{0, 0, 0}
}
