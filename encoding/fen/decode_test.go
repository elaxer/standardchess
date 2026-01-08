package fen_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/encoding/fen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	type args struct {
		fen string
	}
	tests := []struct {
		name      string
		args      args
		wantFEN   string
		wantColor chess.Color
		wantErr   bool
	}{
		{
			"empty",
			args{"8/8/8/8/8/8/8/8"},
			"8/8/8/8/8/8/8/8",
			chess.ColorWhite,
			false,
		},
		{
			"8x8_init_position",
			args{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"},
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			chess.ColorWhite,
			false,
		},
		{
			"white",
			args{"1r1q1n2/2P3b1/3k4/8/2N2B2/1p2P3/2r1n3/3Q1R2 w"},
			"1r1q1n2/2P3b1/3k4/8/2N2B2/1p2P3/2r1n3/3Q1R2",
			chess.ColorWhite,
			false,
		},
		{
			"black",
			args{"3b2n1/pp2p1p1/2k2r2/3P4/8/2N1B1q1/3P4/R3K3 b"},
			"3b2n1/pp2p1p1/2k2r2/3P4/8/2N1B1q1/3P4/R3K3",
			chess.ColorBlack,
			false,
		},
		{
			"full_fen",
			args{"1k1r4/pp3p1p/3r4/6Q1/1P6/2N2P2/P4PKP/R7 b - - 0 22"},
			"1k1r4/pp3p1p/3r4/6Q1/1P6/2N2P2/P4PKP/R7",
			chess.ColorBlack,
			false,
		},
		{
			"6x6",
			args{"rnq1kb/pp1ppp/6/6/PPP1PP/RNQK1B w"},
			"rnq1kb/pp1ppp/6/6/PPP1PP/RNQK1B",
			chess.ColorWhite,
			false,
		},
		{
			"10x10",
			args{
				"1n2r1b1qk/3p3p1p/2n5B1/3P1N2r1/4Q3P1/3b1K3P/4p2B2/3q2P3/R3P1r3/2N2k4 b KQ - 10 34",
			},
			"1n2r1b1qk/3p3p1p/2n5B1/3P1N2r1/4Q3P1/3b1K3P/4p2B2/3q2P3/R3P1r3/2N2k4",
			chess.ColorBlack,
			false,
		},
		{
			"12x8",
			args{
				"2r1k2bnq1p/3p4p1P1/10pp/1P1b2N3p1/kQ10/3K3P2N1/R10R/12 w k h4 14 38",
			},
			"2r1k2bnq1p/3p4p1P1/10pp/1P1b2N3p1/kQ10/3K3P2N1/R10R/12",
			chess.ColorWhite,
			false,
		},
		{
			"7x13",
			args{
				"1q2b2/2p1P1n/1r2B2/2N2pk/3r1PB/1Q2n1K/3P3/1P3r1/2N1k1P/2p1bP1/1R2K2/2q2p1/2B1n2 b KQk g3 5 55",
			},
			"1q2b2/2p1P1n/1r2B2/2N2pk/3r1PB/1Q2n1K/3P3/1P3r1/2N1k1P/2p1bP1/1R2K2/2q2p1/2B1n2",
			chess.ColorBlack,
			false,
		},

		{
			"5x5",
			args{"r1k1b/ppp1p/5/5/PPPP1"},
			"",
			false,
			true,
		},
		{
			"digits",
			args{"2R2/221/11111/32/PPPP1"},
			"",
			false,
			true,
		},
		{
			"empty_string",
			args{""},
			"",
			false,
			true,
		},
		{
			"invalid",
			args{"hello w@rld"},
			"",
			false,
			true,
		},
		{
			"broken_position",
			args{""},
			"6/pN1K2/5/Pn1Q/ppppppp/kbnQBN",
			false,
			true,
		},
		{
			"empty_string",
			args{""},
			"",
			false,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fen.Decode(tt.args.fen)

			require.Truef(t, (err != nil) == tt.wantErr, "Decode() error = %v, wantErr %v", err, tt.wantErr)
			if tt.wantErr {
				return
			}

			assert.Equal(t, tt.wantColor, got.Turn())
			assert.Equal(t, tt.wantFEN, fen.EncodeSquares(got.Squares()))
		})
	}
}
