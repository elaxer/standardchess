package enpassant_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/enpassant"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateMove(t *testing.T) {
	board := standardtest.DecodeFEN("N7/pp1b2k1/3P1nqp/1Pb1p3/P1B1Pp1n/2N2P2/4Q1PK/R6R w - - 1 27")
	_, err := board.MakeMove("g4")
	require.NoError(t, err)

	assert.NoError(
		t,
		enpassant.ValidateMove(
			chess.PositionFromString("f4"),
			chess.PositionFromString("g3"),
			board,
		),
	)
}

func TestValidateMove_CheckAfterMove(t *testing.T) {
	board := standardtest.DecodeFEN("rnbqkbnr/ppp2ppp/8/3P4/8/8/PPP2PPP/RNBK1BNR b kq - 0 1")
	_, err := board.MakeMove("c5")
	require.NoError(t, err)

	err = enpassant.ValidateMove(
		chess.PositionFromString("d5"),
		chess.PositionFromString("c6"),
		board,
	)
	require.Error(t, err)
	assert.Equal(t, "en passant move validation error: there is a check after the move", err.Error())
}
