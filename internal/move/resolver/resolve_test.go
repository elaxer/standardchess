package resolver_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/resolver"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
)

func TestResolveFrom(t *testing.T) {
	type args struct {
		move  *move.Normal
		board chess.Board
	}
	tests := []struct {
		name    string
		args    args
		want    chess.Position
		wantErr bool
	}{
		{
			"empty_from",
			args{
				move: move.NewNormal(chess.NewPositionEmpty(), chess.PositionFromString("e4"), piece.NotationPawn),
				board: standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("d2"): piece.NewPawn(chess.SideWhite),
					chess.PositionFromString("e2"): piece.NewPawn(chess.SideWhite),
					chess.PositionFromString("f2"): piece.NewPawn(chess.SideWhite),
				}),
			},
			chess.PositionFromString("e2"),
			false,
		},
		{
			"same_file",
			args{
				move: move.NewNormal(chess.PositionFromString("a"), chess.PositionFromString("b8"), piece.NotationRook),
				board: standardtest.NewBoard(chess.SideBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("f8"): piece.NewRook(chess.SideBlack),
					chess.PositionFromString("a8"): piece.NewRook(chess.SideBlack),
				}),
			},
			chess.PositionFromString("a8"),
			false,
		},
		{
			"knights",
			args{
				move: move.NewNormal(chess.PositionFromString("g"), chess.PositionFromString("e2"), piece.NotationKnight),
				board: standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("g1"): piece.NewKnight(chess.SideWhite),
					chess.PositionFromString("c3"): piece.NewKnight(chess.SideWhite),
				}),
			},
			chess.PositionFromString("g1"),
			false,
		},
		{
			"same_rank",
			args{
				move: move.NewNormal(chess.PositionFromString("1"), chess.PositionFromString("a5"), piece.NotationRook),
				board: standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): piece.NewRook(chess.SideWhite),
					chess.PositionFromString("a8"): piece.NewRook(chess.SideWhite),
				}),
			},
			chess.PositionFromString("a1"),
			false,
		},
		{
			"full_from",
			args{
				move: move.NewNormal(chess.PositionFromString("f2"), chess.PositionFromString("d4"), piece.NotationBishop),
				board: standardtest.NewBoard(chess.SideBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("b2"): piece.NewBishop(chess.SideBlack),
					chess.PositionFromString("f2"): piece.NewBishop(chess.SideBlack),
					chess.PositionFromString("b6"): piece.NewBishop(chess.SideBlack),
				}),
			},
			chess.PositionFromString("f2"),
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
