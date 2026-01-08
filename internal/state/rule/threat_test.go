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

func TestCheck(t *testing.T) {
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
			"check",
			args{
				standardtest.NewBoardEmpty8x8(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): piece.NewKing(chess.SideWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.SideBlack),
					chess.PositionFromString("a8"): piece.NewRook(chess.SideBlack),
				}),
				chess.SideWhite,
			},
			state.Check,
		},
		{
			"check_bishop",
			args{
				standardtest.NewBoardEmpty8x8(chess.SideBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): piece.NewKing(chess.SideBlack),
					chess.PositionFromString("h8"): piece.NewKing(chess.SideWhite),
					chess.PositionFromString("b4"): piece.NewBishop(chess.SideWhite),
				}),
				chess.SideBlack,
			},
			state.Check,
		},
		{
			"check_pawns",
			args{
				standardtest.NewBoardEmpty8x8(chess.SideBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("d4"): piece.NewKing(chess.SideBlack),
					chess.PositionFromString("c3"): piece.NewKing(chess.SideWhite),
					chess.PositionFromString("e3"): piece.NewBishop(chess.SideWhite),
				}),
				chess.SideBlack,
			},
			state.Check,
		},
		{
			"no_check",
			args{
				standardtest.NewBoardEmpty8x8(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("d4"): piece.NewKing(chess.SideWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.SideBlack),
					chess.PositionFromString("a1"): piece.NewRook(chess.SideBlack),
				}),
				chess.SideWhite,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rule.Check(tt.args.board, tt.args.side)
			assert.Equal(t, tt.want, got)
		})
	}
}
