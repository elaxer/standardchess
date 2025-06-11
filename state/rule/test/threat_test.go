package rule_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
	"github.com/elaxer/standardchess/piece"
	"github.com/elaxer/standardchess/standardtest"
	"github.com/elaxer/standardchess/state/rule"
	"github.com/elaxer/standardchess/state/state"
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
				standardtest.NewBoard(chess.SideWhite, map[position.Position]chess.Piece{
					position.FromString("a1"): piece.NewKing(chess.SideWhite),
					position.FromString("h8"): piece.NewKing(chess.SideBlack),
					position.FromString("a8"): piece.NewRook(chess.SideBlack),
				}),
				chess.SideWhite,
			},
			state.Check,
		},
		{
			"check_bishop",
			args{
				standardtest.NewBoard(chess.SideBlack, map[position.Position]chess.Piece{
					position.FromString("e1"): piece.NewKing(chess.SideBlack),
					position.FromString("h8"): piece.NewKing(chess.SideWhite),
					position.FromString("b4"): piece.NewBishop(chess.SideWhite),
				}),
				chess.SideBlack,
			},
			state.Check,
		},
		{
			"check_pawns",
			args{
				standardtest.NewBoard(chess.SideBlack, map[position.Position]chess.Piece{
					position.FromString("d4"): piece.NewKing(chess.SideBlack),
					position.FromString("c3"): piece.NewKing(chess.SideWhite),
					position.FromString("e3"): piece.NewBishop(chess.SideWhite),
				}),
				chess.SideBlack,
			},
			state.Check,
		},
		{
			"no_check",
			args{
				standardtest.NewBoard(chess.SideWhite, map[position.Position]chess.Piece{
					position.FromString("d4"): piece.NewKing(chess.SideWhite),
					position.FromString("h8"): piece.NewKing(chess.SideBlack),
					position.FromString("a1"): piece.NewRook(chess.SideBlack),
				}),
				chess.SideWhite,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rule.Check(tt.args.board, tt.args.side); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
