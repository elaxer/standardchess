package castling

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/result"
)

var ErrMoveResultValidation = errors.New("castling move result validation error")

type MoveResult struct {
	*result.Abstract
	CastlingType

	InitRookPosition chess.Position
	InitKingPosition chess.Position
}

func (r *MoveResult) CapturedPiece() chess.Piece {
	return nil
}

func (r *MoveResult) Input() string {
	return r.CastlingType.String()
}

func (r *MoveResult) Validate() error {
	if r.Abstract == nil {
		return fmt.Errorf("%w: empty abstract", ErrMoveResultValidation)
	}
	if err := r.Abstract.Validate(); err != nil {
		return fmt.Errorf("%w: %w", ErrMoveResultValidation, err)
	}
	if !r.InitKingPosition.IsFull() {
		return fmt.Errorf("%w: not full initial king position", ErrMoveResultValidation)
	}
	if !r.InitRookPosition.IsFull() {
		return fmt.Errorf("%w: not full initial rook position", ErrMoveResultValidation)
	}

	return nil
}

func (r *MoveResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"move":            r.CastlingType.String(),
		"side":            r.Side(),
		"captured_piece":  nil,
		"board_new_state": r.BoardNewState(),
		"str":             r.String(),
	})
}

func (r *MoveResult) String() string {
	return r.CastlingType.String() + r.Suffix()
}
