package mover

import (
	"fmt"

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

	pieceResult, err := movePiece(promotion.PieceMove, piece.NotationPawn, board)
	if err != nil {
		return nil, err
	}

	promotedPiece, err := piece.New(promotion.PromotedPieceNotation, board.Turn())
	if err != nil {
		return nil, err
	}

	promotedPiece.SetIsMoved(true)

	if err := board.Squares().PlacePiece(promotedPiece, promotion.To); err != nil {
		return nil, err
	}

	pieceResult.Abstract = newAbstractResult(board)

	return &result.Promotion{PieceMove: pieceResult, InputMove: *promotion}, nil
}

func UndoPromotion(promotion *result.Promotion, board chess.Board) error {
	if err := promotion.Validate(); err != nil {
		return err
	}

	p, err := board.Squares().FindByPosition(promotion.InputMove.To)
	if err != nil {
		return err
	}
	if p.Side() != promotion.MoveSide {
		return fmt.Errorf("%w: cannot find promoted piece", Err)
	}

	if err := board.Squares().PlacePiece(promotion.Captured, promotion.InputMove.To); err != nil {
		return err
	}

	pawn := piece.NewPawn(promotion.MoveSide)
	pawn.SetIsMoved(true)

	return board.Squares().PlacePiece(pawn, promotion.FromFull)
}
