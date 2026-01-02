package rule

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/state/state"
)

func Checkmate(board chess.Board, side chess.Side) chess.State {
	if Check(board, side) != nil && len(board.Moves(side)) == 0 {
		return state.Checkmate
	}

	return nil
}
