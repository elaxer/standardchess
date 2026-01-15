package enpassant

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
)

var (
	ErrValidation                  = errors.New("en passant move validation error")
	errValidationNoTargetSquare    = fmt.Errorf("%w: no en passant target square", ErrValidation)
	errValidationNoPawn            = fmt.Errorf("%w: the moving pawn wasn't found", ErrValidation)
	errValidationWrongRank         = fmt.Errorf("%w: the pawn is in the wrong rank", ErrValidation)
	errValidationWrongFile         = fmt.Errorf("%w: the pawn is in the wrong file", ErrValidation)
	errValidationWrongTargetSquare = fmt.Errorf(
		"%w: to position is not equal to the en passant target square",
		ErrValidation,
	)
	errValidationNoPawnToCapture = fmt.Errorf("%w: there is no pawn to capture", ErrValidation)
	errValidationCheckAfterMove  = fmt.Errorf("%w: there is a check after the move", ErrValidation)
)

func ValidateMove(from, to chess.Position, board chess.Board) error {
	if err := from.Validate(); err != nil {
		return err
	}
	if err := to.Validate(); err != nil {
		return err
	}

	enPassantPosition := EnPassantTargetSquare(board)
	if enPassantPosition.IsEmpty() {
		return errValidationNoTargetSquare
	}
	if enPassantPosition != to {
		return errValidationWrongTargetSquare
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

	return validateCheck(from, to, board)
}

func validateCheck(from, to chess.Position, board chess.Board) error {
	pawnToCapturePosition := chess.NewPosition(
		to.File,
		to.Rank-piece.PawnRankDirection(board.Turn()),
	)
	pawnToCapture, err := board.Squares().FindByPosition(pawnToCapturePosition)
	if err != nil {
		return err
	}
	if pawnToCapture == nil {
		return errValidationNoPawnToCapture
	}

	_, kingPosition := board.Squares().FindPiece(piece.NotationKing, board.Turn())

	var checkErr error = nil
	err = board.Squares().MovePieceTemporarily(from, to, func() {
		if err := board.Squares().PlacePiece(nil, pawnToCapturePosition); err != nil {
			panic(err)
		}

		if board.IsSquareAttacked(kingPosition) {
			checkErr = errValidationCheckAfterMove
		}

		if err := board.Squares().PlacePiece(pawnToCapture, pawnToCapturePosition); err != nil {
			panic(err)
		}
	})
	if err != nil {
		return err
	}

	return checkErr
}

func absFile(file chess.File) chess.File {
	if file < 0 {
		return -file
	}

	return file
}
