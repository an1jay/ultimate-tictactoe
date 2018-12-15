package main

// Position represents a game state
type Position struct {
	board             Board
	sideToMove        Color
	boardToPlayOnNext SubBoard
}
