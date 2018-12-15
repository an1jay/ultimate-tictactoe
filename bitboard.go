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
	first  uint32
	second uint32
	third  uint32
}

// NewBitBoard returns a bitboard given a mapping form Square to Bitboard.
func NewBitBoard(map[Square]bool) BitBoard {
	for sq := range numSquaresInBoard {

	}
	return BitBoard{0, 0, 0}
}

// Mapping returns a map from Square to bool, with true for any occupied squares.
func (b *BitBoard) Mapping() map[Square]bool {
	return map[Square]bool{}
}

func (b *BitBoard) UpdateSubboardWins() *BitBoard {

}
