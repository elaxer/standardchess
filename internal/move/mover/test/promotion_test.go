package mover_test

import (
	"testing"

	"github.com/elaxer/chess"
	. "github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/mover"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
)

func TestPromotion_Make(t *testing.T) {
	board := standardtest.NewBoard(SideWhite, map[chess.Position]Piece{
		chess.PositionFromString("d7"): standardtest.NewPiece("P"),
		chess.PositionFromString("a1"): standardtest.NewPiece("K"),
		chess.PositionFromString("a8"): standardtest.NewPiece("k"),
	})

	promotion := move.NewPromotion(chess.NewEmptyPosition(), chess.PositionFromString("d8"), piece.NotationQueen)
	_, err := new(mover.Promotion).Make(promotion, board)
	if err != nil {
		t.Fatalf("promotion failed: %v", err)
	}

	queen, err := board.Squares().FindByPosition(chess.PositionFromString("d8"))
	if err != nil {
		t.Fatal(err)
	}
	if queen == nil {
		t.Fatalf("the queen didn't appear on the board")
	}
	if queen.Notation() != piece.NotationQueen {
		t.Errorf("the piece should be a queen")
	}
}
