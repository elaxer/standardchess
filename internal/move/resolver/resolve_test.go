package resolver_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/internal/move/resolver"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveFrom(t *testing.T) {
	type args struct {
		from, to      chess.Position
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
				to:            chess.PositionFromString("e4"),
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
				from:          chess.PositionFromString("a"),
				to:            chess.PositionFromString("b8"),
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
				from:          chess.PositionFromString("g"),
				to:            chess.PositionFromString("e2"),
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
				from:          chess.PositionFromString("1"),
				to:            chess.PositionFromString("a5"),
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
				from:          chess.PositionFromString("f2"),
				to:            chess.PositionFromString("d4"),
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
			got, err := resolver.ResolveFrom(tt.args.from, tt.args.to, tt.args.pieceNotation, tt.args.board, tt.args.board.Turn())
			require.Truef(t, (err != nil) == tt.wantErr, "ResolveNormal() error = %v, wantErr %v", err, tt.wantErr)

			if tt.wantErr {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func BenchmarkResolve(b *testing.B) {
	board := standardchess.NewBoardFilled()

	from, to := chess.NewPositionEmpty(), chess.PositionFromString("e4")
	turn := board.Turn()

	b.ResetTimer()
	for range b.N {
		_, _ = resolver.ResolveFrom(from, to, piece.NotationPawn, board, turn)
	}
}
