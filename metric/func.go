package metric

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/chess/metric"
	"github.com/elaxer/chess/position"
	"github.com/elaxer/standardchess/move/move"
	"github.com/elaxer/standardchess/move/result"
	"github.com/elaxer/standardchess/move/validator"
	"github.com/elaxer/standardchess/piece"
)

type Castlings = map[string]map[chess.Side]map[string]bool

var AllFuncs = []metric.MetricFunc{
	CastlingAbility,
	EnPassantTargetSquare,
	HalfmoveClock,
}

func CastlingAbility(board chess.Board) metric.Metric {
	callback := func(side chess.Side, board chess.Board, validateObstacle bool) map[string]bool {
		return map[string]bool{
			move.CastlingShort.String(): validator.ValidateCastlingMove(move.CastlingShort, side, board, validateObstacle) == nil,
			move.CastlingLong.String():  validator.ValidateCastlingMove(move.CastlingLong, side, board, validateObstacle) == nil,
		}
	}

	castlings := Castlings{
		"theoretical": {
			chess.SideWhite: callback(chess.SideWhite, board, false),
			chess.SideBlack: callback(chess.SideBlack, board, false),
		},
		"practical": {
			chess.SideWhite: callback(chess.SideWhite, board, true),
			chess.SideBlack: callback(chess.SideBlack, board, true),
		},
	}

	return metric.New("Castling ability", castlings)
}

func EnPassantTargetSquare(board chess.Board) metric.Metric {
	if len(board.MovesHistory()) == 0 {
		return nil
	}

	lastMove := board.MovesHistory()[len(board.MovesHistory())-1]
	normalMove, ok := lastMove.(*result.Normal)
	if !ok || normalMove.InputMove.PieceNotation != piece.NotationPawn {
		return nil
	}

	if normalMove.InputMove.To.Rank != normalMove.FromFull.Rank+(piece.PawnRankDirection(!board.Turn())*2) {
		return nil
	}

	passant := position.New(
		normalMove.InputMove.To.File,
		normalMove.FromFull.Rank+piece.PawnRankDirection(!board.Turn()),
	)

	return metric.New("En passant target square", passant)
}

func HalfmoveClock(board chess.Board) metric.Metric {
	clock := 0
	for _, m := range board.MovesHistory() {
		normalMove, ok := m.(*result.Normal)
		if !ok || normalMove.InputMove.PieceNotation == piece.NotationPawn || normalMove.IsCapture() {
			clock = 0

			continue
		}

		clock++
	}

	return metric.New("Halfmove clock", clock)
}
