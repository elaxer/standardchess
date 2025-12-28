package standardchess

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
)

const (
	NotationRook   = piece.NotationRook
	NotationKnight = piece.NotationKnight
	NotationBishop = piece.NotationBishop
	NotationQueen  = piece.NotationQueen
	NotationKing   = piece.NotationKing
	NotationPawn   = piece.NotationPawn

	WeightRook   = piece.WeightRook
	WeightKnight = piece.WeightKnight
	WeightBishop = piece.WeightBishop
	WeightQueen  = piece.WeightQueen
	WeightKing   = piece.WeightKing
	WeightPawn   = piece.WeightPawn
)

func NewPiece(notation string, side chess.Side) (chess.Piece, error) {
	return piece.New(notation, side)
}

func NewRook(side chess.Side) chess.Piece {
	return piece.NewRook(side)
}

func NewKnight(side chess.Side) chess.Piece {
	return piece.NewKnight(side)
}

func NewBishop(side chess.Side) chess.Piece {
	return piece.NewBishop(side)
}

func NewQueen(side chess.Side) chess.Piece {
	return piece.NewQueen(side)
}

func NewKing(side chess.Side) chess.Piece {
	return piece.NewKing(side)
}

func NewPawn(side chess.Side) chess.Piece {
	return piece.NewPawn(side)
}
