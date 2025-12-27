package fen_test

import (
	"os"
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/chesstest"
	"github.com/elaxer/chess/position"
	"github.com/elaxer/chess/visualizer"
	"github.com/elaxer/standardchess/encoding/fen"
)

func TestDecoder_Decode(t *testing.T) {
	type args struct {
		fen          string
		edgePosition position.Position
	}
	tests := []struct {
		name     string
		args     args
		wantFEN  string
		wantSide chess.Side
		wantErr  bool
	}{
		{
			"empty",
			args{"8/8/8/8/8/8/8/8", position.FromString("h8")},
			"8/8/8/8/8/8/8/8",
			chess.SideWhite,
			false,
		},
		{
			"8x8_init_position",
			args{"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", position.FromString("h8")},
			"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR",
			chess.SideWhite,
			false,
		},
		{
			"white",
			args{"1r1q1n2/2P3b1/3k4/8/2N2B2/1p2P3/2r1n3/3Q1R2 w", position.FromString("h8")},
			"1r1q1n2/2P3b1/3k4/8/2N2B2/1p2P3/2r1n3/3Q1R2",
			chess.SideWhite,
			false,
		},
		{
			"black",
			args{"3b2n1/pp2p1p1/2k2r2/3P4/8/2N1B1q1/3P4/R3K3 b", position.FromString("h8")},
			"3b2n1/pp2p1p1/2k2r2/3P4/8/2N1B1q1/3P4/R3K3",
			chess.SideBlack,
			false,
		},
		{
			"full_fen",
			args{"1k1r4/pp3p1p/3r4/6Q1/1P6/2N2P2/P4PKP/R7 b - - 0 22", position.FromString("h8")},
			"1k1r4/pp3p1p/3r4/6Q1/1P6/2N2P2/P4PKP/R7",
			chess.SideBlack,
			false,
		},
		{
			"6x6",
			args{"rnq1kb/pp1ppp/6/6/PPP1PP/RNQK1B w", position.FromString("f6")},
			"rnq1kb/pp1ppp/6/6/PPP1PP/RNQK1B",
			chess.SideWhite,
			false,
		},
		{
			"10x10",
			args{
				"1n2r1b1qk/3p3p1p/2n5B1/3P1N2r1/4Q3P1/3b1K3P/4p2B2/3q2P3/R3P1r3/2N2k4",
				position.FromString("j10"),
			},
			"1n2r1b1qk/3p3p1p/2n5B1/3P1N2r1/4Q3P1/3b1K3P/4p2B2/3q2P3/R3P1r3/2N2k4",
			chess.SideWhite,
			false,
		},
		{
			"12x8",
			args{
				"2r1k2bnq1p/3p4p1P1/2n3B1P3/1P1b2N3p1/2Q1P1p3R1/3K3P2N1/1p2B2P2p1/12",
				position.FromString("l8"),
			},
			"2r1k2bnq1p/3p4p1P1/2n3B1P3/1P1b2N3p1/2Q1P1p3R1/3K3P2N1/1p2B2P2p1/12",
			chess.SideWhite,
			false,
		},
		{
			"7x13",
			args{
				"1q2b2/2p1P1n/1r2B2/2N2pk/3r1PB/1Q2n1K/3P3/1P3r1/2N1k1P/2p1bP1/1R2K2/2q2p1/2B1n2",
				position.FromString("g13"),
			},
			"1q2b2/2p1P1n/1r2B2/2N2pk/3r1PB/1Q2n1K/3P3/1P3r1/2N1k1P/2p1bP1/1R2K2/2q2p1/2B1n2",
			chess.SideWhite,
			false,
		},

		{
			"5x5",
			args{"r1k1b/ppp1p/5/5/PPPP1", position.FromString("e5")},
			"",
			false,
			true,
		},
		{
			"digits",
			args{"2R2/221/11111/32/PPPP1", position.FromString("e5")},
			"",
			false,
			true,
		},
		{
			"empty_string",
			args{"", position.FromString("e5")},
			"",
			false,
			true,
		},
		{
			"invalid",
			args{"hello w@rld", position.FromString("e5")},
			"",
			false,
			true,
		},
		{
			"broken_position",
			args{"", position.FromString("f6")},
			"6/pN1K2/5/Pn1Q/ppppppp/kbnQBN",
			false,
			true,
		},
		{
			"empty_string",
			args{"", position.FromString("h8")},
			"",
			false,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			boardFactory := &chesstest.FactoryMock{EdgePosition: tt.args.edgePosition}
			decoder := fen.NewDecoder(boardFactory)
			got, err := decoder.Decode(tt.args.fen)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.Decode() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if tt.wantErr {
				return
			}

			if got.Turn() != tt.wantSide {
				t.Errorf("Decoder.Decode() got side = %v, want %v", got.Turn(), tt.wantSide)
			}
			if gotFEN := fen.EncodePiecePlacement(got.Squares()); gotFEN != tt.wantFEN {
				new(visualizer.Visualizer).Fprintln(os.Stdout, got)
				t.Errorf("Decoder.Decode() got =\n%v, want\n%v", gotFEN, tt.wantFEN)
			}
		})
	}
}
