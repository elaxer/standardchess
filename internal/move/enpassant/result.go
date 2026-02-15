package enpassant

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/elaxer/standardchess/internal/move/piecemove"
)

var ErrMoveResultValidation = errors.New("en passant move result validation error")

type MoveResult struct {
	piecemove.PieceMoveResult

	InputMove Move
}

func (r *MoveResult) Input() string {
	return r.InputMove.String()
}

func (r *MoveResult) String() string {
	from := r.FromShortened
	if from.IsEmpty() {
		from.File = r.FromFull.File
	}

	return from.String() + "x" + r.InputMove.To.String() + r.Suffix()
}

func (r *MoveResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"input":           r.InputMove.String(),
		"side":            r.Side(),
		"captured_piece":  r.CapturedPiece(),
		"board_new_state": r.BoardNewState(),
		"uci":             r.uci(),
		"str":             r.String(),
	})
}

func (r *MoveResult) uci() string {
	return r.FromFull.String() + r.InputMove.To.String()
}

func (r *MoveResult) validate() error {
	if err := r.Validate(); err != nil {
		return err
	}
	if r.Captured == nil {
		return fmt.Errorf("%w: must have a captured piece", ErrMoveResultValidation)
	}

	return r.InputMove.Validate()
}
