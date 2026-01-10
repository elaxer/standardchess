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

func TestStalemate(t *testing.T) {
	type args struct {
		board chess.Board
	}
	tests := []struct {
		name string
		args args
		want chess.State
	}{
		{
			"stalemate",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a8"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("b6"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("c7"): piece.NewQueen(chess.ColorBlack),
				}),
			},
			state.Stalemate,
		},
		{
			"no_stalemate",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("a8"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("c7"): piece.NewQueen(chess.ColorWhite),
					chess.PositionFromString("b6"): piece.NewKing(chess.ColorBlack),
				}),
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rule.Stalemate(tt.args.board)
			assert.Equal(t, tt.want, got)
		})
	}
}
