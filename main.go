package main

import (
	"fmt"
	"strconv"

	"github.com/an1jay/ultimate-tictactoe/game"
	"github.com/an1jay/ultimate-tictactoe/players"
	"github.com/an1jay/ultimate-tictactoe/players/evaluators"
)

func main() {
	p1 := players.AlphaBetaPlayer{
		Ev:       evaluators.BigBoardMSE{},
		MaxDepth: 8,
	}
	p2 := players.AlphaBetaPlayer{
		Ev:       evaluators.BigBoardMSE{},
		MaxDepth: 8,
	}

	// p1 := players.RandomPlayer{}
	// p2 := players.RandomPlayer{}
	// fmt.Println(playMatch(&p1, &p2, 10000))
	g := Game{}
	fmt.Println(g.PlayGame(&p1, &p2, true))
	fmt.Printf("Move List: \n %v", g.moveHistory)

}

func playMatch(p1, p2 Player, numGames uint32) []map[game.Result]uint32 {
	fmt.Println("Playing a match")
	g := Game{}
	resultsp1 := map[game.Result]uint32{}
	resultsp2 := map[game.Result]uint32{}
	for i := 0; i < int(numGames/2); i++ {
		// fmt.Println("Playing a game")
		switch g.PlayGame(p1, p2, false)[game.White] {
		case game.Win:
			resultsp1[game.Win]++
		case game.Loss:
			resultsp1[game.Loss]++
		case game.Tie:
			resultsp1[game.Tie]++
		}
	}
	for i := 0; i < int(numGames/2); i++ {
		// fmt.Println("Playing a game")
		switch g.PlayGame(p2, p1, false)[game.White] {
		case game.Win:
			resultsp2[game.Win.Other()]++
		case game.Loss:
			resultsp2[game.Loss.Other()]++
		case game.Tie:
			resultsp2[game.Tie]++
		}
	}
	results := []map[game.Result]uint32{resultsp1, resultsp2}
	return results

}

func strSquare(a game.Square) string {
	return a.String()
}

func makeBBFromString(str string) uint32 {
	num, _ := strconv.ParseUint(str, 2, 32)
	// fmt.Printf("%032b \n", num)
	fmt.Printf("%d \n", num)
	return uint32(num)
}

func calcWinConditions() {
	makeBBFromString("11100000000000000000000000000000")
	makeBBFromString("00011100000000000000000000000000")
	makeBBFromString("00000011100000000000000000000000")
	makeBBFromString("10010010000000000000000000000000")
	makeBBFromString("01001001000000000000000000000000")
	makeBBFromString("00100100100000000000000000000000")
	makeBBFromString("10001000100000000000000000000000")
	makeBBFromString("00101010000000000000000000000000")
}

func testBBUpdateSubboardWins() {
	b := game.BitBoard{Zeroth: 0, First: 469762048 >> 9, Second: 1<<20 + 1<<19 + 1<<18}
	fmt.Println("Before")
	b.PrintRowsAsBits()
	b.UpdateSubBoardWins()

	fmt.Println("After")
	b.PrintRowsAsBits()
}

func testBBManyUpdateSubboardWins() {
	b := game.BitBoard{Zeroth: 1224736768, First: 69762048 >> 9, Second: 1224736768 >> 18}
	fmt.Println("Before")
	b.PrintRowsAsBits()
	b.UpdateSubBoardWins()
	b.UpdateSubBoardWins()

	fmt.Println("After")
	b.PrintRowsAsBits()
}

func testBBWin() {
	b := game.BitBoard{Zeroth: 1224736768, First: 69762048 >> 9, Second: 1224736768 >> 18}
	fmt.Println(b.Win())
}
func testBBMove() {
	b := game.BitBoard{Zeroth: 1224736768, First: 69762048 >> 9, Second: 1224736768 >> 18}
	fmt.Println("Before: ")
	b.PrintRowsAsBits()
	b.Move(game.Square(12))
	fmt.Println("After: ")
	b.PrintRowsAsBits()
}

func testBBDisplay() {
	b := game.BitBoard{Zeroth: 1224736768, First: 69762048 >> 9, Second: 1224736768 >> 18}
	b.Display()
}

func testPLegalMoves() {
	p := game.Position{
		WhiteBB: game.BitBoard{
			Zeroth: makeBBFromString("11100000000000000000000000000000"),
			First:  makeBBFromString("00011000000000000000000000000000"),
			Second: makeBBFromString("10010010000000000000000000000000")},
		BlackBB:              game.BitBoard{},
		SideToMove:           game.White,
		SubBoardToPlayOnNext: game.NoSubBoard,
	}
	fmt.Println(p.LegalMoves())
}
