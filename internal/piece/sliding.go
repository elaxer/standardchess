package piece

import (
	"iter"

	"github.com/elaxer/chess"
)

var (
	orthogonalDirections = []chess.Position{
		chess.NewPosition(1, 0),  // Right
		chess.NewPosition(-1, 0), // Left
		chess.NewPosition(0, 1),  // Up
		chess.NewPosition(0, -1), // Down
	}
	diagonalDirections = []chess.Position{
		chess.NewPosition(1, 1),   // Up-Right
		chess.NewPosition(-1, -1), // Down-Left
		chess.NewPosition(1, -1),  // Down-Right
		chess.NewPosition(-1, 1),  // Up-Left
	}
	allDirections = append(orthogonalDirections, diagonalDirections...)
)

// sliding - структура для фигур, которые могут двигаться по диагонали или вертикали/горизонтали
// (слон, ферзь, ладья).
// Она содержит базовую структуру фигуры и методы для проверки возможности движения.
type sliding struct {
	*abstract
}

func (s *sliding) slide(from, direction chess.Position, squares *chess.Squares) iter.Seq[chess.Position] {
	return func(yield func(chess.Position) bool) {
		for move := range squares.IterByDirection(from, direction) {
			piece, err := squares.FindByPosition(move)
			canContinue := err == nil && piece == nil

			if (s.canMove(piece, s.side) && !yield(move)) || !canContinue {
				return
			}
		}
	}
}
