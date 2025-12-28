package standardchess

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/standardtest"
)

func TestMakeMove_CaptureAddsToCapturedPieces(t *testing.T) {
	b := standardtest.DecodeFEN("r3k2r/ppp2ppp/B1n2n1B/3pp2Q/3PP2q/b1N2N1b/PPP2PPP/R3K2R")
	if got := len(b.CapturedPieces()); got != 0 {
		t.Fatalf("expected 0 captured pieces initially, got %d", got)
	}

	type moveCheck struct {
		move          string
		expectCapture bool
	}

	cases := []moveCheck{
		{"Qh4", true},
		{"Bb2", true},
		{"Nd5", true},
		{"0-0", false},
	}

	capturedCount := 0
	for _, c := range cases {
		res, err := b.MakeMove(chess.StringMove(c.move))
		if err != nil {
			t.Fatalf("MakeMove failed (%s): %v", c.move, err)
		}

		if c.expectCapture {
			if res.CapturedPiece() == nil {
				t.Fatalf("expected move %s to capture a piece, but CapturedPiece is nil", c.move)
			}
			capturedCount++
			if len(b.CapturedPieces()) != capturedCount {
				t.Fatalf("after %s expected %d captured pieces, got %d", c.move, capturedCount, len(b.CapturedPieces()))
			}
			if b.CapturedPieces()[capturedCount-1] != res.CapturedPiece() {
				t.Fatalf("after %s captured piece in board (%v) does not match move result (%v)", c.move, b.CapturedPieces()[capturedCount-1], res.CapturedPiece())
			}
			continue
		}
		if res.CapturedPiece() != nil {
			t.Fatalf("expected no capture for move %s, but CapturedPiece is %v", c.move, res.CapturedPiece())
		}
		if len(b.CapturedPieces()) != capturedCount {
			t.Fatalf("after %s expected captured count to remain %d, got %d", c.move, capturedCount, len(b.CapturedPieces()))
		}
	}
}
