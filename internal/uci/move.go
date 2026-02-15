package uci

import (
	"errors"
	"slices"
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
)

type Move struct {
	From                  chess.Position
	To                    chess.Position
	PromotedPieceNotation string
}

func MoveFromString(uci string) (*Move, error) {
	if len(uci) < 4 {
		return nil, errors.New("too short string")
	}
	if len(uci) > 5 {
		return nil, errors.New("too long string")
	}

	move := &Move{From: chess.PositionFromString(uci[:2])}
	if move.From.IsEmpty() {
		return nil, errors.New("wrong from position format")
	}
	move.To = chess.PositionFromString(uci[2:4])
	if move.To.IsEmpty() {
		return nil, errors.New("wrong to position format")
	}

	if len(uci) == 5 {
		move.PromotedPieceNotation = strings.ToUpper(uci[4:5])
		if !slices.Contains(piece.AllNotations, move.PromotedPieceNotation) {
			return nil, errors.New("unknown promoted piece notation")
		}
	}

	return move, nil
}

func (m *Move) String() string {
	return m.From.String() + m.To.String() + strings.ToLower(m.PromotedPieceNotation)
}
