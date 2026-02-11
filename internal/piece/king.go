package piece

import (
	"encoding/json"
	"math"

	"github.com/elaxer/chess"
)

const (
	NotationKing = "K"
	WeightKing   = math.MaxUint16
)

type King struct {
	*abstract

	pseudoMoves []chess.Position
}

func NewKing(color chess.Color) *King {
	return &King{&abstract{color, false}, make([]chess.Position, 0, 8)}
}

func (k *King) PseudoMoves(from chess.Position, squares *chess.Squares) []chess.Position {
	k.pseudoMoves = k.pseudoMoves[:0]

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

	for _, move := range positions {
		if move.Validate() != nil {
			continue
		}

		if piece, err := squares.FindByPosition(move); err == nil && k.canMove(piece, k.color) {
			k.pseudoMoves = append(k.pseudoMoves, move)
		}
	}

	return k.pseudoMoves
}

func (k *King) Notation() string {
	return NotationKing
}

func (k *King) Weight() uint16 {
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
