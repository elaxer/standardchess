package normal

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/piecemove"
)

var ErrUndoMove = errors.New("cannot undo normal move")

func MakeMove(move *Move, board chess.Board) (*MoveResult, error) {
	if err := move.Validate(); err != nil {
		return nil, err
	}

	pieceMoveResult, err := piecemove.MakeMove(move.PieceMove, move.PieceNotation, board)
	if err != nil {
		return nil, err
	}

	return &MoveResult{PieceMoveResult: pieceMoveResult, InputMove: *move}, nil
}

func UndoMove(move *MoveResult, board chess.Board) error {
	if err := move.Validate(); err != nil {
		return err
	}

	p, err := board.Squares().FindByPosition(move.InputMove.To)
	if err != nil {
		return err
	}
	if p == nil || p.Notation() != move.InputMove.PieceNotation || p.Color() != move.MoveSide {
		return fmt.Errorf("%w: cannot find moved piece", ErrUndoMove)
	}

	if err := board.Squares().PlacePiece(move.Captured, move.InputMove.To); err != nil {
		return err
	}

	if err := board.Squares().PlacePiece(p, move.FromFull); err != nil {
		return err
	}

	p.SetIsMoved(move.WasMoved)

	return nil
}
