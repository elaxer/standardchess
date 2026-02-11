package rule

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/check"
	"github.com/elaxer/standardchess/internal/state"
)

func Checkmate(board chess.Board) chess.State {
	if check.IsCheck(board) && len(board.Moves()) == 0 {
		return state.Checkmate
	}

	return nil
}

func Stalemate(board chess.Board) chess.State {
	if !check.IsCheck(board) && len(board.Moves()) == 0 {
		return state.Stalemate
	}

	return nil
}
