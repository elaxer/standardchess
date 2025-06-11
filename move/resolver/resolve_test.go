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

func TestResolveFrom(t *testing.T) {
	type args struct {
		move  *move.Normal
		board Board
	}
	tests := []struct {
		name    string
		args    args
		want    position.Position
		wantErr bool
	}{
		{
			"empty_from",
			args{
				move: move.NewNormal(position.NewEmpty(), position.FromString("e4"), piece.NotationPawn),
				board: standardtest.NewBoard(SideWhite, map[position.Position]Piece{
					position.FromString("d2"): piece.NewPawn(SideWhite),
					position.FromString("e2"): piece.NewPawn(SideWhite),
					position.FromString("f2"): piece.NewPawn(SideWhite),
				}),
			},
			position.FromString("e2"),
			false,
		},
		{
			"same_file",
			args{
				move: move.NewNormal(position.FromString("a"), position.FromString("b8"), piece.NotationRook),
				board: standardtest.NewBoard(SideBlack, map[position.Position]Piece{
					position.FromString("f8"): piece.NewRook(SideBlack),
					position.FromString("a8"): piece.NewRook(SideBlack),
				}),
			},
			position.FromString("a8"),
			false,
		},
		{
			"knights",
			args{
				move: move.NewNormal(position.FromString("g"), position.FromString("e2"), piece.NotationKnight),
				board: standardtest.NewBoard(SideWhite, map[position.Position]Piece{
					position.FromString("g1"): piece.NewKnight(SideWhite),
					position.FromString("c3"): piece.NewKnight(SideWhite),
				}),
			},
			position.FromString("g1"),
			false,
		},
		{
			"same_rank",
			args{
				move: move.NewNormal(position.FromString("1"), position.FromString("a5"), piece.NotationRook),
				board: standardtest.NewBoard(SideWhite, map[position.Position]Piece{
					position.FromString("a1"): piece.NewRook(SideWhite),
					position.FromString("a8"): piece.NewRook(SideWhite),
				}),
			},
			position.FromString("a1"),
			false,
		},
		{
			"full_from",
			args{
				move: move.NewNormal(position.FromString("f2"), position.FromString("d4"), piece.NotationBishop),
				board: standardtest.NewBoard(SideBlack, map[position.Position]Piece{
					position.FromString("b2"): piece.NewBishop(SideBlack),
					position.FromString("f2"): piece.NewBishop(SideBlack),
					position.FromString("b6"): piece.NewBishop(SideBlack),
				}),
			},
			position.FromString("f2"),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolver.ResolveFrom(tt.args.move.Piece, tt.args.move.PieceNotation, tt.args.board, tt.args.board.Turn())
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveNormal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResolveNormal() = %v, want %v", got, tt.want)
			}
		})
	}
}
