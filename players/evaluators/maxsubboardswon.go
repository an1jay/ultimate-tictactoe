package evaluators

import (
	"math/bits"

	"github.com/an1jay/ultimate-tictactoe/game"
)

// MaxSubBoardsWon evaluates a position
type MaxSubBoardsWon struct{}

// Evaluate evaluates a position and returns a score depending on how favourable the position is
func (b MaxSubBoardsWon) Evaluate(pos *game.Position) float32 {
	igo, winner := pos.GameOver()
	if igo {
		return winner.EvaluationWin()
	}
	// Theory:
	//  * Winning subboards is good
	//  * Number of legal moves is good(ish)
	//  * Winning subboards relative to win condition - e.g. xor bigboard with wc and check

	bbSM := pos.BigBoard(pos.SideToMove)

	// number of won subboards
	nwbbSM := float32(bits.OnesCount32(bbSM)) / 9

	// ---------------------------------------------

	bbSMO := pos.BigBoard(pos.SideToMove.Other())

	// number of won subboards
	nwbbSMO := float32(bits.OnesCount32(bbSMO)) / 9

	value := nwbbSM - nwbbSMO
	return pos.SideToMove.EvaluationCoefficient() * value
}
