package piecemove

import (
	"github.com/elaxer/chess"
	validation "github.com/go-ozzo/ozzo-validation"
)

type PieceMove struct {
	From chess.Position `json:"from"`
	To   chess.Position `json:"to"`
}

func NewPieceMove(from, to chess.Position) PieceMove {
	return PieceMove{from, to}
}

func (m PieceMove) Validate() error {
	return validation.ValidateStruct(
		&m,
		validation.Field(&m.From),
		validation.Field(&m.To, validation.By(chess.ValidationRulePositionIsEmpty)),
	)
}
