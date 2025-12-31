package piece

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
)

// ErrCreate means an error during the piece creating process.
var ErrCreate = errors.New("cannot create piece")

var AllNotations = []string{
	NotationPawn,
	NotationRook,
	NotationKnight,
	NotationBishop,
	NotationQueen,
	NotationKing,
}

// New creates a new chess piece based on the provided notation and side.
// Returns nil if the piece cannot be created.
// See chess.PieceFactory for more details.
func New(notation string, side chess.Side) (chess.Piece, error) {
	switch notation {
	case NotationPawn:
		return NewPawn(side), nil
	case NotationRook:
		return NewRook(side), nil
	case NotationKnight:
		return NewKnight(side), nil
	case NotationBishop:
		return NewBishop(side), nil
	case NotationQueen:
		return NewQueen(side), nil
	case NotationKing:
		return NewKing(side), nil
	}

	return nil, fmt.Errorf("%w: unknown notation", ErrCreate)
}
