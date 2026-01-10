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

type sliding struct {
	*abstract
}

func (s *sliding) slide(
	from, direction chess.Position,
	squares *chess.Squares,
) iter.Seq[chess.Position] {
	return func(yield func(chess.Position) bool) {
		for move := range squares.IterByDirection(from, direction) {
			piece, err := squares.FindByPosition(move)
			canContinue := err == nil && piece == nil

			if (s.canMove(piece, s.color) && !yield(move)) || !canContinue {
				return
			}
		}
	}
}
