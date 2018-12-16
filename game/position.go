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
		sbo, _ := p.SubBoardOver(p.SubBoardToPlayOnNext)
		if sbo {
			panic("LegalMoves: sub board to play on next is full")
		}
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

// Move makes a move on the board, i.e. makes move on c's bitboard
// returns true if valid move, returns false if invalid (does not make invalid move).
func (p *Position) Move(sq Square, c Color) bool {
	if c != p.SideToMove || c == NoColor {
		return false
	}
	switch c {
	case White:
		return p.WhiteBB.Move(sq)
	case Black:
		return p.BlackBB.Move(sq)
	}
	return false // unreachable code xD
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
