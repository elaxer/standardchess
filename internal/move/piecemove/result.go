package piecemove

import (
	"errors"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/result"
	validation "github.com/go-ozzo/ozzo-validation"
)

type PieceMoveResult struct {
	*result.Abstract

	WasMoved      bool
	FromFull      chess.Position
	FromShortened chess.Position
	Captured      chess.Piece
}

func (r PieceMoveResult) CapturedPiece() chess.Piece {
	return r.Captured
}

func (r PieceMoveResult) IsCapture() bool {
	return r.Captured != nil
}

func (r PieceMoveResult) Validate() error {
	if r.Abstract == nil {
		return errors.New("sfoadi")
	}
	if r.Abstract.NewState == nil {
		return errors.New("sfor333333333adi")
	}
	return validation.ValidateStruct(&r,
		validation.Field(&r.FromFull, validation.By(chess.ValidationRulePositionIsEmpty)),
		validation.Field(&r.FromShortened),
	)
}

func (r PieceMoveResult) CaptureString() string {
	if !r.IsCapture() {
		return ""
	}

	return "x"
}
