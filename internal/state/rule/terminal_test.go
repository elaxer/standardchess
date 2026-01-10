package rule_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/elaxer/standardchess/internal/state"
	"github.com/elaxer/standardchess/internal/state/rule"
	"github.com/stretchr/testify/assert"
)

func TestCheckmate(t *testing.T) {
	type args struct {
		board chess.Board
	}
	tests := []struct {
		name string
		args args
		want chess.State
	}{
		{
			"checkmate",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("a8"): piece.NewRook(chess.ColorBlack),
					chess.PositionFromString("b8"): piece.NewRook(chess.ColorBlack),
				}),
			},
			state.Checkmate,
		},
		{
			// no checkmate because the black king can capture the threatening rook
			"no_checkmate",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("a2"): piece.NewRook(chess.ColorBlack),
					chess.PositionFromString("b8"): piece.NewRook(chess.ColorBlack),
				}),
			},

			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rule.Checkmate(tt.args.board)
			assert.Equal(t, tt.want, got)
		})
	}
}
