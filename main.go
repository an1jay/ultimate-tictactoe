package main

import (
	"fmt"
	"strconv"
)

func main() {

	testBBUpdateSubboardWins()

	testBBManyUpdateSubboardWins()

}

func strSquare(a Square) string {
	return a.String()
}

func makeBBFromString(str string) {
	num, _ := strconv.ParseUint(str, 2, 32)
	// fmt.Printf("%032b \n", num)
	fmt.Printf("%d \n", num)
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
	b := BitBoard{1224736768, 469762048 >> 9, 1224736768 >> 18}
	fmt.Println("Before")
	b.PrintRowsAsBits()
	b.UpdateSubBoardWins()

	fmt.Println("After")
	b.PrintRowsAsBits()
}

func testBBManyUpdateSubboardWins() {
	b := BitBoard{1224736768, 469762048 >> 9, 1224736768 >> 18}
	fmt.Println("Before")
	b.PrintRowsAsBits()
	b.UpdateSubBoardWins()
	b.UpdateSubBoardWins()

	fmt.Println("After")
	b.PrintRowsAsBits()
}

func testBBWin() {
	b := BitBoard{1224736768, 469762048 >> 9, 1224736768 >> 18}
	fmt.Println(b.Win())
}
func testBBMove() {
	b := BitBoard{1224736768, 469762048 >> 9, 1224736768 >> 18}
	fmt.Println("Before: ")
	b.PrintRowsAsBits()
	b.Move(Square(12))
	fmt.Println("After: ")
	b.PrintRowsAsBits()
}

func testBBDisplay() {
	b := BitBoard{1224736768, 469762048 >> 9, 1224736768 >> 18}
	b.Display()
}
