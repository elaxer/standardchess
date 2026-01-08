package resolver_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/resolver"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnresolveFrom(t *testing.T) {
	type args struct {
		from, to chess.Position
		board    chess.Board
	}
	tests := []struct {
		name    string
		args    args
		want    chess.Position
		wantErr bool
	}{
		{
			"same_file",
			args{
				chess.PositionFromString("d1"),
				chess.PositionFromString("d4"),
				standardtest.NewBoardEmpty8x8(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("d1"): piece.NewQueen(chess.SideWhite),
					chess.PositionFromString("d8"): piece.NewQueen(chess.SideWhite),
				}),
			},
			chess.Position{Rank: chess.Rank1},
			false,
		},
		{
			"same_rank",
			args{
				chess.PositionFromString("a1"),
				chess.PositionFromString("d1"),
				standardtest.NewBoardEmpty8x8(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): piece.NewRook(chess.SideBlack),
					chess.PositionFromString("g1"): piece.NewRook(chess.SideBlack),
				}),
			},
			chess.Position{File: chess.FileA},
			false,
		},
		{
			"same_file_and_rank",
			args{
				chess.PositionFromString("b7"),
				chess.PositionFromString("d5"),
				standardtest.NewBoardEmpty8x8(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("b7"): piece.NewBishop(chess.SideWhite),
					chess.PositionFromString("f7"): piece.NewBishop(chess.SideWhite),
					chess.PositionFromString("b3"): piece.NewBishop(chess.SideWhite),
				}),
			},
			chess.PositionFromString("b7"),
			false,
		},
		{
			"no_same_file_and_rank",
			args{
				chess.PositionFromString("g1"),
				chess.PositionFromString("e2"),
				standardtest.NewBoardEmpty8x8(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("c3"): piece.NewKnight(chess.SideWhite),
					chess.PositionFromString("g1"): piece.NewKnight(chess.SideWhite),
				}),
			},
			chess.PositionFromString("g"),
			false,
		},
		{
			"no_same_moves",
			args{
				chess.PositionFromString("e2"),
				chess.PositionFromString("e4"),
				standardtest.NewBoardEmpty8x8(chess.SideBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("e2"): piece.NewPawn(chess.SideBlack),
					chess.PositionFromString("f2"): piece.NewPawn(chess.SideBlack),
				}),
			},
			chess.NewPositionEmpty(),
			false,
		},
		{
			"single_pawn_capture",
			args{
				chess.PositionFromString("b7"),
				chess.PositionFromString("c8"),
				standardtest.NewBoardEmpty8x8(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("b7"): piece.NewPawn(chess.SideWhite),
					chess.PositionFromString("c8"): piece.NewPawn(chess.SideBlack),
				}),
			},
			chess.NewPositionEmpty(),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolver.UnresolveFrom(tt.args.from, tt.args.to, tt.args.board)
			require.Truef(t, (err != nil) == tt.wantErr, "UnresolveFrom() error = %v, wantErr %v", err, tt.wantErr)

			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
