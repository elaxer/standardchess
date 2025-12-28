package piece

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess"
)

const (
	NotationKnight = "N"
	WeightKnight   = 3
)

type Knight struct {
	*abstract
}

func NewKnight(side chess.Side) *Knight {
	return &Knight{&abstract{side, false}}
}

func (k *Knight) PseudoMoves(from chess.Position, squares *chess.Squares) chess.PositionSet {
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

	moves := mapset.NewSetWithSize[chess.Position](len(positions))
	for _, move := range positions {
		if piece, err := squares.FindByPosition(move); err == nil && k.canMove(piece, k.side) {
			moves.Add(move)
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
	if k.side == chess.SideBlack {
		return "n"
	}

	return "N"
}

func (k *Knight) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     k.side,
		"notation": k.Notation(),
	})
}
