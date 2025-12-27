package standardtest

import (
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/chesstest"
	"github.com/elaxer/chess/position"
	"github.com/elaxer/standardchess/board"
	"github.com/elaxer/standardchess/encoding/fen"
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
	if len(str) != 1 || !strings.Contains("PRNBQKprnbqk", str) {
		panic("invalid piece string")
	}

	side := chess.SideWhite
	if strings.ToLower(str) == str {
		side = chess.SideBlack
	}

	notation := strings.ToUpper(str)
	if notation == "P" {
		notation = piece.NotationPawn
	}

	piece, err := piece.New(notation, side)
	if err != nil {
		panic(err)
	}

	return piece
}

// NewPieceM creates a new moved piece by string.
// See NewPiece for more details
func NewPieceM(str string) chess.Piece {
	piece := NewPiece(str)
	piece.MarkMoved()

	return piece
}

// DecodeFEN8x8 decodes a FEN string into a chess.Board instance with an 8x8 edge position.
func DecodeFEN8x8(str string) chess.Board {
	return DecodeFEN(str, position.FromString("h8"))
}

// DecodeFEN decodes a FEN string into a chess.Board instance with the specified edge position.
func DecodeFEN(str string, edgePosition position.Position) chess.Board {
	decoder := fen.NewDecoder(&chesstest.FactoryMock{EdgePosition: edgePosition})
	board, err := decoder.Decode(str)
	if err != nil {
		panic(err)
	}

	return board
}
