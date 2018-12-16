package main

import (
	"fmt"
	"time"

	"github.com/an1jay/ultimate-tictactoe/game"
)

// A Player play the game.
// If an object has a Move method - it is a player
type Player interface {
	Move(game.Position) game.Square
}

// PrintIf prints str using fmt.Println() if b is true.
func PrintIf(b bool, str string) {
	if b {
		fmt.Println(str)
	}
}

// PlayGame plays a game of Ultimate TicTacToe, returning a map[Color]int
func PlayGame(white, black Player, verbose bool) map[game.Color]game.Result {
	// make result map
	res := map[game.Color]game.Result{}
	// make newgame position
	pos := game.Position{
		WhiteBB:              game.BitBoard{},
		BlackBB:              game.BitBoard{},
		SideToMove:           game.White,
		SubBoardToPlayOnNext: game.NoSubBoard,
	}

	// main game loop
	for true {
		time.Sleep(1 * time.Second) // Leave in so doesn't crash

		// depending on whose move, get move
		switch pos.SideToMove {
		case game.White:
			PrintIf(verbose, pos.String())
			m := white.Move(pos)
			PrintIf(verbose, fmt.Sprintf("%s plays %s", pos.SideToMove.String(), m.String()))
		case game.Black:
			PrintIf(verbose, pos.String())
			m := black.Move(pos)
			PrintIf(verbose, fmt.Sprintf("%s plays %s", pos.SideToMove.String(), m.String()))
		}

		// set move to other side
		pos.SideToMove = pos.SideToMove.Other()

		//check gameover
		igo, winner := pos.GameOver()
		if igo {
			switch winner {
			case game.White:
				res[game.White] = game.Win
				res[game.Black] = game.Loss
			case game.Black:
				res[game.White] = game.Loss
				res[game.Black] = game.Win
			case game.NoColor:
				res[game.White] = game.Tie
				res[game.Black] = game.Tie
			}
			break
		}

	}
	return res
}
