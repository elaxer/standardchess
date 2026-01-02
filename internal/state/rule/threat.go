package rule

import (
	"slices"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/state/state"
)

func Check(board chess.Board, side chess.Side) chess.State {
	_, kingPosition := board.Squares().FindPiece(piece.NotationKing, side)
	if slices.Contains(board.Moves(!side), kingPosition) {
		return state.Check
	}

	return nil
}
