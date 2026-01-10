// Package resolver contains set of functions for obtaining
// the complete initial position of a move or, conversely,
// trimming unnecessary information about the initial position of a move.
//
// An example of resolving: if, in the starting position of the board,
// we take the move "e4" where the initial position is empty,
// the resolver will find the initial position, which will be "e2".
//
// Example of an resolver: in the starting position of the board,
// it takes the move "e2e4" and cuts off the unnecessary "e2",
// since in the starting position,
// White has only one pawn that can move to the "e4" square.
package resolver

import (
	"errors"
	"fmt"
	"slices"

	"github.com/elaxer/chess"
)

var ErrResolve = errors.New("resolving error")

func ResolveFrom(
	from, to chess.Position,
	pieceNotation string,
	board chess.Board,
) (chess.Position, error) {
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
	for piece := range board.Squares().GetPieces(pieceNotation, board.Turn()) {
		if slices.Contains(board.LegalMoves(piece), to) {
			pieces = append(pieces, piece)
		}
	}

	if len(pieces) == 0 {
		return chess.NewPositionEmpty(), fmt.Errorf("%w: no moves found", ErrResolve)
	}
	if len(pieces) == 1 {
		return board.Squares().GetByPiece(pieces[0]), nil
	}

	resolvedFrom := resolvedFrom(from, board, pieces)

	if !resolvedFrom.IsFull() {
		return chess.NewPositionEmpty(), fmt.Errorf(
			"%w: failed to resolve init position",
			ErrResolve,
		)
	}

	return resolvedFrom, nil
}

func resolvedFrom(
	unresolvedFrom chess.Position,
	board chess.Board,
	pieces []chess.Piece,
) chess.Position {
	resolvedFrom := unresolvedFrom
	for _, piece := range pieces {
		pos := board.Squares().GetByPiece(piece)
		if unresolvedFrom.Rank.IsNull() && pos.File == unresolvedFrom.File {
			resolvedFrom.Rank = pos.Rank
		}
		if unresolvedFrom.File.IsNull() && pos.Rank == unresolvedFrom.Rank {
			resolvedFrom.File = pos.File
		}
	}

	return resolvedFrom
}
