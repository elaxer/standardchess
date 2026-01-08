package normal_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/internal/move/normal"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUndoNormal(t *testing.T) {
	piece := standardtest.NewPiece("Q")
	movedPiece := standardtest.NewPieceM("Q")

	pieceWillBeCaptured := standardtest.NewPiece("p")
	movedPieceWillBeCaptured := standardtest.NewPieceM("p")

	board := standardtest.NewBoardEmpty8x8(chess.SideWhite, map[chess.Position]chess.Piece{
		chess.PositionFromString("a1"): piece,
		chess.PositionFromString("h1"): movedPieceWillBeCaptured,

		chess.PositionFromString("h8"): movedPiece,
		chess.PositionFromString("a8"): pieceWillBeCaptured,
	})

	initFEN := standardtest.EncodeFEN(board)

	firstMoveResult, err := normal.MakeMove(normal.NewMove(
		chess.PositionFromString("a1"),
		chess.PositionFromString("h1"),
		piece.Notation(),
	), board)
	firstMoveResult.SetBoardNewState(chess.StateClear)
	require.NoError(t, err)

	secondMoveResult, err := normal.MakeMove(normal.NewMove(
		chess.PositionFromString("h8"),
		chess.PositionFromString("a8"),
		movedPiece.Notation(),
	), board)
	secondMoveResult.SetBoardNewState(chess.StateClear)
	require.NoError(t, err)

	err = normal.UndoMove(secondMoveResult, board)
	require.NoError(t, err)

	err = normal.UndoMove(firstMoveResult, board)
	require.NoError(t, err)

	assert.Equal(t, initFEN, standardtest.EncodeFEN(board))
	assert.False(t, piece.IsMoved())
	assert.True(t, movedPiece.IsMoved())
	assert.False(t, pieceWillBeCaptured.IsMoved())
	assert.True(t, movedPieceWillBeCaptured.IsMoved())
}

func BenchmarkMakeMove(b *testing.B) {
	board := standardchess.NewBoard()

	move := normal.NewMove(
		chess.PositionFromString("e2"),
		chess.PositionFromString("e4"),
		piece.NotationPawn,
	)
	b.ResetTimer()
	for range b.N {
		_, err := normal.MakeMove(move, board)

		b.StopTimer()
		require.NoError(b, err)
		_, err = board.UndoLastMove()
		require.NoError(b, err)
		board = standardchess.NewBoard()
		b.StartTimer()
	}
}
