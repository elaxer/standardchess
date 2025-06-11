package resolver_test

import (
	"testing"

	. "github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
	"github.com/elaxer/standardchess/move/move"
	"github.com/elaxer/standardchess/move/resolver"
	"github.com/elaxer/standardchess/piece"
	"github.com/elaxer/standardchess/standardtest"
)

func TestUnresolveFrom(t *testing.T) {
	type args struct {
		move  move.Piece
		board Board
	}
	tests := []struct {
		name    string
		args    args
		want    position.Position
		wantErr bool
	}{
		{
			"same_file",
			args{
				move.NewPiece(position.FromString("d1"), position.FromString("d4")),
				standardtest.NewBoard(SideWhite, map[position.Position]Piece{
					position.FromString("d1"): piece.NewQueen(SideWhite),
					position.FromString("d8"): piece.NewQueen(SideWhite),
				}),
			},
			position.Position{Rank: position.Rank1},
			false,
		},
		{
			"same_rank",
			args{
				move.NewPiece(position.FromString("a1"), position.FromString("d1")),
				standardtest.NewBoard(SideWhite, map[position.Position]Piece{
					position.FromString("a1"): piece.NewRook(SideBlack),
					position.FromString("g1"): piece.NewRook(SideBlack),
				}),
			},
			position.Position{File: position.FileA},
			false,
		},
		{
			"same_file_and_rank",
			args{
				move.NewPiece(position.FromString("b7"), position.FromString("d5")),
				standardtest.NewBoard(SideWhite, map[position.Position]Piece{
					position.FromString("b7"): piece.NewBishop(SideWhite),
					position.FromString("f7"): piece.NewBishop(SideWhite),
					position.FromString("b3"): piece.NewBishop(SideWhite),
				}),
			},
			position.FromString("b7"),
			false,
		},
		{
			"no_same_file_and_rank",
			args{
				move.NewPiece(position.FromString("g1"), position.FromString("e2")),
				standardtest.NewBoard(SideWhite, map[position.Position]Piece{
					position.FromString("c3"): piece.NewKnight(SideWhite),
					position.FromString("g1"): piece.NewKnight(SideWhite),
				}),
			},
			position.FromString("g"),
			false,
		},
		{
			"no_same_moves",
			args{
				move.NewPiece(position.FromString("e2"), position.FromString("e4")),
				standardtest.NewBoard(SideBlack, map[position.Position]Piece{
					position.FromString("e2"): piece.NewPawn(SideBlack),
					position.FromString("f2"): piece.NewPawn(SideBlack),
				}),
			},
			position.NewEmpty(),
			false,
		},
		{
			"single_pawn_capture",
			args{
				move.NewPiece(position.FromString("b7"), position.FromString("c8")),
				standardtest.NewBoard(SideWhite, map[position.Position]Piece{
					position.FromString("b7"): piece.NewPawn(SideWhite),
					position.FromString("c8"): piece.NewPawn(SideBlack),
				}),
			},
			position.NewEmpty(),
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
