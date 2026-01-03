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
}

func PawnRankDirection(side chess.Side) chess.Rank {
	if side == chess.SideBlack {
		return -1
	}

	return 1
}

func NewPawn(side chess.Side) *Pawn {
	return &Pawn{&abstract{side, false}}
}

func (p *Pawn) PseudoMoves(from chess.Position, squares *chess.Squares) []chess.Position {
	direction := PawnRankDirection(p.side)
	moves := make([]chess.Position, 0, 4)
	p.appendMovesForward(&moves, from, direction, squares)
	p.appendMovesDiagonal(&moves, from, direction, squares)

	return moves
}

func (p *Pawn) Notation() string {
	return NotationPawn
}

func (p *Pawn) Weight() uint8 {
	return WeightPawn
}

func (p *Pawn) String() string {
	if p.side == chess.SideBlack {
		return "p"
	}

	return "P"
}

func (p *Pawn) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"side":     p.side,
		"notation": p.Notation(),
		"is_moved": p.isMoved,
	})
}

func (p *Pawn) appendMovesForward(moves *[]chess.Position, from chess.Position, rankDir chess.Rank, squares *chess.Squares) {
	positions := [2]chess.Position{
		chess.NewPosition(from.File, from.Rank+rankDir*1),
		chess.NewPosition(from.File, from.Rank+rankDir*2),
	}

	for i, move := range positions {
		piece, err := squares.FindByPosition(move)
		if (err != nil || piece != nil) || (i == 1 && p.isMoved) {
			break
		}

		*moves = append(*moves, move)
	}
}

func (p *Pawn) appendMovesDiagonal(moves *[]chess.Position, from chess.Position, rankDir chess.Rank, squares *chess.Squares) {
	positions := [2]chess.Position{
		chess.NewPosition(from.File+1, from.Rank+rankDir),
		chess.NewPosition(from.File-1, from.Rank+rankDir),
	}
	for _, move := range positions {
		piece, err := squares.FindByPosition(move)
		if err == nil && piece != nil && piece.Side() != p.side {
			*moves = append(*moves, move)
		}
	}
}
