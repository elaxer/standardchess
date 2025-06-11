package standardtest

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/chess/chesstest"
	"github.com/elaxer/chess/position"
	"github.com/elaxer/standardchess/board"
	"github.com/elaxer/standardchess/piece"
)

func NewBoard(turn chess.Side, placement map[position.Position]chess.Piece) chess.Board {
	board, err := board.NewFactory().Create(turn, placement)
	if err != nil {
		panic(err)
	}

	return board
}

func NewBoardFromMoves(moveStrings ...string) chess.Board {
	moves := chesstest.MoveStrings(moveStrings...)

	board, err := board.NewFactory().CreateFromMoves(moves)
	if err != nil {
		panic(err)
	}

	return board
}

// NewPiece creates a new piece by string.
// Created piece marked as not moved.
// P, R, N, B, Q, K - creates white piece
// p, r, n, b, q, k - creates black piece
func NewPiece(str string) chess.Piece {
	piece, err := piece.NewFactory().CreateFromString(str)
	if err != nil {
		panic(err)
	}

	return piece
}

// NewPieceM creates a new moved piece by string.
// See NewPiece for more details
func NewPieceM(str string) chess.Piece {
	piece, err := piece.NewFactory().CreateFromString(str)
	if err != nil {
		panic(err)
	}

	piece.MarkMoved()

	return piece
}
