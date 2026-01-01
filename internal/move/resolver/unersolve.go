package resolver

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
)

var ErrUnresolve = errors.New("unresolving error")

func UnresolveFrom(from, to chess.Position, board chess.Board) (chess.Position, error) {
	if !from.IsFull() {
		return chess.NewPositionEmpty(), fmt.Errorf("%w: from position is not full", ErrUnresolve)
	}

	piece, err := board.Squares().FindByPosition(from)
	if err != nil {
		return chess.NewPositionEmpty(), err
	}
	if piece == nil {
		return chess.NewPositionEmpty(), fmt.Errorf("%w: piece not found", ErrResolve)
	}

	hasSamePiece, hasSameFile, hasSameRank := false, false, false
	for _, samePiece := range board.Squares().GetPieces(piece.Notation(), piece.Side()) {
		samePiecePosition := board.Squares().GetByPiece(samePiece)
		if samePiecePosition == from || !board.LegalMoves(samePiece).ContainsOne(to) {
			continue
		}

		hasSamePiece = true
		hasSameFile = hasSameFile || samePiecePosition.File == from.File
		hasSameRank = hasSameRank || samePiecePosition.Rank == from.Rank
		if hasSameFile && hasSameRank {
			break
		}
	}

	unresolvedFrom := chess.NewPositionEmpty()
	if hasSameRank || (hasSamePiece && !hasSameFile) {
		unresolvedFrom.File = from.File
	}
	if hasSameFile {
		unresolvedFrom.Rank = from.Rank
	}

	return unresolvedFrom, nil
}
