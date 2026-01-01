package castling

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
			"O character",
			args{"O-O"},
			TypeShort,
			false,
		},
		{
			"All characters",
			args{"O-o-0+"},
			TypeLong,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TypeFromString(tt.args.str)

			require.Truef(t, (err != nil) == tt.wantErr, "CastlingFromString() error = %v, wantErr %v", err, tt.wantErr)
			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCastling_String(t *testing.T) {
	type fields struct {
		move CastlingType
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
			assert.Equal(t, tt.want, tt.fields.move.String())
		})
	}
}
