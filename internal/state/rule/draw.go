package rule

import (
	"slices"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/normal"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/state/state"
)

func Stalemate(board chess.Board, side chess.Side) chess.State {
	if Check(board, side) == nil && board.Moves(side).Cardinality() == 0 {
		return state.Stalemate
	}

	return nil
}

func FiftyMoves(board chess.Board, side chess.Side) chess.State {
	moves := slices.Clone(board.MoveHistory())
	slices.Reverse(moves)

	count := 0
	for _, move := range moves {
		normalMove, ok := move.(*normal.MoveResult)
		if !ok || normalMove.InputMove.PieceNotation == piece.NotationPawn || normalMove.IsCapture() {
			count = 0
		} else {
			count++
		}
	}

	if count >= 50 {
		return state.DrawFiftyMoves
	}

	return nil
}
