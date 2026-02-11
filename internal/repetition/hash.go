package repetition

import (
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/castling"
	"github.com/elaxer/standardchess/internal/move/enpassant"
)

func hash(board chess.Board) string {
	var hash strings.Builder

	hash.WriteString(board.Turn().String())

	for _, row := range board.Squares().IterOverRows(false) {
		for _, piece := range row {
			if piece == nil {
				hash.WriteRune('.')
			} else {
				hash.WriteString(piece.Notation() + piece.Color().String())
			}
		}
	}

	if castling.ValidateMoveWithObstacle(castling.TypeShort, chess.ColorWhite, board) != nil {
		hash.WriteRune('K')
	} else {
		hash.WriteRune('.')
	}
	if castling.ValidateMoveWithObstacle(castling.TypeLong, chess.ColorWhite, board) != nil {
		hash.WriteRune('Q')
	} else {
		hash.WriteRune('.')
	}
	if castling.ValidateMoveWithObstacle(castling.TypeShort, chess.ColorBlack, board) != nil {
		hash.WriteRune('k')
	} else {
		hash.WriteRune('.')
	}
	if castling.ValidateMoveWithObstacle(castling.TypeLong, chess.ColorBlack, board) != nil {
		hash.WriteRune('q')
	} else {
		hash.WriteRune('.')
	}

	if position := enpassant.EnPassantTargetSquare(board); !position.IsEmpty() {
		hash.WriteString(position.String())
	} else {
		hash.WriteRune('.')
	}

	return hash.String()
}
