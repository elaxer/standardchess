// Package metric contains metric functions for a chessboard.
package metric

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/chess/metric"
	"github.com/elaxer/standardchess/internal/move/castling"
	"github.com/elaxer/standardchess/internal/move/enpassant"
	"github.com/elaxer/standardchess/internal/move/normal"
	"github.com/elaxer/standardchess/internal/piece"
)

var AllFuncs = []metric.MetricFunc{
	CastlingAbility,
	EnPassantTargetSquare,
	HalfmoveClock,
}

type Castlings = map[string]map[chess.Color]map[string]bool

func CastlingAbility(board chess.Board) metric.Metric {
	callback := func(side chess.Color, board chess.Board, practical bool) map[string]bool {
		if !practical {
			return map[string]bool{
				castling.TypeShort.String(): castling.ValidateMoveWithObstacle(
					castling.TypeShort,
					side,
					board,
				) == nil,
				castling.TypeLong.String(): castling.ValidateMoveWithObstacle(
					castling.TypeLong,
					side,
					board,
				) == nil,
			}
		}

		if board.Turn() != side {
			return map[string]bool{
				castling.TypeShort.String(): false,
				castling.TypeLong.String():  false,
			}
		}

		return map[string]bool{
			castling.TypeShort.String(): castling.ValidateMove(castling.TypeShort, board) == nil,
			castling.TypeLong.String():  castling.ValidateMove(castling.TypeLong, board) == nil,
		}
	}

	castlings := Castlings{
		"theoretical": {
			chess.ColorWhite: callback(chess.ColorWhite, board, false),
			chess.ColorBlack: callback(chess.ColorBlack, board, false),
		},
		"practical": {
			chess.ColorWhite: callback(chess.ColorWhite, board, true),
			chess.ColorBlack: callback(chess.ColorBlack, board, true),
		},
	}

	return metric.New("Castling ability", castlings)
}

func EnPassantTargetSquare(board chess.Board) metric.Metric {
	targetPosition := enpassant.EnPassantPosition(board)
	if targetPosition.IsEmpty() {
		return nil
	}

	return metric.New("En passant target square", targetPosition)
}

func HalfmoveClock(board chess.Board) metric.Metric {
	clock := 0
	for _, m := range board.MoveHistory() {
		normalMove, ok := m.(*normal.MoveResult)
		if !ok || normalMove.InputMove.PieceNotation == piece.NotationPawn ||
			normalMove.IsCapture() {
			clock = 0

			continue
		}

		clock++
	}

	return metric.New("Halfmove clock", clock)
}
