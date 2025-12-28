package result

import (
	"encoding/json"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/piece"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Normal struct {
	Piece
	InputMove move.Normal
}

func (r *Normal) Move() chess.Move {
	return &r.InputMove
}

func (r *Normal) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Piece),
		validation.Field(&r.InputMove),
	)
}

func (r *Normal) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"move":            r.InputMove.String(),
		"side":            r.Side(),
		"captured_piece":  r.CapturedPiece(),
		"board_new_state": r.BoardNewState(),
		"str":             r.String(),
	})
}

func (r *Normal) String() string {
	from := r.FromShortened
	if from.IsEmpty() && r.IsCapture() && r.InputMove.PieceNotation == piece.NotationPawn {
		from.File = r.FromFull.File
	}

	return fmt.Sprintf("%s%s%s%s%s", r.InputMove.PieceNotation, from, r.captureString(), r.InputMove.To, r.suffix())
}
