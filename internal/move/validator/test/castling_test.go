package validator_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/validator"
	"github.com/elaxer/standardchess/internal/standardtest"
)

// todo добавить тесты с новым параметром Side
func TestValidateCastlingMove(t *testing.T) {
	type args struct {
		castlingType move.Castling
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
				move.CastlingShort,
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
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
				move.CastlingLong,
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
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
				move.CastlingShort,
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
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
				move.CastlingShort,
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPieceM("R"),
				}),
			},
			true,
		},
		{
			"let",
			args{
				move.CastlingShort,
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
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
				move.CastlingShort,
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
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
				move.CastlingShort,
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
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
				move.CastlingShort,
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
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
			args{move.CastlingShort, standardtest.DecodeFEN("12/12/12/12/12/3K3P2N1")},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateCastlingMove(tt.args.castlingType, tt.args.board.Turn(), tt.args.board, true)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCastlingMove() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
