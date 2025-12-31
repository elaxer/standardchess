package mover

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/result"
)

func MakeNormal(normal *move.Normal, board chess.Board) (*result.Normal, error) {
	if err := normal.Validate(); err != nil {
		return nil, err
	}

	pieceResult, err := movePiece(normal.PieceMove, normal.PieceNotation, board)
	if err != nil {
		return nil, err
	}

	pieceResult.Abstract = newAbstractResult(board)

	return &result.Normal{PieceMove: pieceResult, InputMove: *normal}, nil
}

func UndoNormal(normal *result.Normal, board chess.Board) error {
	if err := normal.Validate(); err != nil {
		return err
	}

	p, err := board.Squares().FindByPosition(normal.InputMove.To)
	if err != nil {
		return err
	}
	if p == nil || p.Notation() != normal.InputMove.PieceNotation || p.Side() != normal.MoveSide {
		return fmt.Errorf("%w: cannot find moved piece", ErrUndoMove)
	}

	if err := board.Squares().PlacePiece(normal.Captured, normal.InputMove.To); err != nil {
		return err
	}

	if err := board.Squares().PlacePiece(p, normal.FromFull); err != nil {
		return err
	}

	p.SetIsMoved(normal.WasMoved)

	return nil
}
