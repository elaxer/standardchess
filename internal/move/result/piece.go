package result

import (
	"github.com/elaxer/chess"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Piece struct {
	Abstract

	FromFull       chess.Position
	FromShortened  chess.Position
	ACapturedPiece chess.Piece
}

func (r Piece) CapturedPiece() chess.Piece {
	return r.ACapturedPiece
}

func (r Piece) IsCapture() bool {
	return r.ACapturedPiece != nil
}

func (r Piece) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Abstract),
		validation.Field(&r.FromFull, validation.By(chess.ValidationRulePositionIsEmpty)),
		validation.Field(&r.FromShortened),
	)
}

func (r Piece) captureString() string {
	if !r.IsCapture() {
		return ""
	}

	return "x"
}
