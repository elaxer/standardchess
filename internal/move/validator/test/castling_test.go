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
	type fields struct {
		board chess.Board
	}
	type args struct {
		castling move.Castling
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"short",
			fields{
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("g8"): standardtest.NewPiece("Q"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
					chess.PositionFromString("b6"): standardtest.NewPiece("r"),
				}),
			},
			args{move.CastlingShort},
			false,
		},
		{
			"long",
			fields{
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
					chess.PositionFromString("g6"): standardtest.NewPiece("r"),
				}),
			},
			args{move.CastlingLong},
			false,
		},
		{
			"king_is_walked",
			fields{
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPieceM("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"rook_is_walked",
			fields{
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPieceM("R"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"let",
			fields{
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("g1"): standardtest.NewPiece("N"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"obstacle",
			fields{
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
					chess.PositionFromString("g1"): standardtest.NewPiece("n"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"future_check",
			fields{
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
					chess.PositionFromString("g8"): standardtest.NewPiece("r"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"attacked_castling_square",
			fields{
				standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): standardtest.NewPiece("K"),
					chess.PositionFromString("a1"): standardtest.NewPiece("R"),
					chess.PositionFromString("h1"): standardtest.NewPiece("R"),
					chess.PositionFromString("f8"): standardtest.NewPiece("r"),
				}),
			},
			args{move.CastlingShort},
			true,
		},
		{
			"another_piece_instead_rook",
			fields{standardtest.DecodeFEN("12/12/12/12/12/3K3P2N1")},
			args{move.CastlingShort},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateCastlingMove(tt.args.castling, tt.fields.board.Turn(), tt.fields.board, true)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCastlingMove() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
