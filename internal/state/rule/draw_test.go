package rule_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/elaxer/standardchess/internal/state/rule"
	"github.com/elaxer/standardchess/internal/state/state"
	"github.com/stretchr/testify/assert"
)

func TestStalemate(t *testing.T) {
	type args struct {
		board chess.Board
		side  chess.Side
	}
	tests := []struct {
		name string
		args args
		want chess.State
	}{
		{
			"stalemate",
			args{
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a8"): piece.NewKing(chess.SideWhite),
					chess.PositionFromString("b6"): piece.NewKing(chess.SideBlack),
					chess.PositionFromString("c7"): piece.NewQueen(chess.SideBlack),
				}),
				chess.SideWhite,
			},
			state.Stalemate,
		},
		{
			"no_stalemate",
			args{
				standardtest.NewBoard(chess.SideBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("a8"): piece.NewKing(chess.SideWhite),
					chess.PositionFromString("c7"): piece.NewQueen(chess.SideWhite),
					chess.PositionFromString("b6"): piece.NewKing(chess.SideBlack),
				}),
				chess.SideBlack,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rule.Stalemate(tt.args.board, tt.args.side)
			assert.Equal(t, tt.want, got)
		})
	}
}
