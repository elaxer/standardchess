package piece

import (
	"encoding/json"

	"github.com/elaxer/chess"
)

const (
	NotationBishop = "B"
	WeightBishop   = 3
)

type Bishop struct {
	*sliding

	pseudoMoves []chess.Position
}

func NewBishop(color chess.Color) *Bishop {
	return &Bishop{&sliding{&abstract{color, false}}, make([]chess.Position, 0, 13)}
}

func (b *Bishop) PseudoMoves(from chess.Position, squares *chess.Squares) []chess.Position {
	b.pseudoMoves = b.pseudoMoves[:0]

	for _, direction := range diagonalDirections {
		for move := range b.slide(from, direction, squares) {
			b.pseudoMoves = append(b.pseudoMoves, move)
		}
	}

	return b.pseudoMoves
}

func (b *Bishop) Notation() string {
	return NotationBishop
}

func (b *Bishop) Weight() uint16 {
	return WeightBishop
}

func (b *Bishop) String() string {
	if b.color == chess.ColorBlack {
		return "b"
	}

	return "B"
}

func (b *Bishop) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"color":    b.color,
		"notation": b.Notation(),
		"is_moved": b.isMoved,
	})
}
