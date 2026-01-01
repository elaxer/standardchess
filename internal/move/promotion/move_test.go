package promotion

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPromotionFromString(t *testing.T) {
	type args struct {
		notation string
	}
	tests := []struct {
		name    string
		args    args
		want    *Move
		wantErr bool
	}{
		{
			"promotion",
			args{"e8=Q"},
			NewMove(chess.NewPositionEmpty(), chess.PositionFromString("e8"), piece.NotationQueen),
			false,
		},
		{
			"from_file",
			args{"fe8=R"},
			NewMove(chess.PositionFromString("f"), chess.PositionFromString("e8"), piece.NotationRook),
			false,
		},
		{
			"with_check",
			args{"d1=N+"},
			NewMove(chess.NewPositionEmpty(), chess.PositionFromString("d1"), piece.NotationKnight),
			false,
		},
		{
			"with_checkmate",
			args{"a8=R#"},
			NewMove(chess.NewPositionEmpty(), chess.PositionFromString("a8"), piece.NotationRook),
			false,
		},
		{
			"with_capture",
			args{"xc8=B"},
			NewMove(chess.NewPositionEmpty(), chess.PositionFromString("c8"), piece.NotationBishop),
			false,
		},
		{
			"invalid_piece",
			args{"c1=K"},
			nil,
			true,
		},
		{
			"invalid_file",
			args{"w8=B"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MoveFromString(tt.args.notation)

			require.Truef(t, (err != nil) == tt.wantErr, "PromotionFromString() error = %v, wantErr %v", err, tt.wantErr)
			if !tt.wantErr {
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestPromotion_String(t *testing.T) {
	type fields struct {
		promotion *Move
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"promotion",
			fields{NewMove(chess.NewPositionEmpty(), chess.PositionFromString("a1"), piece.NotationRook)},
			"a1=R",
		},
		{
			"from_file",
			fields{NewMove(chess.PositionFromString("f"), chess.PositionFromString("e8"), piece.NotationRook)},
			"fe8=R",
		},
		{
			"full_from",
			fields{NewMove(chess.PositionFromString("b2"), chess.PositionFromString("b1"), piece.NotationKnight)},
			"b2b1=N",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.promotion.String())
		})
	}
}
