package move

import (
	"github.com/elaxer/chess"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Piece struct {
	From chess.Position `json:"from"`
	To   chess.Position `json:"to"`
}

func NewPiece(from, to chess.Position) Piece {
	return Piece{from, to}
}

func (m Piece) Validate() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(&m.From),
		validation.Field(&m.To, validation.By(chess.ValidationRulePositionIsEmpty)),
	)
}

func (m Piece) ValidateStrict() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(&m.From, validation.By(chess.ValidationRulePositionIsEmpty)),
		validation.Field(&m.To, validation.By(chess.ValidationRulePositionIsEmpty)),
	)
}
