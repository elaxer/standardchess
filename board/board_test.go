package board

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/encoding/fen"
	"github.com/elaxer/standardchess/piece"
)

func TestMakeMove_CaptureAddsToCapturedPieces(t *testing.T) {
	fenStr := "r3k2r/ppp2ppp/B1n2n1B/3pp2Q/3PP2q/b1N2N1b/PPP2PPP/R3K2R w KQkq - 0 1"
	b, err := fen.NewDecoder(NewFactory(), piece.NewFactory()).Decode(fenStr)
	if err != nil {
		t.Fatalf("failed to create board from FEN: %v", err)
	}

	sb := b.(*board)

	if got := len(sb.capturedPieces); got != 0 {
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
			if len(sb.capturedPieces) != capturedCount {
				t.Fatalf("after %s expected %d captured pieces, got %d", c.move, capturedCount, len(sb.capturedPieces))
			}
			if sb.capturedPieces[capturedCount-1] != res.CapturedPiece() {
				t.Fatalf("after %s captured piece in board (%v) does not match move result (%v)", c.move, sb.capturedPieces[capturedCount-1], res.CapturedPiece())
			}
			continue
		}
		if res.CapturedPiece() != nil {
			t.Fatalf("expected no capture for move %s, but CapturedPiece is %v", c.move, res.CapturedPiece())
		}
		if len(sb.capturedPieces) != capturedCount {
			t.Fatalf("after %s expected captured count to remain %d, got %d", c.move, capturedCount, len(sb.capturedPieces))
		}
	}
}
