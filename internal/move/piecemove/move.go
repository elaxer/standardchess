// Package piecemove contains low-level logic for moving pieces on the board.
package piecemove

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
)

var ErrMoveValidation = errors.New("piece move validation error")

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
	if !m.To.IsFull() {
		return fmt.Errorf("%w: to position is not full", ErrMoveValidation)
	}

	return nil
}
