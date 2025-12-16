package fen

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/chess/encoding/fen"
	"github.com/elaxer/chess/metric"
	standardmetric "github.com/elaxer/standardchess/metric"
	"github.com/elaxer/standardchess/move/move"
)

// NewEncoder creates a new FEN encoder for the standard chess variant.
// Encoder.Encode() method will return a FEN string representing the current state of the board,
// including the turn, castling rights, en passant target square, halfmove clock, and fullmove counter.
//
// See github.com/elaxer/chess/encoding/fen/encoder for more details.
func NewEncoder() *fen.Encoder {
	return &fen.Encoder{
		MetricFuncs: []metric.MetricFunc{
			castlingMetric,
			standardmetric.EnPassantTargetSquare,
			standardmetric.HalfmoveClock,
			metric.FullmoveCounter,
		},
	}
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
