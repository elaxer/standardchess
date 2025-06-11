package piece

import (
	"iter"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
)

var (
	orthogonalDirections = []position.Position{
		position.New(1, 0),  // Right
		position.New(-1, 0), // Left
		position.New(0, 1),  // Up
		position.New(0, -1), // Down
	}
	diagonalDirections = []position.Position{
		position.New(1, 1),   // Up-Right
		position.New(-1, -1), // Down-Left
		position.New(1, -1),  // Down-Right
		position.New(-1, 1),  // Up-Left
	}
	allDirections = append(orthogonalDirections, diagonalDirections...)
)

// sliding - структура для фигур, которые могут двигаться по диагонали или вертикали/горизонтали
// (слон, ферзь, ладья).
// Она содержит базовую структуру фигуры и методы для проверки возможности движения.
type sliding struct {
	*abstract
}

func (s *sliding) slide(from, direction position.Position, squares *chess.Squares) iter.Seq[position.Position] {
	return func(yield func(position position.Position) bool) {
		for move := range squares.IterByDirection(from, direction) {
			piece, err := squares.FindByPosition(move)
			canContinue := err == nil && piece == nil

			if (s.canMove(piece, s.side) && !yield(move)) || !canContinue {
				return
			}
		}
	}
}
