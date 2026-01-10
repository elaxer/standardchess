package piecemove

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/result"
	"github.com/elaxer/standardchess/internal/resolver"
)

func MakeMove(
	move PieceMove,
	movingPieceNotation string,
	board chess.Board,
) (PieceMoveResult, error) {
	fullFrom, err := resolver.ResolveFrom(move.From, move.To, movingPieceNotation, board)
	if err != nil {
		return PieceMoveResult{}, err
	}
	pieceMove := NewPieceMove(fullFrom, move.To)

	shortenedFrom, err := resolver.UnresolveFrom(pieceMove.From, pieceMove.To, board)
	if err != nil {
		return PieceMoveResult{}, err
	}
	if err := ValidateMove(pieceMove, movingPieceNotation, board); err != nil {
		return PieceMoveResult{}, err
	}

	capturedPiece, err := board.Squares().MovePiece(pieceMove.From, pieceMove.To)
	if err != nil {
		return PieceMoveResult{}, err
	}

	piece, err := board.Squares().FindByPosition(pieceMove.To)
	if err != nil {
		return PieceMoveResult{}, err
	}

	wasMoved := piece.IsMoved()

	piece.SetIsMoved(true)

	return PieceMoveResult{
		WasMoved:      wasMoved,
		FromFull:      fullFrom,
		FromShortened: shortenedFrom,
		Captured:      capturedPiece,
		Abstract:      &result.Abstract{MoveSide: board.Turn()},
	}, nil
}
