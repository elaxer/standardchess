package move

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
)

func TestPromotionFromString(t *testing.T) {
	type args struct {
		notation string
	}
	tests := []struct {
		name    string
		args    args
		want    *Promotion
		wantErr bool
	}{
		{
			"promotion",
			args{"e8=Q"},
			NewPromotion(chess.NewPositionEmpty(), chess.PositionFromString("e8"), piece.NotationQueen),
			false,
		},
		{
			"from_file",
			args{"fe8=R"},
			NewPromotion(chess.PositionFromString("f"), chess.PositionFromString("e8"), piece.NotationRook),
			false,
		},
		{
			"with_check",
			args{"d1=N+"},
			NewPromotion(chess.NewPositionEmpty(), chess.PositionFromString("d1"), piece.NotationKnight),
			false,
		},
		{
			"with_checkmate",
			args{"a8=R#"},
			NewPromotion(chess.NewPositionEmpty(), chess.PositionFromString("a8"), piece.NotationRook),
			false,
		},
		{
			"with_capture",
			args{"xc8=B"},
			NewPromotion(chess.NewPositionEmpty(), chess.PositionFromString("c8"), piece.NotationBishop),
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
			got, err := PromotionFromString(tt.args.notation)
			if (err != nil) != tt.wantErr {
				t.Errorf("PromotionFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && *got != *tt.want {
				t.Errorf("PromotionFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromotion_String(t *testing.T) {
	type fields struct {
		promotion *Promotion
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"promotion",
			fields{NewPromotion(chess.NewPositionEmpty(), chess.PositionFromString("a1"), piece.NotationRook)},
			"a1=R",
		},
		{
			"from_file",
			fields{NewPromotion(chess.PositionFromString("f"), chess.PositionFromString("e8"), piece.NotationRook)},
			"fe8=R",
		},
		{
			"full_from",
			fields{NewPromotion(chess.PositionFromString("b2"), chess.PositionFromString("b1"), piece.NotationKnight)},
			"b2b1=N",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.promotion.String(); got != tt.want {
				t.Errorf("Promotion.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
