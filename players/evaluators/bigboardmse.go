package evaluators

import (
	"math/bits"

	"github.com/an1jay/ultimate-tictactoe/game"
)

// BigBoardMSE evaluates a position
type BigBoardMSE struct{}

// Evaluate evaluates a position and returns a score depending on how favourable the position is
func (b BigBoardMSE) Evaluate(pos *game.Position) float32 {
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

	value := 0.8 * (fMseSMO - fMseSM)
	return pos.SideToMove.EvaluationCoefficient() * value
}
