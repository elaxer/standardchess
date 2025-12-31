package standardchess_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_board_MakeMove_CaptureAddsToCapturedPieces(t *testing.T) {
	b := standardtest.DecodeFEN("r3k2r/ppp2ppp/B1n2n1B/3pp2Q/3PP2q/b1N2N1b/PPP2PPP/R3K2R")
	require.Equal(t, 0, len(b.CapturedPieces()))

	cases := []struct {
		move          string
		expectCapture bool
	}{
		{"Qh4", true},
		{"Bb2", true},
		{"Nd5", true},
		{"0-0", false},
	}

	capturedCount := 0
	for _, c := range cases {
		res, err := b.MakeMove(chess.StringMove(c.move))
		require.NoErrorf(t, err, "MakeMove failed (%s)", c.move)

		if !c.expectCapture {
			require.Nilf(t, res.CapturedPiece(), "expected no capture for move %s, but CapturedPiece is %v", c.move, res.CapturedPiece())
			require.Equal(t, capturedCount, len(b.CapturedPieces()))

			continue
		}

		require.NotNilf(t, res.CapturedPiece(), "expected move %s to capture a piece, but CapturedPiece is nil", c.move)
		capturedCount++

		require.Equal(t, capturedCount, len(b.CapturedPieces()))
		require.Equalf(t, b.CapturedPieces()[capturedCount-1], res.CapturedPiece(), "after %s captured piece in board does not match move result", c.move)
	}
}

func Test_board_UndoLastMove(t *testing.T) {
	moves := []string{
		"e4", "c6",
		"d4", "d5",
		"e5", "Bf5",
		"Nc3", "e6",
		"g4", "Bg6",
		"Nge2", "c5",
		"Be3", "Ne7",
		"f4", "h5",
		"f5", "exf5",
		"g5", "Nbc6",
		"Nf4", "a6",
		"Bg2", "cxd4",
		"Bxd4", "Nxd4",
		"Qxd4", "Nc6",
		"Qf2", "Bb4",
		"0-0-0", "Bxc3",
		"bxc3", "Qa5",
		"Rxd5", "Qxc3",
		"Qc5", "Qxc5",
		"Rxc5", "0-0",
		"Bxc6", "bxc6",
		"Rd1", "Rab8",
		"c4", "Rfd8",
		"Rd6", "Kf8",
		"Rcc6", "Rdc8",
		"Kc2", "h4",
		"Rxc8+", "Rxc8",
		"Kc3", "a5",
		"Ra6", "Rb8",
		"Rxa5", "Rb1",
		"c5", "Re1",
		"Ra8+", "Ke7",
		"Ra7+", "Ke8",
		"Nd3", "Re3",
		"Kd2", "Rh3",
		"c6", "Rxh2+",
		"Ke3", "Rc2",
		"e6", "h3",
		"Nb4", "f4+",
		"Kd4", "h2",
		"Ra8+", "Ke7",
		"Rh8", "Rd2+",
		"Kc5", "Be4",
		"c7", "Bb7",
		"Kb6", "Bc8",
		"Rxc8", "h1=Q",
		"Re8+", "Kxe8",
		"c8=Q+", "Ke7",
		"Nc6+", "Qxc6+",
		"Qxc6", "Rd6",
	}

	board := standardchess.NewBoardFilled()
	initFEN := standardtest.EncodeFEN(board)

	for _, move := range moves {
		_, err := board.MakeMove(chess.StringMove(move))
		require.NoError(t, err)
	}
	for i := range board.MoveHistory() {
		_, err := board.UndoLastMove()
		require.NoErrorf(t, err, "No %d", i+1)
	}

	afterFEN := standardtest.EncodeFEN(board)
	assert.Equal(t, initFEN, afterFEN)
}

func TestNewBoardFromMoves(t *testing.T) {
	type args struct {
		moves []string
	}
	tests := []struct {
		name    string
		args    args
		wantFEN string
	}{
		{
			"https://www.chess.com/games/view/14842105",
			args{[]string{
				"e4", "c6",
				"d4", "d5",
				"e5", "Bf5",
				"Nc3", "e6",
				"g4", "Bg6",
				"Nge2", "c5",
				"Be3", "Ne7",
				"f4", "h5",
				"f5", "exf5",
				"g5", "Nbc6",
				"Nf4", "a6",
				"Bg2", "cxd4",
				"Bxd4", "Nxd4",
				"Qxd4", "Nc6",
				"Qf2", "Bb4",
				"0-0-0", "Bxc3",
				"bxc3", "Qa5",
				"Rxd5", "Qxc3",
				"Qc5", "Qxc5",
				"Rxc5", "0-0",
				"Bxc6", "bxc6",
				"Rd1", "Rab8",
				"c4", "Rfd8",
				"Rd6", "Kf8",
				"Rcc6", "Rdc8",
				"Kc2", "h4",
				"Rxc8+", "Rxc8",
				"Kc3", "a5",
				"Ra6", "Rb8",
				"Rxa5", "Rb1",
				"c5", "Re1",
				"Ra8+", "Ke7",
				"Ra7+", "Ke8",
				"Nd3", "Re3",
				"Kd2", "Rh3",
				"c6", "Rxh2+",
				"Ke3", "Rc2",
				"e6", "h3",
				"Nb4", "f4+",
				"Kd4", "h2",
				"Ra8+", "Ke7",
				"Rh8", "Rd2+",
				"Kc5", "Be4",
				"c7", "Bb7",
				"Kb6", "Bc8",
				"Rxc8", "h1=Q",
				"Re8+", "Kxe8",
				"c8=Q+", "Ke7",
				"Nc6+", "Qxc6+",
				"Qxc6", "Rd6",
			}},
			"8/4kpp1/1KQrP3/6P1/5p2/8/P7/8 w - - 1 51",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := standardtest.NewBoardFromMoves(tt.args.moves...)
			fen := standardtest.EncodeFEN(board)

			assert.Equalf(t, tt.wantFEN, fen, "Expected position \"%s\", got - %s", tt.wantFEN, fen)
		})
	}
}
