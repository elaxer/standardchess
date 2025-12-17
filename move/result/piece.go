package result

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Piece struct {
	Abstract
	FromFull       position.Position `json:"from_full"`
	FromShortened  position.Position `json:"from_shortened"`
	ACapturedPiece chess.Piece       `json:"captured_piece"`
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
		validation.Field(&r.FromFull, validation.By(position.ValidationRuleIsEmpty)),
		validation.Field(&r.FromShortened),
	)
}

func (r Piece) captureString() string {
	if r.IsCapture() {
		return ""
	}

	return "x"
}
