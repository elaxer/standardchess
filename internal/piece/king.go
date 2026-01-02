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

func NewKing(side chess.Side) *King {
	return &King{&abstract{side, false}}
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

		if piece, err := squares.FindByPosition(move); err == nil && k.canMove(piece, k.side) {
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
	if k.side == chess.SideBlack {
		return "k"
	}

	return "K"
}

func (k *King) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     k.side,
		"notation": k.Notation(),
		"is_moved": k.isMoved,
	})
}
