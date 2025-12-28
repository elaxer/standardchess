package result

import (
	"encoding/json"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Castling struct {
	Abstract
	move.Castling
}

func (r *Castling) CapturedPiece() chess.Piece {
	return nil
}

func (r *Castling) Move() chess.Move {
	return r.Castling
}

func (r *Castling) Validate() error {
	return validation.ValidateStruct(r, validation.Field(&r.Abstract))
}

func (r *Castling) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"move":            r.Castling.String(),
		"side":            r.Side(),
		"captured_piece":  nil,
		"board_new_state": r.BoardNewState(),
		"str":             r.String(),
	})
}

func (r *Castling) String() string {
	return fmt.Sprintf("%s%s", r.Castling, r.suffix())
}
