package game

import "fmt"

// Position represents a game state
type Position struct {
	WhiteBB              BitBoard
	BlackBB              BitBoard
	SideToMove           Color
	SubBoardToPlayOnNext SubBoard // No SubBoard if Subboard is over
}

// GameOver returns true and Color of victor if game is over, else returns false
func (p *Position) GameOver() (bool, Color) {
	if p.WhiteBB.Win() {
		return true, White
	}
	if p.BlackBB.Win() {
		return true, Black
	}
	if len(p.LegalMoves()) == 0 {
		return true, NoColor
	}
	return false, NoColor
}

// LegalMoves returns a slice of all possible moves for the side to move
func (p *Position) LegalMoves() []Square {
	if p.SubBoardToPlayOnNext != NoSubBoard {
		return p.UnoccupiedSubBoardSquares(p.SubBoardToPlayOnNext)
	}

	lm := []Square{}
	for _, sb := range p.NotOverSubBoards() {
		// fmt.Println("Subboard: ", sb.String())
		// fmt.Println("Unoccupied sq: ", p.UnoccupiedSubBoardSquares(sb))
		lm = append(lm, p.UnoccupiedSubBoardSquares(sb)...)
	}

	return lm
}

// CountLegalMoves returns number of legal moves
func (p *Position) CountLegalMoves() uint8 {
	if p.SubBoardToPlayOnNext != NoSubBoard {
		return p.CountUnoccupiedSubBoardSquares(p.SubBoardToPlayOnNext)
	}

	var clm uint8
	for _, sb := range p.NotOverSubBoards() {
		clm += p.CountUnoccupiedSubBoardSquares(sb)
	}

	return clm
}

// IsMoveLegal returns a bool of whether a square is legal or not
func (p *Position) IsMoveLegal(sq Square) bool {
	if p.SubBoardToPlayOnNext != NoSubBoard {
		if !(p.WhiteBB.Occupied(sq) || p.BlackBB.Occupied(sq)) && (sq.SubBoard() == p.SubBoardToPlayOnNext) {
			return true
		}
		return false
	}
	subbo, _ := p.SubBoardOver(sq.SubBoard())
	if !(p.WhiteBB.Occupied(sq) || p.BlackBB.Occupied(sq)) && !subbo {
		return true
	}
	return false

}

// SubBoardOver returns true and Color of victor if SubBoard is over, else returns false
func (p *Position) SubBoardOver(sb SubBoard) (bool, Color) {
	if p.WhiteBB.SubBoardWin(sb) {
		return true, White
	}
	if p.BlackBB.SubBoardWin(sb) {
		return true, Black
	}
	if p.WhiteBB.MoveCount(sb)+p.BlackBB.MoveCount(sb) == uint8(numSquaresInSubBoard) {
		return true, NoColor
	}
	return false, NoColor
}

// UnoccupiedSubBoardSquares returns a slice of squares that are unoccupied in sb
func (p *Position) UnoccupiedSubBoardSquares(sb SubBoard) []Square {
	usbs := []Square{}
	for _, sq := range sb.Squares() {
		if !(p.WhiteBB.Occupied(sq) || p.BlackBB.Occupied(sq)) {
			usbs = append(usbs, sq)
		}
	}
	return usbs
}

// CountUnoccupiedSubBoardSquares returns number of unoccupied squares in a subboard
func (p *Position) CountUnoccupiedSubBoardSquares(sb SubBoard) uint8 {
	var cusbs uint8
	for _, sq := range sb.Squares() {
		if !(p.WhiteBB.Occupied(sq) || p.BlackBB.Occupied(sq)) {
			cusbs++
		}
	}
	return cusbs
}

// NotOverSubBoards returns a list of subboards that are not over
func (p *Position) NotOverSubBoards() []SubBoard {
	nosb := []SubBoard{}
	for _, sb := range AllSubBoards() {
		sbo, _ := p.SubBoardOver(sb)
		if !sbo {
			nosb = append(nosb, sb)
		}
	}
	return nosb
}

// SqInSlice checks whether Square is in a Square slice
func SqInSlice(sq Square, sqs []Square) bool {
	for _, square := range sqs {
		if square == sq {
			return true
		}
	}
	return false
}

// Move makes a move on the board, i.e. makes move on c's bitboard
// returns true if valid move, returns false if invalid (does not make invalid move).
func (p *Position) Move(sq Square, c Color) bool {
	rv := false
	// if wrong side tries to move - error
	if c != p.SideToMove || c == NoColor {
		return false
	}

	// // make slice of legal moves
	// lm := p.LegalMoves()

	// if move is legal
	if p.IsMoveLegal(sq) {
		switch c {
		case White:
			rv = p.WhiteBB.Move(sq)
		case Black:
			rv = p.BlackBB.Move(sq)
		}
		// update subboards
		sssb := sq.SubSquareSubBoard()
		so, _ := p.SubBoardOver(sssb)
		if so {
			p.SubBoardToPlayOnNext = NoSubBoard
		} else {
			p.SubBoardToPlayOnNext = sssb
		}
		p.SideToMove = p.SideToMove.Other()
	}
	return rv
}

// BigBoard returns a 3x3 board of the subboards for the specific Color
func (p *Position) BigBoard(c Color) uint32 {
	switch c {
	case White:
		return p.WhiteBB.BigBoard()
	case Black:
		return p.BlackBB.BigBoard()
	}
	return 0
}

// Copy returns a pointer to a copy of the position
func (p *Position) Copy() *Position {
	return &Position{
		WhiteBB:              p.WhiteBB,
		BlackBB:              p.BlackBB,
		SideToMove:           p.SideToMove,
		SubBoardToPlayOnNext: p.SubBoardToPlayOnNext,
	}
}

// String returns a string representation of the board, suitable for printing.
func (p *Position) String() string {
	return p.funcString(getCharForSquare)
}

func getCharForSquare(p *Position, sq Square) string {
	if p.WhiteBB.Occupied(sq) {
		return "W"
	} else if p.BlackBB.Occupied(sq) {
		return "B"
	}
	return " "
}

func (p *Position) funcString(f func(*Position, Square) string) string {

	return fmt.Sprintf(
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
			"%s %s %s | %s %s %s | %s %s %s \n"+

			"%s to move on %s. \n",

		f(p, B0S0), f(p, B0S1), f(p, B0S2), f(p, B1S0), f(p, B1S1), f(p, B1S2), f(p, B2S0), f(p, B2S1), f(p, B2S2),
		f(p, B0S3), f(p, B0S4), f(p, B0S5), f(p, B1S3), f(p, B1S4), f(p, B1S5), f(p, B2S3), f(p, B2S4), f(p, B2S5),
		f(p, B0S6), f(p, B0S7), f(p, B0S8), f(p, B1S6), f(p, B1S7), f(p, B1S8), f(p, B2S6), f(p, B2S7), f(p, B2S8),
		// ----------------------------------------------------------------------------
		f(p, B3S0), f(p, B3S1), f(p, B3S2), f(p, B4S0), f(p, B4S1), f(p, B4S2), f(p, B5S0), f(p, B5S1), f(p, B5S2),
		f(p, B3S3), f(p, B3S4), f(p, B3S5), f(p, B4S3), f(p, B4S4), f(p, B4S5), f(p, B5S3), f(p, B5S4), f(p, B5S5),
		f(p, B3S6), f(p, B3S7), f(p, B3S8), f(p, B4S6), f(p, B4S7), f(p, B4S8), f(p, B5S6), f(p, B5S7), f(p, B5S8),
		// ----------------------------------------------------------------------------
		f(p, B6S0), f(p, B6S1), f(p, B6S2), f(p, B7S0), f(p, B7S1), f(p, B7S2), f(p, B8S0), f(p, B8S1), f(p, B8S2),
		f(p, B6S3), f(p, B6S4), f(p, B6S5), f(p, B7S3), f(p, B7S4), f(p, B7S5), f(p, B8S3), f(p, B8S4), f(p, B8S5),
		f(p, B6S6), f(p, B6S7), f(p, B6S8), f(p, B7S6), f(p, B7S7), f(p, B7S8), f(p, B8S6), f(p, B8S7), f(p, B8S8),

		p.SideToMove.String(), p.SubBoardToPlayOnNext.String(),
	)
}
