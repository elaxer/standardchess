package castling

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/chesstest"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var edgePosition = chess.PositionFromString("h8")

func TestFromNotation(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    CastlingType
		wantErr bool
	}{
		{
			"short",
			args{"0-0"},
			TypeShort,
			false,
		},
		{
			"long",
			args{"0-0-0"},
			TypeLong,
			false,
		},
		{
			"short_with_check",
			args{"0-0+"},
			TypeShort,
			false,
		},
		{
			"short_with_checkmate",
			args{"0-0#"},
			TypeShort,
			false,
		},
		{
			"long_with_check",
			args{"0-0-0+"},
			TypeLong,
			false,
		},
		{
			"long_with_checkmate",
			args{"0-0-0#"},
			TypeLong,
			false,
		},
		{
			"O_character",
			args{"O-O"},
			TypeShort,
			false,
		},
		{
			"all_characters",
			args{"O-o-0+"},
			TypeLong,
			false,
		},

		{
			"too_short",
			args{"0"},
			false,
			true,
		},
		{
			"too_long",
			args{"0-0-0-0"},
			false,
			true,
		},
		{
			"non_ascii",
			args{"о-О"},
			false,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromNotation(tt.args.str)

			require.Truef(
				t,
				(err != nil) == tt.wantErr,
				"FromNotation() error = %v, wantErr %v",
				err,
				tt.wantErr,
			)
			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCastlingType_String(t *testing.T) {
	type fields struct {
		castlingType CastlingType
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"short",
			fields{(TypeShort)},
			"O-O",
		},
		{
			"long",
			fields{(TypeLong)},
			"O-O-O",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.castlingType.String())
		})
	}
}

func TestFromUCI(t *testing.T) {
	tests := []struct {
		uciStr  string
		squares *chess.Squares
		turn    chess.Color
		want    CastlingType
		wantErr bool
	}{
		{
			"e1g1",
			chesstest.MustSquaresFromPlacement(edgePosition, map[chess.Position]chess.Piece{
				chess.PositionFromString("e1"): piece.NewKing(chess.ColorWhite),
			}),
			chess.ColorWhite,
			TypeShort,
			false,
		},
		{
			"e1c1",
			chesstest.MustSquaresFromPlacement(edgePosition, map[chess.Position]chess.Piece{
				chess.PositionFromString("e1"): piece.NewKing(chess.ColorWhite),
			}),
			chess.ColorWhite,
			TypeLong,
			false,
		},
		{
			"e8g8",
			chesstest.MustSquaresFromPlacement(edgePosition, map[chess.Position]chess.Piece{
				chess.PositionFromString("e8"): piece.NewKing(chess.ColorBlack),
			}),
			chess.ColorBlack,
			TypeShort,
			false,
		},
		{
			"e8c8",
			chesstest.MustSquaresFromPlacement(edgePosition, map[chess.Position]chess.Piece{
				chess.PositionFromString("e8"): piece.NewKing(chess.ColorBlack),
			}),
			chess.ColorBlack,
			TypeLong,
			false,
		},

		{
			"e1g1",
			chesstest.MustSquaresFromPlacement(edgePosition, map[chess.Position]chess.Piece{
				chess.PositionFromString("e1"): piece.NewKing(chess.ColorBlack),
			}),
			chess.ColorWhite,
			false,
			true,
		},
		{
			"a3g2",
			chesstest.MustSquaresFromPlacement(edgePosition, map[chess.Position]chess.Piece{
				chess.PositionFromString("a3"): piece.NewKing(chess.ColorWhite),
			}),
			chess.ColorWhite,
			false,
			true,
		},
		{
			"e1c1",
			chesstest.MustSquaresFromPlacement(edgePosition, map[chess.Position]chess.Piece{
				chess.PositionFromString("e1"): piece.NewBishop(chess.ColorWhite),
			}),
			chess.ColorWhite,
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.uciStr, func(t *testing.T) {
			got, err := FromUCI(tt.uciStr, tt.squares, tt.turn)
			if err != nil && !tt.wantErr {
				t.Fatalf("wantErr = false, got err = %v", err)
			}
			if err == nil && tt.wantErr {
				t.Fatal("expected error, got nil")
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
