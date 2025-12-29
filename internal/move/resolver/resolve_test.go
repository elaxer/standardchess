package resolver_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/resolver"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveFrom(t *testing.T) {
	type args struct {
		pieceMove     move.Piece
		pieceNotation string
		board         chess.Board
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
				pieceMove:     move.Piece{To: chess.PositionFromString("e4")},
				pieceNotation: piece.NotationPawn,
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
				pieceMove:     move.Piece{From: chess.PositionFromString("a"), To: chess.PositionFromString("b8")},
				pieceNotation: piece.NotationRook,
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
				pieceMove:     move.Piece{From: chess.PositionFromString("g"), To: chess.PositionFromString("e2")},
				pieceNotation: piece.NotationKnight,
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
				pieceMove:     move.Piece{From: chess.PositionFromString("1"), To: chess.PositionFromString("a5")},
				pieceNotation: piece.NotationRook,
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
				pieceMove:     move.Piece{From: chess.PositionFromString("f2"), To: chess.PositionFromString("d4")},
				pieceNotation: piece.NotationBishop,
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
			got, err := resolver.ResolveFrom(tt.args.pieceMove, tt.args.pieceNotation, tt.args.board, tt.args.board.Turn())
			require.Truef(t, (err != nil) == tt.wantErr, "ResolveNormal() error = %v, wantErr %v", err, tt.wantErr)

			if tt.wantErr {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
