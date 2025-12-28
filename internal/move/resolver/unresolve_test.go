package resolver_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/resolver"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
)

func TestUnresolveFrom(t *testing.T) {
	type args struct {
		move  move.Piece
		board chess.Board
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
				move.NewPiece(chess.PositionFromString("d1"), chess.PositionFromString("d4")),
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
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
				move.NewPiece(chess.PositionFromString("a1"), chess.PositionFromString("d1")),
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
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
				move.NewPiece(chess.PositionFromString("b7"), chess.PositionFromString("d5")),
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
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
				move.NewPiece(chess.PositionFromString("g1"), chess.PositionFromString("e2")),
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
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
				move.NewPiece(chess.PositionFromString("e2"), chess.PositionFromString("e4")),
				standardtest.NewBoard(chess.SideBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("e2"): piece.NewPawn(chess.SideBlack),
					chess.PositionFromString("f2"): piece.NewPawn(chess.SideBlack),
				}),
			},
			chess.NewEmptyPosition(),
			false,
		},
		{
			"single_pawn_capture",
			args{
				move.NewPiece(chess.PositionFromString("b7"), chess.PositionFromString("c8")),
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("b7"): piece.NewPawn(chess.SideWhite),
					chess.PositionFromString("c8"): piece.NewPawn(chess.SideBlack),
				}),
			},
			chess.NewEmptyPosition(),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolver.UnresolveFrom(tt.args.move, tt.args.board)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnresolveFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UnresolveFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
