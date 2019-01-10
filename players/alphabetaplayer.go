package players

import (
	"fmt"
	"time"

	"github.com/an1jay/ultimate-tictactoe/game"
	"github.com/an1jay/ultimate-tictactoe/players/evaluators"
)

// AlphaBetaPlayer plays UTTT using minimax and alpha beta pruning
type AlphaBetaPlayer struct {
	Ev        evaluators.Evaluator
	MaxDepth  uint8
	nodeCount uint64
}

// ChooseMove asks a AlphaBetaPlayer to play using minimax and alpha beta pruning
func (a *AlphaBetaPlayer) ChooseMove(pos game.Position) game.Square {

	fmt.Println("AlphaBeta Player thinks ...")
	a.nodeCount = 0
	t0 := time.Now()
	bestScore := pos.SideToMove.Other().EvaluationWin() // -1 for white, 1 for black
	var bestMove game.Square

	alpha := -DefaultVal
	beta := DefaultVal

	switch pos.SideToMove {
	case game.White:
		for _, lgm := range pos.LegalMoves() {
			newPos := pos.Copy()
			newPos.Move(lgm, game.White)
			scr := a.AlphaBeta(a.MaxDepth, game.Black, alpha, beta, newPos) // Is Black <TODO> ?
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
			scr := a.AlphaBeta(a.MaxDepth, game.White, alpha, beta, newPos) // Is White <TODO> ?
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
	fmt.Printf("AlphaBeta Player explore %d nodes in %.02f seconds at %.02f nodes/s \n", a.nodeCount, dt, float64(a.nodeCount)/dt)
	return bestMove
}

// AlphaBeta returns a minimax score for a given evaluator and position using alpha beta pruning
func (a *AlphaBetaPlayer) AlphaBeta(depth uint8, side game.Color, alpha float32, beta float32, pos *game.Position) float32 {

	a.nodeCount++

	// if at a terminal node, evaluate
	igo, _ := pos.GameOver()
	if igo || depth == 0 {
		return a.Ev.Evaluate(pos)
	}

	value := DefaultVal

	// recursively find Minimax value of a position
	switch side {
	case game.White:
		value *= -1
		for _, lgm := range pos.LegalMoves() {
			newPos := pos.Copy()
			newPos.Move(lgm, game.White)
			value = max(value, a.AlphaBeta(depth-1, game.Black, alpha, beta, newPos))
			alpha = max(alpha, value)
			if alpha >= beta {
				break
			}
		}
	case game.Black:
		value *= 1
		for _, lgm := range pos.LegalMoves() {
			newPos := pos.Copy()
			newPos.Move(lgm, game.Black)
			value = min(value, a.AlphaBeta(depth-1, game.White, alpha, beta, newPos))
			beta = min(beta, value)
			if alpha >= beta {
				break
			}
		}
	}
	return value
}
