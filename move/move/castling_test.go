package move

import (
	"testing"
)

func TestCastlingFromString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    Castling
		wantErr bool
	}{
		{
			"short",
			args{"0-0"},
			CastlingShort,
			false,
		},
		{
			"long",
			args{"0-0-0"},
			CastlingLong,
			false,
		},
		{
			"short_with_check",
			args{"0-0+"},
			CastlingShort,
			false,
		},
		{
			"short_with_checkmate",
			args{"0-0#"},
			CastlingShort,
			false,
		},
		{
			"long_with_check",
			args{"0-0-0+"},
			CastlingLong,
			false,
		},
		{
			"long_with_checkmate",
			args{"0-0-0#"},
			CastlingLong,
			false,
		},
		{
			"O character",
			args{"O-O"},
			CastlingShort,
			false,
		},
		{
			"All characters",
			args{"O-o-0+"},
			CastlingLong,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CastlingFromString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("CastlingFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("CastlingFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCastling_String(t *testing.T) {
	type fields struct {
		move Castling
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"castling_short",
			fields{(CastlingShort)},
			"O-O",
		},
		{
			"castling_long",
			fields{(CastlingLong)},
			"O-O-O",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.move.String(); got != tt.want {
				t.Errorf("Castling.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
