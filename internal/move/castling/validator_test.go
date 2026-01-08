package castling_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/castling"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
)

// todo добавить тесты с новым параметром Side
func TestValidateCastlingMove(t *testing.T) {
	type args struct {
		castlingType castling.CastlingType
		board        chess.Board
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"short",
			args{
				castling.TypeShort,
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("g8"): standardtest.NewPiece("Q"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
					chess.PositionFromString("b6"): standardtest.NewPiece("r"),
				}),
			},
			false,
		},
		{
			"long",
			args{
				castling.TypeLong,
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
					chess.PositionFromString("g6"): standardtest.NewPiece("r"),
				}),
			},
			false,
		},
		{
			"king_is_walked",
			args{
				castling.TypeShort,
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPieceM("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
				}),
			},
			true,
		},
		{
			"rook_is_walked",
			args{
				castling.TypeShort,
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPieceM("R"),
				}),
			},
			true,
		},
		{
			"opposite_side_rook",
			args{
				castling.TypeShort,
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("r"),
				}),
			},
			true,
		},
		{
			"let",
			args{
				castling.TypeShort,
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("g1"): standardtest.NewPiece("N"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
				}),
			},
			true,
		},
		{
			"obstacle",
			args{
				castling.TypeShort,
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
					chess.PositionFromString("g1"): standardtest.NewPiece("n"),
				}),
			},
			true,
		},
		{
			"future_check",
			args{
				castling.TypeShort,
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
					chess.PositionFromString("g8"): standardtest.NewPiece("r"),
				}),
			},
			true,
		},
		{
			"attacked_castling_square",
			args{
				castling.TypeShort,
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
					chess.PositionFromString("f8"): standardtest.NewPiece("r"),
				}),
			},
			true,
		},
		{
			"another_piece_instead_rook",
			args{castling.TypeShort, standardtest.DecodeFEN("12/12/12/12/12/3K3P2N1")},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := castling.ValidateMove(tt.args.castlingType, tt.args.board.Turn(), tt.args.board, true)
			assert.True(t, (err != nil) == tt.wantErr)
		})
	}
}
