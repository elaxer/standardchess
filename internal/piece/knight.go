package piece

import (
	"encoding/json"

	"github.com/elaxer/chess"
)

const (
	NotationKnight = "N"
	WeightKnight   = 3
)

type Knight struct {
	*abstract
}

func NewKnight(color chess.Color) *Knight {
	return &Knight{&abstract{color, false}}
}

func (k *Knight) PseudoMoves(from chess.Position, squares *chess.Squares) []chess.Position {
	positions := [8]chess.Position{
		chess.NewPosition(from.File+1, from.Rank+2),
		chess.NewPosition(from.File-1, from.Rank+2),
		chess.NewPosition(from.File+2, from.Rank+1),
		chess.NewPosition(from.File-2, from.Rank+1),
		chess.NewPosition(from.File-1, from.Rank-2),
		chess.NewPosition(from.File-2, from.Rank-1),
		chess.NewPosition(from.File+2, from.Rank-1),
		chess.NewPosition(from.File+1, from.Rank-2),
	}

	moves := make([]chess.Position, 0, len(positions))
	for _, move := range positions {
		if piece, err := squares.FindByPosition(move); err == nil && k.canMove(piece, k.color) {
			moves = append(moves, move)
		}
	}

	return moves
}

func (k *Knight) Notation() string {
	return NotationKnight
}

func (k *Knight) Weight() uint8 {
	return WeightKnight
}

func (k *Knight) String() string {
	if k.color == chess.ColorBlack {
		return "n"
	}

	return "N"
}

func (k *Knight) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     k.color,
		"notation": k.Notation(),
		"is_moved": k.isMoved,
	})
}
