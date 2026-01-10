// Package piece contains the logic of chess pieces,
// their variables and constants, and functions for creating pieces.
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

func New(notation string, color chess.Color) (chess.Piece, error) {
	switch notation {
	case NotationPawn:
		return NewPawn(color), nil
	case NotationRook:
		return NewRook(color), nil
	case NotationKnight:
		return NewKnight(color), nil
	case NotationBishop:
		return NewBishop(color), nil
	case NotationQueen:
		return NewQueen(color), nil
	case NotationKing:
		return NewKing(color), nil
	}

	return nil, fmt.Errorf("%w: unknown notation", ErrCreate)
}
