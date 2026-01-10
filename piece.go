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

// NewPiece creates a new chess piece based on the provided notation and side.
// Returns nil if the piece cannot be created.
func NewPiece(notation string, color chess.Color) (chess.Piece, error) {
	return piece.New(notation, color)
}

func NewRook(color chess.Color) chess.Piece {
	return piece.NewRook(color)
}

func NewKnight(color chess.Color) chess.Piece {
	return piece.NewKnight(color)
}

func NewBishop(color chess.Color) chess.Piece {
	return piece.NewBishop(color)
}

func NewQueen(color chess.Color) chess.Piece {
	return piece.NewQueen(color)
}

func NewKing(color chess.Color) chess.Piece {
	return piece.NewKing(color)
}

func NewPawn(color chess.Color) chess.Piece {
	return piece.NewPawn(color)
}
