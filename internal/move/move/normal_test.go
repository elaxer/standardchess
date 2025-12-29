package move

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalFromString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"pawn",
			args{"e4"},
			"e4",
			false,
		},
		{
			"rook",
			args{"Rd8"},
			"Rd8",
			false,
		},
		{
			"bishop",
			args{"Ba1"},
			"Ba1",
			false,
		},
		{
			"knight",
			args{"Nc3"},
			"Nc3",
			false,
		},
		{
			"queen",
			args{"Qc6"},
			"Qc6",
			false,
		},
		{
			"king",
			args{"Kb7"},
			"Kb7",
			false,
		},
		{
			"unknown_piece",
			args{"Zk9"},
			"",
			true,
		},
		{
			"wrong_file",
			args{"x3"},
			"",
			true,
		},
		{
			"wrong_rank",
			args{"d21"},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalFromString(tt.args.str)

			require.Truef(t, (err != nil) == tt.wantErr, "NormalFromString() error = %v, wantErr %v", err, tt.wantErr)
			if !tt.wantErr {
				assert.Equal(t, tt.want, got.String())
			}
		})
	}
}

func TestNormal_String(t *testing.T) {
	tests := []struct {
		name string
		move *Normal
		want string
	}{
		{
			"queen",
			NewNormal(chess.NewPositionEmpty(), chess.PositionFromString("a8"), piece.NotationQueen),
			"Qa8",
		},
		{
			"pawn",
			NewNormal(chess.NewPositionEmpty(), chess.PositionFromString("e4"), piece.NotationPawn),
			"e4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.move.String())
		})
	}
}
