package mover

import (
	"github.com/elaxer/chess"

	"github.com/elaxer/standardchess/move/move"
	"github.com/elaxer/standardchess/move/result"
	"github.com/elaxer/standardchess/piece"
)

// Promotion это структура, реализующая интерфейс Mover для выполнения и проверки ходов превращения фигур.
// Она отвечает за логику, связанную с превращением пешки в другую фигуру на шахматной доске.
type Promotion struct {
}

func (m *Promotion) Make(promotion *move.Promotion, board chess.Board) (*result.Promotion, error) {
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
