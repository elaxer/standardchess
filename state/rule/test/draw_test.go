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
				standardtest.NewBoard(chess.SideWhite, map[position.Position]chess.Piece{
					position.FromString("a8"): piece.NewKing(chess.SideWhite),
					position.FromString("b6"): piece.NewKing(chess.SideBlack),
					position.FromString("c7"): piece.NewQueen(chess.SideBlack),
				}),
				chess.SideWhite,
			},
			state.Stalemate,
		},
		{
			"no_stalemate",
			args{
				standardtest.NewBoard(chess.SideBlack, map[position.Position]chess.Piece{
					position.FromString("a8"): piece.NewKing(chess.SideWhite),
					position.FromString("c7"): piece.NewQueen(chess.SideWhite),
					position.FromString("b6"): piece.NewKing(chess.SideBlack),
				}),
				chess.SideBlack,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rule.Stalemate(tt.args.board, tt.args.side); got != tt.want {
				t.Errorf("Stalemate() = %v, want %v", got, tt.want)
			}
		})
	}
}
