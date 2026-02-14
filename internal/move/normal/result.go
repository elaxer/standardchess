package normal

import (
	"encoding/json"

	"github.com/elaxer/standardchess/internal/move/piecemove"
	"github.com/elaxer/standardchess/internal/piece"
)

type MoveResult struct {
	piecemove.PieceMoveResult

	InputMove Move
}

func (r *MoveResult) Input() string {
	return r.InputMove.String()
}

func (r *MoveResult) UCI() string {
	return r.FromFull.String() + r.InputMove.To.String()
}

func (r *MoveResult) String() string {
	from := r.FromShortened
	if from.IsEmpty() && r.IsCapture() && r.InputMove.PieceNotation == piece.NotationPawn {
		from.File = r.FromFull.File
	}

	return r.InputMove.PieceNotation + from.String() + r.CaptureString() + r.InputMove.To.String() + r.Suffix()
}

func (r *MoveResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"input":           r.InputMove.String(),
		"side":            r.Side(),
		"captured_piece":  r.CapturedPiece(),
		"board_new_state": r.BoardNewState(),
		"uci":             r.UCI(),
		"str":             r.String(),
	})
}

func (r *MoveResult) validate() error {
	if err := r.Validate(); err != nil {
		return err
	}

	return r.InputMove.validate()
}
