package enpassant

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/piecemove"
)

type Move struct {
	piecemove.PieceMove
}

func NewEnPassant(from, to chess.Position) *Move {
	return &Move{piecemove.NewPieceMove(from, to)}
}

func (m *Move) Validate() error {
	return m.PieceMove.Validate()
}

func (m *Move) String() string {
	return m.From.String() + m.To.String()
}
