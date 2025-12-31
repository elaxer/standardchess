package mover_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/mover"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUndoNormal(t *testing.T) {
	piece := standardtest.NewPiece("Q")
	movedPiece := standardtest.NewPieceM("Q")

	pieceWillBeCaptured := standardtest.NewPiece("p")
	movedPieceWillBeCaptured := standardtest.NewPieceM("p")

	board := standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
		chess.PositionFromString("a1"): piece,
		chess.PositionFromString("h1"): movedPieceWillBeCaptured,

		chess.PositionFromString("h8"): movedPiece,
		chess.PositionFromString("a8"): pieceWillBeCaptured,
	})

	initFEN := standardtest.EncodeFEN(board)

	firstMoveResult, err := mover.MakeNormal(move.NewNormal(
		chess.PositionFromString("a1"),
		chess.PositionFromString("h1"),
		piece.Notation(),
	), board)
	require.NoError(t, err)

	secondMoveResult, err := mover.MakeNormal(move.NewNormal(
		chess.PositionFromString("h8"),
		chess.PositionFromString("a8"),
		movedPiece.Notation(),
	), board)
	require.NoError(t, err)

	err = mover.UndoNormal(secondMoveResult, board)
	require.NoError(t, err)

	err = mover.UndoNormal(firstMoveResult, board)
	require.NoError(t, err)

	assert.Equal(t, initFEN, standardtest.EncodeFEN(board))
	assert.False(t, piece.IsMoved())
	assert.True(t, movedPiece.IsMoved())
	assert.False(t, pieceWillBeCaptured.IsMoved())
	assert.True(t, movedPieceWillBeCaptured.IsMoved())
}
