package players

import (
	"github.com/an1jay/ultimate-tictactoe/game"
)

// Evaluator evaluates a position and returns a score depending on how favourable the position is
// -1 is best for black, 1 is best for white
type Evaluator interface {
	Evaluate(p *game.Position) float32
}

// FirstEvaluator is basic
type FirstEvaluator struct{}

// Evaluate evaluates a position and returns a score depending on how favourable the position is
func (b FirstEvaluator) Evaluate(p *game.Position) float32 {
	igo, winner := p.GameOver()
	if igo {
		return winner.EvaluationCoefficient()
	}
	value := float32(p.CountLegalMoves()) / 500
	return p.SideToMove.EvaluationCoefficient() * value
}
