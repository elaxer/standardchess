// Package castling contains code for validating,
// executing and cancelling castling moves on a chessboard.
package castling

import (
	"errors"
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/uci"
)

const (
	TypeShort CastlingType = true
	TypeLong  CastlingType = false
)

var (
	ErrInvalidNotation = errors.New("wrong castling notation")
	ErrInvalidUCI      = errors.New("uci move string is not castling")
)

type CastlingType bool

func FromNotation(str string) (CastlingType, error) {
	str = strings.TrimRight(str, "+#")

	parts := strings.Split(str, "-")
	if len(parts) < 2 || len(parts) > 3 {
		return false, ErrInvalidNotation
	}
	for _, part := range parts {
		if part != "0" && part != "o" && part != "O" {
			return false, ErrInvalidNotation
		}
	}

	return len(parts) == 2, nil
}

func FromUCI(uciStr string, squares *chess.Squares, turn chess.Color) (CastlingType, error) {
	uciMove, err := uci.MoveFromString(uciStr)
	if err != nil {
		return false, err
	}

	p, err := squares.FindByPosition(uciMove.From)
	if err != nil {
		return false, err
	}

	isValidPiece := p != nil && p.Notation() == piece.NotationKing && p.Color() == turn
	if !isValidPiece || uciMove.From != KingInitPosition(turn) {
		return false, ErrInvalidUCI
	}

	switch uciMove.To {
	case KingCastledPosition(TypeShort, turn):
		return TypeShort, nil
	case KingCastledPosition(TypeLong, turn):
		return TypeLong, nil
	}

	return false, ErrInvalidUCI
}

func (m CastlingType) IsShort() bool {
	return m == TypeShort
}

func (m CastlingType) IsLong() bool {
	return m == TypeLong
}

func (m CastlingType) String() string {
	switch m {
	case TypeShort:
		return "O-O"
	case TypeLong:
		return "O-O-O"
	default:
		panic("unknown catling type")
	}
}
