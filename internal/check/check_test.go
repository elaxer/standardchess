package check_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/check"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
)

func TestIsCheck(t *testing.T) {
	type args struct {
		board chess.Board
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"check",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("a8"): piece.NewRook(chess.ColorBlack),
				}),
			},
			true,
		},
		{
			"check_bishop",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("h8"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("b4"): piece.NewBishop(chess.ColorWhite),
				}),
			},
			true,
		},
		{
			"check_pawns",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("d4"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("c3"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("e3"): piece.NewBishop(chess.ColorWhite),
				}),
			},
			true,
		},
		{
			"no_check",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("d4"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("a1"): piece.NewRook(chess.ColorBlack),
				}),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := check.IsCheck(tt.args.board)
			assert.Equal(t, tt.want, got)
		})
	}
}
