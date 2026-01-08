package normal

import (
	"encoding/json"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/piecemove"
	"github.com/elaxer/standardchess/internal/piece"
)

type MoveResult struct {
	piecemove.PieceMoveResult

	InputMove Move
}

func (r *MoveResult) Move() chess.Move {
	return &r.InputMove
}

func (r *MoveResult) Validate() error {
	if err := r.PieceMoveResult.Validate(); err != nil {
		return err
	}

	return r.InputMove.Validate()
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
	if from.IsEmpty() && r.IsCapture() && r.InputMove.PieceNotation == piece.NotationPawn {
		from.File = r.FromFull.File
	}

	return r.InputMove.PieceNotation + from.String() + r.CaptureString() + r.InputMove.To.String() + r.Suffix()
}
