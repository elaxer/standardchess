package piece

import (
	"encoding/json"

	"github.com/elaxer/chess"
)

const (
	NotationRook = "R"
	WeightRook   = 5
)

type Rook struct {
	*sliding
}

func NewRook(color chess.Color) *Rook {
	return &Rook{&sliding{&abstract{color, false}}}
}

func (r *Rook) PseudoMoves(from chess.Position, squares *chess.Squares) []chess.Position {
	moves := make([]chess.Position, 0, 14)
	for _, direction := range orthogonalDirections {
		for move := range r.slide(from, direction, squares) {
			moves = append(moves, move)
		}
	}

	return moves
}

func (r *Rook) Notation() string {
	return NotationRook
}

func (r *Rook) Weight() uint16 {
	return WeightRook
}

func (r *Rook) String() string {
	if r.color == chess.ColorBlack {
		return "r"
	}

	return "R"
}

func (r *Rook) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"color":    r.Color(),
		"notation": r.Notation(),
		"is_moved": r.isMoved,
	})
}
