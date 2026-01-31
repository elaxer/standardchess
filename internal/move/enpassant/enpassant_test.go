package enpassant_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/enpassant"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnPassantTargetSquare(t *testing.T) {
	board := standardtest.DecodeFEN("rnQ4r/pp2p1kp/3p2pn/1BpP2N1/5P2/2BK4/P1P3qP/8 b - - 5 18")
	_, err := board.MakeMove("e5")
	require.NoError(t, err)

	position := enpassant.EnPassantTargetSquare(board)
	require.Equal(t, chess.PositionFromString("e6"), position)

	pawn, err := board.Squares().FindByPosition(chess.PositionFromString("d5"))
	require.NoError(t, err)

	pawnMoves := board.LegalMoves(pawn)
	assert.Contains(t, pawnMoves, position)
}
