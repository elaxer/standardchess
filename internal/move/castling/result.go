package castling

import (
	"encoding/json"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/result"
	validation "github.com/go-ozzo/ozzo-validation"
)

type MoveResult struct {
	result.Abstract
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
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Abstract),
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
	return fmt.Sprintf("%s%s", r.CastlingType, r.Suffix())
}
