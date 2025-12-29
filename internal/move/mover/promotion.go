package mover

import (
	"github.com/elaxer/chess"

	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/result"
	"github.com/elaxer/standardchess/internal/piece"
)

// MakePromotion отвечает за логику, связанную с превращением пешки в другую фигуру на шахматной доске.
func MakePromotion(promotion *move.Promotion, board chess.Board) (*result.Promotion, error) {
	if err := promotion.Validate(); err != nil {
		return nil, err
	}

	pieceResult, err := movePiece(promotion.Piece, piece.NotationPawn, board)
	if err != nil {
		return nil, err
	}

	promotedPiece, err := piece.New(promotion.PromotedPieceNotation, board.Turn())
	if err != nil {
		return nil, err
	}

	promotedPiece.MarkMoved()

	if err := board.Squares().PlacePiece(promotedPiece, promotion.To); err != nil {
		return nil, err
	}

	pieceResult.Abstract = newAbstractResult(board)

	return &result.Promotion{Piece: pieceResult, InputMove: *promotion}, nil
}
