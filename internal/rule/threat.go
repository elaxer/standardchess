package rule

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/state"
)

func Check(board chess.Board) chess.State {
	_, kingPosition := board.Squares().FindPiece(piece.NotationKing, board.Turn())
	if board.IsSquareAttacked(kingPosition) {
		return state.Check
	}

	return nil
}
