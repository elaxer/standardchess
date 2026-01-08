package piecemove

import (
	"errors"

	"github.com/elaxer/chess"
)

type PieceMove struct {
	From chess.Position `json:"from"`
	To   chess.Position `json:"to"`
}

func NewPieceMove(from, to chess.Position) PieceMove {
	return PieceMove{from, to}
}

func (m PieceMove) Validate() error {
	if err := m.From.Validate(); err != nil {
		return err
	}
	if err := m.To.Validate(); err != nil {
		return err
	}
	// todo
	if !m.To.IsFull() {
		return errors.New("not full todo")
	}

	return nil
}
