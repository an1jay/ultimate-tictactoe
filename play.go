package main

import (
	"fmt"
	"time"

	"github.com/an1jay/ultimate-tictactoe/game"
)

// A Player plays the game.
// If an object has a Move method - it is a player
type Player interface {
	ChooseMove(game.Position) game.Square
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
		time.Sleep(200 * time.Millisecond) // Leave in so doesn't crash

		toMove := pos.SideToMove

		// depending on whose move, get move
		if toMove == game.White {
			PrintIf(verbose, pos.String())
			mW := white.ChooseMove(pos)
			// illegal move loses
			if !pos.Move(mW, game.White) {
				res[game.White] = game.Loss
				res[game.Black] = game.Win
				fmt.Println("Illegal White  Move")
				break
			}
			pos.Move(mW, game.White)
			PrintIf(verbose, fmt.Sprintf("White plays %s", mW.String()))
		} else if toMove == game.Black {
			PrintIf(verbose, pos.String())
			mB := black.ChooseMove(pos)
			// llegal move loses
			if !pos.Move(mB, game.Black) {
				res[game.White] = game.Win
				res[game.Black] = game.Loss
				fmt.Println("Illegal Black Move")
				break
			}
			pos.Move(mB, game.Black)
			PrintIf(verbose, fmt.Sprintf("Black plays %s", mB.String()))
		}
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
	PrintIf(verbose, "Final Position")
	PrintIf(verbose, pos.String())
	return res
}
