package player

import (
	"fmt"

	"github.com/an1jay/ultimate-tictactoe/game"
)

// HumanPlayer allows a human to play a game of Ultimate TicTacToe
type HumanPlayer struct{}

// Move asks a human player to move through stdout/stdin.
func (h *HumanPlayer) Move(pos game.Position) game.Square {
	var b, s uint
	for true {
		fmt.Println("Enter Move: (B0S0 - B8S8)")
		fmt.Scanf("B%dS%d", &b, &s)
		if b >= 0 && b < 9 && s >= 0 && s < 9 {
			break
		}
	}
	return game.NewSquare(b, s)
}
