package fen

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/castling"
	"github.com/stretchr/testify/require"
)

func TestFromString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		str     string
		want    FEN
		wantErr bool
	}{
		{
			"full",
			"1k1r4/pp3p1p/3r4/6Q1/1P6/2N2P2/P4PKP/R7 b Kkq c9 0 22",
			FEN{
				placement: "1k1r4/pp3p1p/3r4/6Q1/1P6/2N2P2/P4PKP/R7",
				turn:      chess.ColorBlack,
				castlings: map[chess.Color]map[castling.CastlingType]bool{
					chess.ColorWhite: {
						castling.TypeShort: true,
						castling.TypeLong:  false,
					},
					chess.ColorBlack: {
						castling.TypeShort: true,
						castling.TypeLong:  true,
					},
				},
				enPassantSquare: chess.PositionFromString("c9"),
				halfmoveClock:   0,
				moveNumber:      22,
			},
			false,
		},
		{
			"negative",
			"1k1r4/pp3p1p/3r4/6Q1/1P6/2N22/P4PKP/R7X c JqIe - 1 -3",
			FEN{},
			true,
		},
		{
			"empty",
			"8/8/8/8/8/8/8/8 w - - 0 1",
			FEN{
				placement: "8/8/8/8/8/8/8/8",
				turn:      chess.ColorWhite,
				castlings: map[chess.Color]map[castling.CastlingType]bool{
					chess.ColorWhite: {
						castling.TypeShort: false,
						castling.TypeLong:  false,
					},
					chess.ColorBlack: {
						castling.TypeShort: false,
						castling.TypeLong:  false,
					},
				},
				enPassantSquare: chess.NewPositionEmpty(),
				halfmoveClock:   0,
				moveNumber:      1,
			},
			false,
		},
		{
			"12x8",
			"1q2b2/2p1P1n/1r2B2/2N2pk/3r1PB/1Q2n1K/3P3/1P3r1/2N1k1P/2p1bP1/1R2K2/2q2p1/2B1n2 b KQk g3 5 55",
			FEN{
				placement: "1q2b2/2p1P1n/1r2B2/2N2pk/3r1PB/1Q2n1K/3P3/1P3r1/2N1k1P/2p1bP1/1R2K2/2q2p1/2B1n2",
				turn:      chess.ColorBlack,
				castlings: map[chess.Color]map[castling.CastlingType]bool{
					chess.ColorWhite: {
						castling.TypeShort: true,
						castling.TypeLong:  true,
					},
					chess.ColorBlack: {
						castling.TypeShort: true,
						castling.TypeLong:  false,
					},
				},
				enPassantSquare: chess.PositionFromString("g3"),
				halfmoveClock:   5,
				moveNumber:      55,
			},
			false,
		},
		{
			"no_castlings",
			"rnq1kb/pp1ppp/6/6/PPP1PP/RNQK1B w - p13 0 1",
			FEN{
				placement: "rnq1kb/pp1ppp/6/6/PPP1PP/RNQK1B",
				turn:      chess.ColorWhite,
				castlings: map[chess.Color]map[castling.CastlingType]bool{
					chess.ColorWhite: {
						castling.TypeShort: false,
						castling.TypeLong:  false,
					},
					chess.ColorBlack: {
						castling.TypeShort: false,
						castling.TypeLong:  false,
					},
				},
				enPassantSquare: chess.PositionFromString("p13"),
				halfmoveClock:   0,
				moveNumber:      1,
			},
			false,
		},
		{
			"no_enpassant_target_square",
			"1r1q1n2/2P3b1/3k4/8/2N2B2/1p2P3/2r1n3/3Q1R2 w q - 14 101",
			FEN{
				placement: "1r1q1n2/2P3b1/3k4/8/2N2B2/1p2P3/2r1n3/3Q1R2",
				turn:      chess.ColorWhite,
				castlings: map[chess.Color]map[castling.CastlingType]bool{
					chess.ColorWhite: {
						castling.TypeShort: false,
						castling.TypeLong:  false,
					},
					chess.ColorBlack: {
						castling.TypeShort: false,
						castling.TypeLong:  true,
					},
				},
				enPassantSquare: chess.NewPositionEmpty(),
				halfmoveClock:   14,
				moveNumber:      101,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := FromString(tt.str)
			require.True(t, (gotErr != nil) == tt.wantErr)

			if tt.wantErr {
				return
			}

			require.Equal(t, tt.want, got)
		})
	}
}

func TestFEN_String(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		str string
	}{
		{
			"full",
			"1k1r4/pp3p1p/3r4/6Q1/1P6/2N2P2/P4PKP/R7 w Kkq c9 0 22",
		},
		{
			"no_castlings",
			"1n2r1b1qk/3p3p1p/2n5B1/3P1N2r1/4Q3P1/3b1K3P/4p2B2/3q2P3/R3P1r3/2N2k4 b - g10 10 34",
		},
		{
			"no_enpassant_target_square",
			"1k1r4/pp3p1p/3r4/6Q1/1P6/2N2P2/P4PKP/R7 b KQkq - 0 22",
		},
		{
			"no_castlings_and_enpassant_target_square",
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w - - 0 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := FromString(tt.str)
			require.NoError(t, err)
			require.Equal(t, tt.str, f.String())
		})
	}
}
