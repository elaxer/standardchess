package rule

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/state/state"
)

func Checkmate(board chess.Board, side chess.Side) chess.State {
	if Check(board, side) != nil && board.Moves(side).Cardinality() == 0 {
		return state.Checkmate
	}

	return nil
}
