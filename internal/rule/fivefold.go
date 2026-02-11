package rule

import (
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/castling"
	"github.com/elaxer/standardchess/internal/move/enpassant"
	"github.com/elaxer/standardchess/internal/state"
)

type fivefoldRepetition struct {
	hashes []string
}

func NewFivefoldRepetition() *fivefoldRepetition {
	return &fivefoldRepetition{make([]string, 0)}
}

func (r *fivefoldRepetition) Rule(board chess.Board) chess.State {
	if len(r.hashes) < 5 {
		return nil
	}

	currentHash := r.hashes[len(r.hashes)-1]
	repetitions := 0

	for _, h := range r.hashes {
		if h == currentHash {
			repetitions++
		}
	}

	if repetitions < 5 {
		return nil
	}

	return state.FivefoldRepetition
}

func (r *fivefoldRepetition) AfterBoardMakeMove(board chess.Board) {
	r.hashes = append(r.hashes, r.hash(board))
}

func (r *fivefoldRepetition) AfterBoardUndoMove(board chess.Board) {
	if len(r.hashes) > 0 {
		r.hashes = r.hashes[:len(r.hashes)-1]
	}
}

func (r *fivefoldRepetition) hash(board chess.Board) string {
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
