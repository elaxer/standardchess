package result

import (
	"encoding/json"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	validation "github.com/go-ozzo/ozzo-validation"
)

type EnPassant struct {
	Piece

	InputMove move.EnPassant
}

func (r *EnPassant) Move() chess.Move {
	return &r.InputMove
}

func (r *EnPassant) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Piece),
		validation.Field(&r.ACapturedPiece, validation.NotNil),
		validation.Field(&r.InputMove),
	)
}

func (r *EnPassant) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"move":            r.InputMove.String(),
		"side":            r.Side(),
		"captured_piece":  r.CapturedPiece(),
		"board_new_state": r.BoardNewState(),
		"str":             r.String(),
	})
}

func (r *EnPassant) String() string {
	from := r.FromShortened
	if from.IsEmpty() {
		from.File = r.FromFull.File
	}

	return fmt.Sprintf("%sx%s%s", from, r.InputMove.To, r.suffix())
}
