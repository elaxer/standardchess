package resolver

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
)

func UnresolveFrom(move move.Piece, board chess.Board) (chess.Position, error) {
	if err := move.ValidateStrict(); err != nil {
		return chess.NewEmptyPosition(), err
	}

	piece, err := board.Squares().FindByPosition(move.From)
	if err != nil {
		return chess.NewEmptyPosition(), err
	}
	if piece == nil {
		return chess.NewEmptyPosition(), fmt.Errorf("%w: piece not found", Err)
	}

	hasSamePiece, hasSameFile, hasSameRank := false, false, false
	for _, samePiece := range board.Squares().GetPieces(piece.Notation(), piece.Side()) {
		samePiecePosition := board.Squares().GetByPiece(samePiece)
		if samePiecePosition == move.From || !board.LegalMoves(samePiece).ContainsOne(move.To) {
			continue
		}

		hasSamePiece = true
		hasSameFile = hasSameFile || samePiecePosition.File == move.From.File
		hasSameRank = hasSameRank || samePiecePosition.Rank == move.From.Rank
		if hasSameFile && hasSameRank {
			break
		}
	}

	unresolvedFrom := chess.NewEmptyPosition()
	if hasSameRank || (hasSamePiece && !hasSameFile) {
		unresolvedFrom.File = move.From.File
	}
	if hasSameFile {
		unresolvedFrom.Rank = move.From.Rank
	}

	return unresolvedFrom, nil
}
