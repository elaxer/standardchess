package piece

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
)

const (
	NotationRook = "R"
	WeightRook   = 5
)

type Rook struct {
	*sliding
}

func NewRook(side chess.Side) *Rook {
	return &Rook{&sliding{&abstract{side, false}}}
}

func (r *Rook) PseudoMoves(from position.Position, squares *chess.Squares) position.Set {
	moves := mapset.NewSetWithSize[position.Position](14)
	for _, direction := range orthogonalDirections {
		for move := range r.slide(from, direction, squares) {
			moves.Add(move)
		}
	}

	return moves
}

func (r *Rook) Notation() string {
	return NotationRook
}

func (r *Rook) Weight() uint8 {
	return WeightRook
}

func (r *Rook) String() string {
	if r.side == chess.SideBlack {
		return "r"
	}

	return "R"
}

func (r *Rook) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     r.Side(),
		"notation": r.Notation(),
	})
}
