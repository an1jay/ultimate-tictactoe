package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println("Ultimate ")

	// x := Square(76)
	// fmt.Printf("%v", x)
	// b := BitBoard{0, 0, 0}
	// _ = b.UpdateSubboardWins()

	makeBBFromString("11111111100000000000000000000000")
	makeBBFromString("00000000011111111100000000000000")
	makeBBFromString("00000000000000000011111111100000")
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
