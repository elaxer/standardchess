package uci_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/uci"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMoveFromString(t *testing.T) {
	tests := []struct {
		uci     string
		want    *uci.Move
		wantErr bool
	}{
		{
			"e2e4",
			&uci.Move{
				From: chess.PositionFromString("e2"),
				To:   chess.PositionFromString("e4"),
			},
			false,
		},
		{
			"a2b1q",
			&uci.Move{
				From:                  chess.PositionFromString("a2"),
				To:                    chess.PositionFromString("b1"),
				PromotedPieceNotation: "Q",
			},
			false,
		},
		{
			"C7C8R",
			&uci.Move{
				From:                  chess.PositionFromString("c7"),
				To:                    chess.PositionFromString("c8"),
				PromotedPieceNotation: "R",
			},
			false,
		},
		{
			"b1b88",
			nil,
			true,
		},
		{
			"ab1b8",
			nil,
			true,
		},
		{
			"b1b8p",
			nil,
			true,
		},
		{
			"b1b8bb",
			nil,
			true,
		},
		{
			"x0x1",
			nil,
			true,
		},
		{
			"1234",
			nil,
			true,
		},
		{
			"1234b",
			nil,
			true,
		},
		{
			"bbbb",
			nil,
			true,
		},
		{
			"bbbbb",
			nil,
			true,
		},

		{
			"b2",
			nil,
			true,
		},
		{
			"cxd5",
			nil,
			true,
		},
		{
			"Nh7#",
			nil,
			true,
		},
		{
			"cxd8=Q+",
			nil,
			true,
		},
		{
			"O-O-O",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.uci, func(t *testing.T) {
			got, err := uci.MoveFromString(tt.uci)
			if err != nil && !tt.wantErr {
				t.Fatal(err)
			}
			if err == nil && tt.wantErr {
				t.Fatal("expected error, got nil")
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestMove_String(t *testing.T) {
	tests := []struct {
		uci  string
		want string
	}{
		{
			"a1b2",
			"a1b2",
		},
		{
			"b3h8q",
			"b3h8q",
		},
		{
			"H2H1R",
			"h2h1r",
		},
	}
	for _, tt := range tests {
		t.Run(tt.uci, func(t *testing.T) {
			m, err := uci.MoveFromString(tt.uci)
			require.NoError(t, err)
			assert.Equal(t, tt.want, m.String())
		})
	}
}
