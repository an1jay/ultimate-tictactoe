package players

import (
	"math/bits"

	"github.com/an1jay/ultimate-tictactoe/game"
)

// Evaluator evaluates a position and returns a score depending on how favourable the position is
// -1 is best for black, 1 is best for white
type Evaluator interface {
	Evaluate(p *game.Position) float32
}

// MaxLegalMoveEvaluator is not very useful...
type MaxLegalMoveEvaluator struct{}

// Evaluate evaluates a position and returns a score depending on how favourable the position is
func (m MaxLegalMoveEvaluator) Evaluate(pos *game.Position) float32 {
	igo, winner := pos.GameOver()
	if igo {
		return winner.EvaluationWin()
	}
	value := float32(pos.CountLegalMoves()) / 500
	return pos.SideToMove.EvaluationCoefficient() * value
}

// FirstEvaluator evaluates a position
type FirstEvaluator struct{}

// Evaluate evaluates a position and returns a score depending on how favourable the position is
func (b FirstEvaluator) Evaluate(pos *game.Position) float32 {
	igo, winner := pos.GameOver()
	if igo {
		return winner.EvaluationWin()
	}
	// Theory:
	//  * Winning subboards is good
	//  * Number of legal moves is good(ish)
	//  * Winning subboards relative to win condition - e.g. xor bigboard with wc and check

	bbSM := pos.BigBoard(pos.SideToMove)

	// calculate lowest mse with a win condition on a board
	bestMSESM := 1000
	mseSM := 0
	for _, wc := range game.WinConditions {
		mseSM = bits.OnesCount32(wc ^ bbSM)
		if mseSM < bestMSESM {
			bestMSESM = mseSM
		}
	}
	fMseSM := float32(mseSM) / 9

	// calculate number of legal moves
	nlm := float32(pos.CountLegalMoves())

	// number of won subboards
	nwbbSM := float32(bits.OnesCount32(bbSM)) / 9

	// ---------------------------------------------

	bbSMO := pos.BigBoard(pos.SideToMove.Other())

	// calculate lowest mse with a win condition on a board
	bestMSESMO := 1000
	mseSMO := 0
	for _, wc := range game.WinConditions {
		mseSMO = bits.OnesCount32(wc ^ bbSMO)
		if mseSMO < bestMSESMO {
			bestMSESMO = mseSM
		}
	}
	fMseSMO := float32(mseSMO) / 9

	// number of won subboards
	nwbbSMO := float32(bits.OnesCount32(bbSMO)) / 9

	value := .0002*nlm + (-0.08)*(fMseSM-fMseSMO) + 0.04*(nwbbSM-nwbbSMO)
	return pos.SideToMove.EvaluationCoefficient() * value
}
