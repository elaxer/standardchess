package rule

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/state/state"
)

func Checkmate(board chess.Board) chess.State {
	if Check(board) != nil && len(board.Moves()) == 0 {
		return state.Checkmate
	}

	return nil
}
