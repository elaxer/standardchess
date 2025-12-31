package fen

import (
	"fmt"
	"iter"
	"strconv"
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/metric"
	"github.com/elaxer/standardchess/internal/move/move"
	standardmetric "github.com/elaxer/standardchess/metric"
)

var metricFuncs = []metric.MetricFunc{
	castlingMetric,
	standardmetric.EnPassantTargetSquare,
	standardmetric.HalfmoveClock,
	metric.FullmoveCounter,
}

// Encode encodes the given chess board into a FEN string.
// If the board is nil, it returns an empty string.
// If MetricFuncs are provided, it appends their results to the FEN string.
// The format of the FEN string is:
// <piece placement> <turn> [<metric1> <metric2> ...].
// If no metrics are provided, it will only include the piece placement and turn.
// If metric functions return nil, it will append a dash ("-") for that metric.
func Encode(board chess.Board) string {
	if board == nil {
		return ""
	}

	var fen strings.Builder
	fmt.Fprintf(&fen, "%s %s", EncodeSquares(board.Squares()), board.Turn())

	for _, metricFunc := range metricFuncs {
		fmt.Fprintf(&fen, " %v", callMetricFunc(metricFunc, board))
	}

	return fen.String()
}

// EncodeSquares encodes the piece placement of the given squares into a FEN string.
// It iterates through the squares by rows and encodes each row.
// Each row is represented by a string of piece string representation, with empty squares represented by numbers.
func EncodeSquares(squares *chess.Squares) string {
	fen := ""
	var fenSb strings.Builder
	for _, row := range squares.IterOverRows(true) {
		fenSb.WriteString(encodeRow(row) + "/")
	}
	fen += fenSb.String()

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

func callMetricFunc(metricFunc metric.MetricFunc, board chess.Board) any {
	metric := metricFunc(board)
	if metric == nil {
		return "-"
	}

	return metric.Value()
}

func castlingMetric(board chess.Board) metric.Metric {
	theoretical := standardmetric.CastlingAbility(board).Value().(standardmetric.Castlings)["theoretical"]
	str := ""
	if theoretical[chess.SideWhite][move.CastlingShort.String()] {
		str += "K"
	}
	if theoretical[chess.SideWhite][move.CastlingShort.String()] {
		str += "Q"
	}
	if theoretical[chess.SideBlack][move.CastlingShort.String()] {
		str += "k"
	}
	if theoretical[chess.SideBlack][move.CastlingShort.String()] {
		str += "q"
	}

	if str == "" {
		return nil
	}

	return metric.New("Castling Ability", str)
}
