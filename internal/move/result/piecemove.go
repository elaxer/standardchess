package result

import (
	"github.com/elaxer/chess"
	validation "github.com/go-ozzo/ozzo-validation"
)

type PieceMove struct {
	Abstract

	WasMoved      bool
	FromFull      chess.Position
	FromShortened chess.Position
	Captured      chess.Piece
}

func (r PieceMove) CapturedPiece() chess.Piece {
	return r.Captured
}

func (r PieceMove) IsCapture() bool {
	return r.Captured != nil
}

func (r PieceMove) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Abstract),
		validation.Field(&r.FromFull, validation.By(chess.ValidationRulePositionIsEmpty)),
		validation.Field(&r.FromShortened),
	)
}

func (r PieceMove) captureString() string {
	if !r.IsCapture() {
		return ""
	}

	return "x"
}
