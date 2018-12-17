package player

import (
	"fmt"

	"github.com/an1jay/ultimate-tictactoe/game"
)

// HumanPlayer allows a human to play a game of Ultimate TicTacToe
type HumanPlayer struct{}

// ChooseMove asks a human player to move through stdout/stdin.
func (h *HumanPlayer) ChooseMove(pos game.Position) game.Square {
	var b, s int = -1, -1
	for true {
		fmt.Println("Enter Move: (B0S0 - B8S8)")
		fmt.Scanf("B%dS%d", &b, &s)
		if b >= 0 && b < 9 && s >= 0 && s < 9 {
			break
		}
	}
	fmt.Println("Got Move, giving it to game")
	return game.NewSquare(b, s)
}
