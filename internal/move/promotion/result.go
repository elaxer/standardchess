package promotion

import (
	"encoding/json"
	"strings"

	"github.com/elaxer/standardchess/internal/move/piecemove"
)

type MoveResult struct {
	piecemove.PieceMoveResult

	InputMove Move
}

func (r *MoveResult) Input() string {
	return r.InputMove.String()
}

func (r *MoveResult) String() string {
	from := r.FromShortened
	if from.IsEmpty() && r.IsCapture() {
		from.File = r.FromFull.File
	}

	return from.String() + r.CaptureString() + r.InputMove.To.String() +
		"=" + r.InputMove.PromotedPieceNotation + r.Suffix()
}

func (r *MoveResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"input":           r.InputMove.String(),
		"side":            r.Side(),
		"captured_piece":  r.CapturedPiece(),
		"board_new_state": r.BoardNewState(),
		"uci":             r.uci(),
		"str":             r.String(),
	})
}

func (r *MoveResult) uci() string {
	return r.FromFull.String() + r.InputMove.To.String() + strings.ToLower(r.InputMove.PromotedPieceNotation)
}

func (r *MoveResult) validate() error {
	if err := r.Validate(); err != nil {
		return err
	}

	return r.InputMove.validate()
}
