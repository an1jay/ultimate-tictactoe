package evaluators

import "github.com/an1jay/ultimate-tictactoe/game"

// MaxLegalMove is not very useful...
type MaxLegalMove struct{}

// Evaluate evaluates a position and returns a score depending on how favourable the position is
func (m MaxLegalMove) Evaluate(pos *game.Position) float32 {
	igo, winner := pos.GameOver()
	if igo {
		return winner.EvaluationWin()
	}
	value := float32(pos.CountLegalMoves()) / 150
	return pos.SideToMove.EvaluationCoefficient() * value
}
