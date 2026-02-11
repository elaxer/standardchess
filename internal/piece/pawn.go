package piece

import (
	"encoding/json"

	"github.com/elaxer/chess"
)

const (
	NotationPawn = ""
	WeightPawn   = 1
)

type Pawn struct {
	*abstract

	pseudoMoves []chess.Position
}

func PawnRankDirection(color chess.Color) chess.Rank {
	if color == chess.ColorBlack {
		return -1
	}

	return 1
}

func NewPawn(color chess.Color) *Pawn {
	return &Pawn{&abstract{color, false}, make([]chess.Position, 0, 4)}
}

func (p *Pawn) PseudoMoves(from chess.Position, squares *chess.Squares) []chess.Position {
	p.pseudoMoves = p.pseudoMoves[:0]

	direction := PawnRankDirection(p.color)

	p.appendMovesForward(from, direction, squares)
	p.appendMovesDiagonal(from, direction, squares)

	return p.pseudoMoves
}

func (p *Pawn) Notation() string {
	return NotationPawn
}

func (p *Pawn) Weight() uint16 {
	return WeightPawn
}

func (p *Pawn) String() string {
	if p.color == chess.ColorBlack {
		return "p"
	}

	return "P"
}

func (p *Pawn) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"color":    p.color,
		"notation": p.Notation(),
		"is_moved": p.isMoved,
	})
}

func (p *Pawn) appendMovesForward(
	from chess.Position,
	rankDir chess.Rank,
	squares *chess.Squares,
) {
	positions := [2]chess.Position{
		chess.NewPosition(from.File, from.Rank+rankDir*1),
		chess.NewPosition(from.File, from.Rank+rankDir*2),
	}

	for i, move := range positions {
		piece, err := squares.FindByPosition(move)
		if (err != nil || piece != nil) || (i == 1 && p.isMoved) {
			break
		}

		p.pseudoMoves = append(p.pseudoMoves, move)
	}
}

func (p *Pawn) appendMovesDiagonal(
	from chess.Position,
	rankDir chess.Rank,
	squares *chess.Squares,
) {
	positions := [2]chess.Position{
		chess.NewPosition(from.File+1, from.Rank+rankDir),
		chess.NewPosition(from.File-1, from.Rank+rankDir),
	}
	for _, move := range positions {
		piece, err := squares.FindByPosition(move)
		if err == nil && piece != nil && piece.Color() != p.color {
			p.pseudoMoves = append(p.pseudoMoves, move)
		}
	}
}
