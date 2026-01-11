package piece

import (
	"encoding/json"
	"math"

	"github.com/elaxer/chess"
)

const (
	NotationKing = "K"
	WeightKing   = math.MaxUint8
)

type King struct {
	*abstract
}

func NewKing(color chess.Color) *King {
	return &King{&abstract{color, false}}
}

func (k *King) PseudoMoves(from chess.Position, squares *chess.Squares) []chess.Position {
	positions := [8]chess.Position{
		chess.NewPosition(from.File, from.Rank+1),
		chess.NewPosition(from.File, from.Rank-1),
		chess.NewPosition(from.File+1, from.Rank),
		chess.NewPosition(from.File-1, from.Rank),
		chess.NewPosition(from.File+1, from.Rank+1),
		chess.NewPosition(from.File-1, from.Rank-1),
		chess.NewPosition(from.File+1, from.Rank-1),
		chess.NewPosition(from.File-1, from.Rank+1),
	}

	moves := make([]chess.Position, 0, len(positions))
	for _, move := range positions {
		if move.Validate() != nil {
			continue
		}

		if piece, err := squares.FindByPosition(move); err == nil && k.canMove(piece, k.color) {
			moves = append(moves, move)
		}
	}

	return moves
}

func (k *King) Notation() string {
	return NotationKing
}

func (k *King) Weight() uint8 {
	return WeightKing
}

func (k *King) String() string {
	if k.color == chess.ColorBlack {
		return "k"
	}

	return "K"
}

func (k *King) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"color":    k.color,
		"notation": k.Notation(),
		"is_moved": k.isMoved,
	})
}
