package enpassant

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/piecemove"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Move struct {
	piecemove.PieceMove
}

func NewEnPassant(from, to chess.Position) *Move {
	return &Move{piecemove.NewPieceMove(from, to)}
}

func (m *Move) Validate() error {
	return validation.ValidateStruct(m, validation.Field(&m.PieceMove))
}

func (m *Move) String() string {
	return m.From.String() + m.To.String()
}
