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

	pseudoMoves []chess.Position
}

func NewKnight(color chess.Color) *Knight {
	return &Knight{&abstract{color, false}, make([]chess.Position, 0, 4)}
}

func (k *Knight) PseudoMoves(from chess.Position, squares *chess.Squares) []chess.Position {
	k.pseudoMoves = k.pseudoMoves[:0]

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

	for _, move := range positions {
		if piece, err := squares.FindByPosition(move); err == nil && k.canMove(piece, k.color) {
			k.pseudoMoves = append(k.pseudoMoves, move)
		}
	}

	return k.pseudoMoves
}

func (k *Knight) Notation() string {
	return NotationKnight
}

func (k *Knight) Weight() uint16 {
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
		"color":    k.color,
		"notation": k.Notation(),
		"is_moved": k.isMoved,
	})
}
