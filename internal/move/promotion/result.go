package promotion

import (
	"encoding/json"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/piecemove"
	validation "github.com/go-ozzo/ozzo-validation"
)

type MoveResult struct {
	piecemove.PieceMoveResult

	InputMove Move
}

func (r *MoveResult) Move() chess.Move {
	return &r.InputMove
}

func (r *MoveResult) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.PieceMoveResult),
		validation.Field(&r.InputMove),
	)
}

func (r *MoveResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"move":            r.InputMove.String(),
		"side":            r.Side(),
		"captured_piece":  r.CapturedPiece(),
		"board_new_state": r.BoardNewState(),
		"str":             r.String(),
	})
}

func (r *MoveResult) String() string {
	from := r.FromShortened
	if from.IsEmpty() && r.IsCapture() {
		from.File = r.FromFull.File
	}

	return fmt.Sprintf(
		"%s%s%s=%s%s",
		from,
		r.CaptureString(),
		r.InputMove.To,
		r.InputMove.PromotedPieceNotation,
		r.Suffix(),
	)
}
