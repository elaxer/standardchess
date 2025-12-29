package mover

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/result"
)

func MakeNormal(normal *move.Normal, board chess.Board) (*result.Normal, error) {
	if err := normal.Validate(); err != nil {
		return nil, err
	}

	pieceResult, err := movePiece(normal.Piece, normal.PieceNotation, board)
	if err != nil {
		return nil, err
	}

	pieceResult.Abstract = newAbstractResult(board)

	return &result.Normal{Piece: pieceResult, InputMove: *normal}, nil
}
