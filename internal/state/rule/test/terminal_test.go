package rule_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/elaxer/standardchess/internal/state/rule"
	"github.com/elaxer/standardchess/internal/state/state"
)

// todo test opposite turn
func TestCheckmate(t *testing.T) {
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
			"checkmate",
			args{
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): piece.NewKing(chess.SideWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.SideBlack),
					chess.PositionFromString("a8"): piece.NewRook(chess.SideBlack),
					chess.PositionFromString("b8"): piece.NewRook(chess.SideBlack),
				}),
				chess.SideWhite,
			},
			state.Checkmate,
		},
		{
			// no checkmate because the black king can capture the threatening rook
			"no_checkmate",
			args{
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): piece.NewKing(chess.SideWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.SideBlack),
					chess.PositionFromString("a2"): piece.NewRook(chess.SideBlack),
					chess.PositionFromString("b8"): piece.NewRook(chess.SideBlack),
				}),
				chess.SideWhite,
			},

			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rule.Checkmate(tt.args.board, tt.args.side); got != tt.want {
				t.Errorf("Checkmate() = %v, want %v", got, tt.want)
			}
		})
	}
}
