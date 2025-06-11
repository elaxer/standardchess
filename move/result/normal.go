package result

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/move/move"
	"github.com/elaxer/standardchess/piece"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Normal struct {
	Piece
	InputMove move.Normal `json:"input"`
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

func (r *Normal) String() string {
	from := r.FromShortened
	if from.IsEmpty() && r.IsCapture() && r.InputMove.PieceNotation == piece.NotationPawn {
		from.File = r.FromFull.File
	}

	return fmt.Sprintf("%s%s%s%s%s", r.InputMove.PieceNotation, from, r.captureString(), r.InputMove.To, r.suffix())
}
