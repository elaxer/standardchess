package result

import (
	"encoding/json"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Promotion struct {
	PieceMove

	InputMove move.Promotion
}

func (r *Promotion) Move() chess.Move {
	return &r.InputMove
}

func (r *Promotion) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.PieceMove),
		validation.Field(&r.InputMove),
	)
}

func (r *Promotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"move":            r.InputMove.String(),
		"side":            r.Side(),
		"captured_piece":  r.CapturedPiece(),
		"board_new_state": r.BoardNewState(),
		"str":             r.String(),
	})
}

func (r *Promotion) String() string {
	from := r.FromShortened
	if from.IsEmpty() && r.IsCapture() {
		from.File = r.FromFull.File
	}

	return fmt.Sprintf("%s%s%s=%s%s", from, r.captureString(), r.InputMove.To, r.InputMove.PromotedPieceNotation, r.suffix())
}
