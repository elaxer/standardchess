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
}

func NewQueen(color chess.Color) *Queen {
	return &Queen{&sliding{&abstract{color, false}}}
}

func (q *Queen) Side() chess.Color {
	return q.color
}

func (q *Queen) PseudoMoves(from chess.Position, squares *chess.Squares) []chess.Position {
	moves := make([]chess.Position, 0, 27)
	for _, direction := range allDirections {
		for move := range q.slide(from, direction, squares) {
			moves = append(moves, move)
		}
	}

	return moves
}

func (q *Queen) Notation() string {
	return NotationQueen
}

func (q *Queen) Weight() uint8 {
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
		"side":     q.color,
		"notation": q.Notation(),
		"is_moved": q.isMoved,
	})
}
