package piece

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
)

const (
	NotationBishop = "B"
	WeightBishop   = 3
)

type Bishop struct {
	*sliding
}

func NewBishop(side chess.Side) *Bishop {
	return &Bishop{&sliding{&abstract{side, false}}}
}

func (b *Bishop) PseudoMoves(from position.Position, squares *chess.Squares) position.Set {
	moves := mapset.NewSetWithSize[position.Position](13)
	for _, direction := range diagonalDirections {
		for move := range b.slide(from, direction, squares) {
			moves.Add(move)
		}
	}

	return moves
}

func (b *Bishop) Notation() string {
	return NotationBishop
}

func (b *Bishop) Weight() uint8 {
	return WeightBishop
}

func (b *Bishop) String() string {
	if b.side == chess.SideBlack {
		return "b"
	}

	return "B"
}

func (b *Bishop) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     b.side,
		"notation": b.Notation(),
	})
}
