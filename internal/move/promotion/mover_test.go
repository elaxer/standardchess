package promotion_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/promotion"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMakePromotion_White(t *testing.T) {
	board := standardtest.NewBoardEmpty8x8(chess.SideWhite, map[chess.Position]chess.Piece{
		chess.PositionFromString("d7"): standardtest.NewPiece("P"),
		chess.PositionFromString("a1"): standardtest.NewPiece("K"),
		chess.PositionFromString("a8"): standardtest.NewPiece("k"),
	})

	move := promotion.NewMove(chess.NewPositionEmpty(), chess.PositionFromString("d8"), piece.NotationQueen)

	promotionRes, err := promotion.MakeMove(move, board)
	require.NoError(t, err)
	require.NotNil(t, promotionRes)

	queen, err := board.Squares().FindByPosition(chess.PositionFromString("d8"))
	require.NoError(t, err)
	require.NotNil(t, queen)
	assert.Equal(t, piece.NotationQueen, queen.Notation())
}

func TestMakePromotion_Black(t *testing.T) {
	board := standardtest.NewBoardEmpty8x8(chess.SideBlack, map[chess.Position]chess.Piece{
		chess.PositionFromString("d2"): standardtest.NewPiece("p"),
		chess.PositionFromString("a8"): standardtest.NewPiece("K"),
		chess.PositionFromString("a1"): standardtest.NewPiece("k"),
		chess.PositionFromString("c1"): standardtest.NewPiece("B"),
	})

	move := promotion.NewMove(chess.NewPositionEmpty(), chess.PositionFromString("c1"), piece.NotationRook)

	promotionRes, err := promotion.MakeMove(move, board)
	require.NoError(t, err)
	require.NotNil(t, promotionRes)

	rook, err := board.Squares().FindByPosition(chess.PositionFromString("c1"))
	require.NoError(t, err)
	require.NotNil(t, rook)
	assert.Equal(t, piece.NotationRook, rook.Notation())
}
