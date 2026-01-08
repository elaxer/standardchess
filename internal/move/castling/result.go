package castling

import (
	"encoding/json"
	"errors"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/result"
)

type MoveResult struct {
	*result.Abstract
	CastlingType

	InitRookPosition chess.Position
	InitKingPosition chess.Position
}

func (r *MoveResult) CapturedPiece() chess.Piece {
	return nil
}

func (r *MoveResult) Move() chess.Move {
	return r.CastlingType
}

func (r *MoveResult) Validate() error {
	if r.Abstract == nil || r.NewState == nil {
		return errors.New("aosdjfioj")
	}
	if !r.InitKingPosition.IsFull() {
		return errors.New("todo")
	}
	if !r.InitRookPosition.IsFull() {
		return errors.New("todo")
	}

	return nil
}

func (r *MoveResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"move":            r.CastlingType.String(),
		"side":            r.Side(),
		"captured_piece":  nil,
		"board_new_state": r.BoardNewState(),
		"str":             r.String(),
	})
}

func (r *MoveResult) String() string {
	return r.CastlingType.String() + r.Suffix()
}
