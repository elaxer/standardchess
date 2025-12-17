package result

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/move/move"
	validation "github.com/go-ozzo/ozzo-validation"
)

type EnPassant struct {
	Piece
	Input move.EnPassant `json:"input"`
}

func (r *EnPassant) Move() chess.Move {
	return &r.Input
}

func (r *EnPassant) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Piece),
		validation.Field(&r.ACapturedPiece, validation.NotNil),
		validation.Field(&r.Input),
	)
}

func (r *EnPassant) String() string {
	from := r.FromShortened
	if from.IsEmpty() {
		from.File = r.FromFull.File
	}

	return fmt.Sprintf("%sx%s%s", from, r.Input.To, r.suffix())
}
