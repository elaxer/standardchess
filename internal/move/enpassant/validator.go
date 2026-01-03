package enpassant

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
)

var (
	ErrValidation               = errors.New("en passant move validation error")
	errValidationNoTargetSquare = fmt.Errorf("%w: no en passant target square", ErrValidation)
	errValidationNoPawn         = fmt.Errorf("%w: the moving pawn wasn't found", ErrValidation)
	errValidationWrongRank      = fmt.Errorf("%w: the pawn is in the wrong rank", ErrValidation)
	errValidationWrongFile      = fmt.Errorf("%w: the pawn is in the wrong file", ErrValidation)
)

func ValidateMove(from, to chess.Position, board chess.Board) error {
	enPassantPosition := EnPassantPosition(board)
	if enPassantPosition.IsEmpty() {
		return errValidationNoTargetSquare
	}

	p, err := board.Squares().FindByPosition(from)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrValidation, err)
	}
	if p == nil || p.Notation() != piece.NotationPawn {
		return errValidationNoPawn
	}
	if from.Rank != enPassantRank(board.Turn()) {
		return errValidationWrongRank
	}
	if absFile(from.File-enPassantPosition.File) != 1 {
		return errValidationWrongFile
	}

	return nil
}

func absFile(file chess.File) chess.File {
	if file < 0 {
		return -file
	}

	return file
}
