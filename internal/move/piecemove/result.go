package piecemove

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/result"
)

var ErrMoveResultValidation = errors.New("piece move result validation error")

type PieceMoveResult struct {
	*result.Abstract

	WasMoved      bool
	FromFull      chess.Position
	FromShortened chess.Position
	Captured      chess.Piece
}

func (r PieceMoveResult) CapturedPiece() chess.Piece {
	return r.Captured
}

func (r PieceMoveResult) IsCapture() bool {
	return r.Captured != nil
}

func (r PieceMoveResult) Validate() error {
	if r.Abstract == nil {
		return fmt.Errorf("%w: empty abstract result", ErrMoveResultValidation)
	}
	if err := r.Abstract.Validate(); err != nil {
		return err
	}
	if err := r.FromFull.Validate(); err != nil {
		return err
	}
	if !r.FromFull.IsFull() {
		return fmt.Errorf("%w: from full position is not full", ErrMoveResultValidation)
	}

	return r.FromShortened.Validate()
}

func (r PieceMoveResult) CaptureString() string {
	if !r.IsCapture() {
		return ""
	}

	return "x"
}
