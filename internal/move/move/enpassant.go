package move

import (
	"fmt"

	"github.com/elaxer/chess"
	validation "github.com/go-ozzo/ozzo-validation"
)

type EnPassant struct {
	Piece
}

func NewEnPassant(from, to chess.Position) *EnPassant {
	return &EnPassant{NewPiece(from, to)}
}

func (m *EnPassant) Validate() error {
	return validation.ValidateStruct(m, validation.Field(&m.Piece))
}

func (m *EnPassant) String() string {
	return fmt.Sprintf("%s%s", m.From, m.To)
}
