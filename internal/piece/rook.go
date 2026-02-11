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

	pseudoMoves []chess.Position
}

func NewRook(color chess.Color) *Rook {
	return &Rook{&sliding{&abstract{color, false}}, make([]chess.Position, 0, 14)}
}

func (r *Rook) PseudoMoves(from chess.Position, squares *chess.Squares) []chess.Position {
	r.pseudoMoves = r.pseudoMoves[:0]
	for _, direction := range orthogonalDirections {
		for move := range r.slide(from, direction, squares) {
			r.pseudoMoves = append(r.pseudoMoves, move)
		}
	}

	return r.pseudoMoves
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
