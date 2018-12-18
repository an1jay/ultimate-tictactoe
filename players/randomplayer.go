package players

import (
	"math/rand"
	"time"

	"github.com/an1jay/ultimate-tictactoe/game"
)

// RandomPlayer plays UTTT randomly
type RandomPlayer struct{}

// ChooseMove asks a random player to move randomly from choice of legal moves.
func (h *RandomPlayer) ChooseMove(pos game.Position) game.Square {
	rand.Seed(time.Now().UnixNano())
	lm := pos.LegalMoves()
	return lm[rand.Intn(len(lm))]
}
