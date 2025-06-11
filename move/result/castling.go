package result

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/move/move"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Castling struct {
	Abstract
	move.Castling `json:"castling"`
}

func (r *Castling) Move() chess.Move {
	return r.Castling
}

func (r *Castling) Validate() error {
	return validation.ValidateStruct(r, validation.Field(&r.Abstract))
}

func (r *Castling) String() string {
	return fmt.Sprintf("%s%s", r.Castling, r.suffix())
}
