package piece

import (
	"iter"

	"github.com/elaxer/chess"
)

var (
	orthogonalDirections = []chess.Position{
		chess.DirectionRight,
		chess.DirectionLeft,
		chess.DirectionTop,
		chess.DirectionBottom,
	}
	diagonalDirections = []chess.Position{
		chess.DirectionTopRight,
		chess.DirectionBottomLeft,
		chess.DirectionBottomRight,
		chess.DirectionTopLeft,
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
