package standardtest

import (
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/chesstest"
	"github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/encoding/fen"
	"github.com/elaxer/standardchess/internal/piece"
)

func NewBoard(turn chess.Side, placement map[chess.Position]chess.Piece) chess.Board {
	board, err := standardchess.NewBoard(turn, placement)
	if err != nil {
		panic(err)
	}

	return board
}

func NewBoardFromMoves(moveStrings ...string) chess.Board {
	moves := chesstest.MoveStrings(moveStrings...)

	board, err := standardchess.NewBoardFromMoves(moves)
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

func EncodeFEN(b chess.Board) string {
	return fen.Encode(b)
}

func EncodeFenRows(b chess.Board) string {
	return fen.EncodeSquares(b.Squares())
}

// DecodeFEN decodes a FEN string into a chess.Board instance with the specified edge position.
func DecodeFEN(str string) chess.Board {
	board, err := fen.Decode(str)
	if err != nil {
		panic(err)
	}

	return board
}
