package castling

import (
	"encoding/json"
	"errors"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/result"
	validation "github.com/go-ozzo/ozzo-validation"
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
	if r.Abstract == nil || r.Abstract.NewState == nil {
		return errors.New("aosdjfioj")
	}

	return validation.ValidateStruct(
		r,
		validation.Field(&r.InitKingPosition, validation.By(chess.ValidationRulePositionIsEmpty)),
		validation.Field(&r.InitRookPosition, validation.By(chess.ValidationRulePositionIsEmpty)),
	)
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
