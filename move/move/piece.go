package move

import (
	"github.com/elaxer/chess/position"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Piece struct {
	From position.Position `json:"from"`
	To   position.Position `json:"to"`
}

func NewPiece(from, to position.Position) Piece {
	return Piece{from, to}
}

func (m Piece) Validate() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(&m.From),
		validation.Field(&m.To, validation.By(position.ValidationRuleIsEmpty)),
	)
}

func (m Piece) ValidateStrict() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(&m.From, validation.By(position.ValidationRuleIsEmpty)),
		validation.Field(&m.To, validation.By(position.ValidationRuleIsEmpty)),
	)
}
