package players

import (
	"fmt"
	"time"

	"github.com/an1jay/ultimate-tictactoe/game"
	"github.com/an1jay/ultimate-tictactoe/players/evaluators"
)

// MinimaxPlayer plays UTTT randomly
type MinimaxPlayer struct {
	Ev        evaluators.Evaluator
	MaxDepth  uint8
	nodeCount uint64
}

// ChooseMove asks a random player to move randomly from choice of legal moves.
func (m *MinimaxPlayer) ChooseMove(pos game.Position) game.Square {

	fmt.Println("Minimax Player thinks ...")
	m.nodeCount = 0
	t0 := time.Now()
	bestScore := pos.SideToMove.Other().EvaluationWin() // -1 for white, 1 for black
	var bestMove game.Square

	switch pos.SideToMove {
	case game.White:
		for _, lgm := range pos.LegalMoves() {
			newPos := pos.Copy()
			newPos.Move(lgm, game.White)
			scr := m.Minimax(m.MaxDepth, game.Black, newPos) // Is Black <TODO> ?
			if scr == game.White.EvaluationWin() {
				return lgm
			}
			if scr >= bestScore {
				bestScore = scr
				bestMove = lgm
			}
		}
	case game.Black:
		for _, lgm := range pos.LegalMoves() {
			newPos := pos.Copy()
			newPos.Move(lgm, game.Black)
			scr := m.Minimax(m.MaxDepth, game.White, newPos) // Is White <TODO> ?
			if scr == game.Black.EvaluationWin() {
				return lgm
			}
			if scr <= bestScore {
				bestScore = scr
				bestMove = lgm
			}
		}
	}
	dt := time.Since(t0).Seconds()
	fmt.Printf("Minimax Player explore %d nodes in %.02f seconds at %.02f nodes/s \n", m.nodeCount, dt, float64(m.nodeCount)/dt)
	return bestMove
}

// DefaultVal is larger than maximum possible evaluation
const DefaultVal float32 = 1000

// Minimax returns a minimax score for a given evaluator and position
func (m *MinimaxPlayer) Minimax(depth uint8, side game.Color, pos *game.Position) float32 {

	m.nodeCount++

	// if at a terminal node, evaluate
	igo, _ := pos.GameOver()
	if igo || depth == 0 {
		return m.Ev.Evaluate(pos)
	}

	value := DefaultVal

	// recursively find Minimax value of a position
	switch side {
	case game.White:
		value *= -1
		for _, lgm := range pos.LegalMoves() {
			newPos := pos.Copy()
			newPos.Move(lgm, game.White)
			value = max(value, m.Minimax(depth-1, game.Black, newPos))
		}
	case game.Black:
		value *= 1
		for _, lgm := range pos.LegalMoves() {
			newPos := pos.Copy()
			newPos.Move(lgm, game.Black)
			value = min(value, m.Minimax(depth-1, game.White, newPos))
		}
	}
	return value
}

func max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}

func min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}
