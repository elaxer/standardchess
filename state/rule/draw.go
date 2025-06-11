package rule

import (
	"slices"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/move/result"
	"github.com/elaxer/standardchess/piece"
	"github.com/elaxer/standardchess/state/state"
)

func Stalemate(board chess.Board, side chess.Side) chess.State {
	if Check(board, side) == nil && board.Moves(side).Cardinality() == 0 {
		return state.Stalemate
	}

	return nil
}

func FiftyMoves(board chess.Board, side chess.Side) chess.State {
	moves := slices.Clone(board.MovesHistory())
	slices.Reverse(moves)

	count := 0
	for _, move := range moves {
		normalMove, ok := move.(*result.Normal)
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
