package standardchess_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/require"
)

func TestMakeMove_CaptureAddsToCapturedPieces(t *testing.T) {
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
