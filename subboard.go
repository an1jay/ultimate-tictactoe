package main

// SubBoard represents one of the 9 3*3 boards which comprise an Ultimate Tic Tac Toe board.
type SubBoard uint8

// Enumerating the possible subboards
const (
	NoSubBoard SubBoard = iota
	SubBoard0
	SubBoard1
	SubBoard2
	SubBoard3
	SubBoard4
	SubBoard5
	SubBoard6
	SubBoard7
	SubBoard8
)
