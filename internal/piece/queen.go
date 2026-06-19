package piece

import (
	"encoding/json"

	"github.com/elaxer/chess"
)

const (
	NotationQueen = "Q"
	WeightQueen   = 9
)

type Queen struct {
	*sliding

	pseudoMoves []chess.Position
}

func NewQueen(color chess.Color) *Queen {
	return &Queen{&sliding{&abstract{color, false}}, make([]chess.Position, 0, 27)}
}

func (q *Queen) Side() chess.Color {
	return q.color
}

func (q *Queen) PseudoMoves(from chess.Position, squares *chess.Squares) []chess.Position {
	q.pseudoMoves = q.pseudoMoves[:0]

	for _, direction := range chess.DirectionsAll {
		for move := range q.slide(from, direction, squares) {
			q.pseudoMoves = append(q.pseudoMoves, move)
		}
	}

	return q.pseudoMoves
}

func (q *Queen) Notation() string {
	return NotationQueen
}

func (q *Queen) Weight() uint16 {
	return WeightQueen
}

func (q *Queen) String() string {
	if q.color == chess.ColorBlack {
		return "q"
	}

	return "Q"
}

func (q *Queen) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"color":    q.color,
		"notation": q.Notation(),
		"is_moved": q.isMoved,
	})
}
