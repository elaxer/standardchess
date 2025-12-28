package mover_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/mover"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
)

func TestCastling_Make_Short(t *testing.T) {
	b, _ := standardchess.New(chess.SideWhite, nil)

	king := piece.NewKing(chess.SideWhite)
	rook := piece.NewRook(chess.SideWhite)

	b.Squares().PlacePiece(king, chess.PositionFromString("e1"))
	b.Squares().PlacePiece(rook, chess.PositionFromString("h1"))

	if got, err := new(mover.Castling).Make(move.CastlingShort, b); err != nil {
		t.Fatalf("castling failed: %v", err)
	} else if got == nil {
		t.Fatalf("expected a valid move result, got nil")
	}

	if pos := b.Squares().GetByPiece(king); pos != chess.PositionFromString("g1") {
		t.Errorf("king should be on g1, got %s", pos)
	}
	if pos := b.Squares().GetByPiece(rook); pos != chess.PositionFromString("f1") {
		t.Errorf("rook should be on f1, got %s", pos)
	}
}

func TestCastling_Make_Long(t *testing.T) {
	b, _ := standardchess.New(chess.SideWhite, nil)

	king := piece.NewKing(chess.SideWhite)
	rook := piece.NewRook(chess.SideWhite)

	b.Squares().PlacePiece(king, chess.PositionFromString("e1"))
	b.Squares().PlacePiece(rook, chess.PositionFromString("a1"))

	if got, err := new(mover.Castling).Make(move.CastlingLong, b); err != nil {
		t.Fatalf("castling failed: %v", err)
	} else if got == nil {
		t.Fatalf("expected a valid move result, got nil")
	}

	if pos := b.Squares().GetByPiece(king); pos != chess.PositionFromString("c1") {
		t.Errorf("king should be on c1, got %s", pos)
	}
	if pos := b.Squares().GetByPiece(rook); pos != chess.PositionFromString("d1") {
		t.Errorf("rook should be on d1, got %s", pos)
	}
}

func TestCastling_Make_Black(t *testing.T) {
	b, _ := standardchess.New(chess.SideBlack, nil)

	king := piece.NewKing(chess.SideBlack)
	rook := piece.NewRook(chess.SideBlack)

	b.Squares().PlacePiece(king, chess.PositionFromString("e8"))
	b.Squares().PlacePiece(rook, chess.PositionFromString("h8"))

	if got, err := new(mover.Castling).Make(move.CastlingShort, b); err != nil {
		t.Fatalf("castling failed: %v", err)
	} else if got == nil {
		t.Fatalf("expected a valid move result, got nil")
	}

	if pos := b.Squares().GetByPiece(king); pos != chess.PositionFromString("g8") {
		t.Errorf("king should be on g8, got %s", pos)
	}
	if pos := b.Squares().GetByPiece(rook); pos != chess.PositionFromString("f8") {
		t.Errorf("rook should be on f8, got %s", pos)
	}
}

func TestCastling_Make_Negative(t *testing.T) {
	_, err := new(mover.Castling).Make(move.CastlingShort,
		standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
			chess.PositionFromString("e1"): standardtest.NewPiece("K"),
			chess.PositionFromString("h1"): standardtest.NewPiece("R"),
			chess.PositionFromString("f1"): standardtest.NewPiece("N"),
		}))
	if err == nil {
		t.Errorf("Castling.Make() should have returned an error, got nil")
	}
}
