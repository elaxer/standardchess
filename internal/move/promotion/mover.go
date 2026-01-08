package promotion

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"

	"github.com/elaxer/standardchess/internal/move/piecemove"
	"github.com/elaxer/standardchess/internal/piece"
)

var ErrUndo = errors.New("cannot undo promotion move")

// MakeMove отвечает за логику, связанную с превращением пешки в другую фигуру на шахматной доске.
func MakeMove(move *Move, board chess.Board) (*MoveResult, error) {
	if err := move.Validate(); err != nil {
		return nil, err
	}

	pieceResult, err := piecemove.MakeMove(move.PieceMove, piece.NotationPawn, board)
	if err != nil {
		return nil, err
	}

	promotedPiece, err := piece.New(move.PromotedPieceNotation, board.Turn())
	if err != nil {
		return nil, err
	}

	promotedPiece.SetIsMoved(true)

	if err := board.Squares().PlacePiece(promotedPiece, move.To); err != nil {
		return nil, err
	}

	return &MoveResult{PieceMoveResult: pieceResult, InputMove: *move}, nil
}

func UndoPromotion(move *MoveResult, board chess.Board) error {
	if err := move.Validate(); err != nil {
		return err
	}

	p, err := board.Squares().FindByPosition(move.InputMove.To)
	if err != nil {
		return err
	}
	if p.Color() != move.MoveSide {
		return fmt.Errorf("%w: cannot find promoted piece", ErrUndo)
	}

	if err := board.Squares().PlacePiece(move.Captured, move.InputMove.To); err != nil {
		return err
	}

	pawn := piece.NewPawn(move.MoveSide)
	pawn.SetIsMoved(true)

	return board.Squares().PlacePiece(pawn, move.FromFull)
}
