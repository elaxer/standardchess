package move

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

			require.Truef(t, (err != nil) == tt.wantErr, "CastlingFromString() error = %v, wantErr %v", err, tt.wantErr)
			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
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
			"short",
			fields{(CastlingShort)},
			"O-O",
		},
		{
			"long",
			fields{(CastlingLong)},
			"O-O-O",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.fields.move.String())
		})
	}
}
