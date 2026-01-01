package resolver

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
)

var ErrResolve = errors.New("resolving error")

// ResolveFrom определяет стартовую позицию фигуры, которая будет ходить.
func ResolveFrom(from, to chess.Position, pieceNotation string, board chess.Board, turn chess.Side) (chess.Position, error) {
	if from.IsFull() {
		return from, nil
	}
	if err := from.Validate(); err != nil {
		return chess.NewPositionEmpty(), err
	}
	if err := to.Validate(); err != nil {
		return chess.NewPositionEmpty(), err
	}

	pieces := make([]chess.Piece, 0, 8)
	for _, piece := range board.Squares().GetPieces(pieceNotation, turn) {
		if board.LegalMoves(piece).ContainsOne(to) {
			pieces = append(pieces, piece)
		}
	}

	if len(pieces) == 0 {
		return chess.NewPositionEmpty(), fmt.Errorf("%w: no moves found", ErrResolve)
	}
	if len(pieces) == 1 {
		return board.Squares().GetByPiece(pieces[0]), nil
	}

	resolvedFrom := from
	for _, piece := range pieces {
		pos := board.Squares().GetByPiece(piece)
		if from.Rank.IsNull() && pos.File == from.File {
			resolvedFrom.Rank = pos.Rank
		}
		if from.File.IsNull() && pos.Rank == from.Rank {
			resolvedFrom.File = pos.File
		}
	}

	return resolvedFrom, nil
}
