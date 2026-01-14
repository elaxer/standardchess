package fen

import (
	"iter"
	"strconv"
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/castling"
	"github.com/elaxer/standardchess/internal/move/enpassant"
	"github.com/elaxer/standardchess/metric"
)

// Encode encodes the given chess board into a FEN string.
// If the board is nil, it returns an empty string.
// If MetricFuncs are provided, it appends their results to the FEN string.
// The format of the FEN string is:
// <piece placement> <turn> [<metric1> <metric2> ...].
// If no metrics are provided, it will only include the piece placement and turn.
// If metric functions return nil, it will append a dash ("-") for that metric.
func Encode(board chess.Board) FEN {
	if board == nil {
		return FEN{}
	}

	return FEN{
		placement:       encodeSquares(board.Squares()),
		turn:            board.Turn(),
		castlings:       castlings(board),
		enPassantSquare: enpassant.EnPassantTargetSquare(board),
		halfmoveClock:   metric.HalfmoveClock(board).Value().(int),
		moveNumber:      len(board.MoveHistory())/2 + 1,
	}
}

func encodeSquares(squares *chess.Squares) string {
	var fenSb strings.Builder
	for _, row := range squares.IterOverRows(true) {
		fenSb.WriteString(encodeRow(row) + "/")
	}

	fen := fenSb.String()

	return fen[:len(fen)-1]
}

func encodeRow(row iter.Seq2[chess.File, chess.Piece]) string {
	var rowSb strings.Builder
	emptySquares := 0
	for _, piece := range row {
		if piece == nil {
			emptySquares++

			continue
		}

		if emptySquares > 0 {
			rowSb.WriteString(strconv.Itoa(emptySquares))
			emptySquares = 0
		}

		rowSb.WriteString(piece.String())
	}

	if emptySquares > 0 {
		rowSb.WriteString(strconv.Itoa(emptySquares))
	}

	return rowSb.String()
}

func castlings(board chess.Board) map[chess.Color]map[castling.CastlingType]bool {
	return map[chess.Color]map[castling.CastlingType]bool{
		chess.ColorWhite: {
			castling.TypeShort: castling.ValidateMoveWithObstacle(castling.TypeShort, chess.ColorWhite, board) == nil,
			castling.TypeLong:  castling.ValidateMoveWithObstacle(castling.TypeLong, chess.ColorWhite, board) == nil,
		},
		chess.ColorBlack: {
			castling.TypeShort: castling.ValidateMoveWithObstacle(castling.TypeShort, chess.ColorBlack, board) == nil,
			castling.TypeLong:  castling.ValidateMoveWithObstacle(castling.TypeLong, chess.ColorBlack, board) == nil,
		},
	}
}
