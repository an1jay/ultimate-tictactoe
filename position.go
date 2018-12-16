package main

// Position represents a game state
type Position struct {
	WhiteBB              BitBoard
	BlackBB              BitBoard
	sideToMove           Color
	subBoardToPlayOnNext SubBoard // No SubBoard if Subboard is over
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
	if p.subBoardToPlayOnNext != NoSubBoard {
		sbo, _ := p.SubBoardOver(p.subBoardToPlayOnNext)
		if sbo {
			panic("LegalMoves: sub board to play on next is full")
		}
		return p.UnoccupiedSubBoardSquares(p.subBoardToPlayOnNext)
	}

	lm := make([]Square, 50) // 50 is estimate of unoccupied squares, may need to change
	for _, sb := range p.NotOverSubBoards() {
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
	usbs := make([]Square, numSquaresInSubBoard)
	for _, sq := range sb.Squares() {
		if !(p.WhiteBB.Occupied(sq) || p.BlackBB.Occupied(sq)) {
			usbs = append(usbs, sq)
		}
	}
	return usbs
}

// NotOverSubBoards returns a list of sub boards that are not over
func (p *Position) NotOverSubBoards() []SubBoard {
	nosb := make([]SubBoard, numSubBoardsInBoard)
	for _, sb := range AllSubBoards() {
		sbo, _ := p.SubBoardOver(sb)
		if !sbo {
			nosb = append(nosb, sb)
		}
	}
	return nosb
}
