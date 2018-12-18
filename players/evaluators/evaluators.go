package evaluators

import (
	"github.com/an1jay/ultimate-tictactoe/game"
)

// Evaluator evaluates a position and returns a score depending on how favourable the position is
// -1 is best for black, 1 is best for white
type Evaluator interface {
	Evaluate(p *game.Position) float32
}
