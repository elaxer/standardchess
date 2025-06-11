package mover_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
	"github.com/elaxer/standardchess/board"
	"github.com/elaxer/standardchess/move/move"
	"github.com/elaxer/standardchess/move/mover"
	"github.com/elaxer/standardchess/piece"
	"github.com/elaxer/standardchess/standardtest"
)

func TestCastling_Make_Short(t *testing.T) {
	b, _ := board.NewFactory().Create(chess.SideWhite, nil)

	king := piece.NewKing(chess.SideWhite)
	rook := piece.NewRook(chess.SideWhite)

	b.Squares().PlacePiece(king, position.FromString("e1"))
	b.Squares().PlacePiece(rook, position.FromString("h1"))

	if got, err := new(mover.Castling).Make(move.CastlingShort, b); err != nil {
		t.Fatalf("castling failed: %v", err)
	} else if got == nil {
		t.Fatalf("expected a valid move result, got nil")
	}

	if pos := b.Squares().GetByPiece(king); pos != position.FromString("g1") {
		t.Errorf("king should be on g1, got %s", pos)
	}
	if pos := b.Squares().GetByPiece(rook); pos != position.FromString("f1") {
		t.Errorf("rook should be on f1, got %s", pos)
	}
}

func TestCastling_Make_Long(t *testing.T) {
	b, _ := board.NewFactory().Create(chess.SideWhite, nil)

	king := piece.NewKing(chess.SideWhite)
	rook := piece.NewRook(chess.SideWhite)

	b.Squares().PlacePiece(king, position.FromString("e1"))
	b.Squares().PlacePiece(rook, position.FromString("a1"))

	if got, err := new(mover.Castling).Make(move.CastlingLong, b); err != nil {
		t.Fatalf("castling failed: %v", err)
	} else if got == nil {
		t.Fatalf("expected a valid move result, got nil")
	}

	if pos := b.Squares().GetByPiece(king); pos != position.FromString("c1") {
		t.Errorf("king should be on c1, got %s", pos)
	}
	if pos := b.Squares().GetByPiece(rook); pos != position.FromString("d1") {
		t.Errorf("rook should be on d1, got %s", pos)
	}
}

func TestCastling_Make_Black(t *testing.T) {
	b, _ := board.NewFactory().Create(chess.SideBlack, nil)

	king := piece.NewKing(chess.SideBlack)
	rook := piece.NewRook(chess.SideBlack)

	b.Squares().PlacePiece(king, position.FromString("e8"))
	b.Squares().PlacePiece(rook, position.FromString("h8"))

	if got, err := new(mover.Castling).Make(move.CastlingShort, b); err != nil {
		t.Fatalf("castling failed: %v", err)
	} else if got == nil {
		t.Fatalf("expected a valid move result, got nil")
	}

	if pos := b.Squares().GetByPiece(king); pos != position.FromString("g8") {
		t.Errorf("king should be on g8, got %s", pos)
	}
	if pos := b.Squares().GetByPiece(rook); pos != position.FromString("f8") {
		t.Errorf("rook should be on f8, got %s", pos)
	}
}

func TestCastling_Make_Negative(t *testing.T) {
	_, err := new(mover.Castling).Make(move.CastlingShort,
		standardtest.NewBoard(chess.SideWhite, map[position.Position]chess.Piece{
			position.FromString("e1"): standardtest.NewPiece("K"),
			position.FromString("h1"): standardtest.NewPiece("R"),
			position.FromString("f1"): standardtest.NewPiece("N"),
		}))
	if err == nil {
		t.Errorf("Castling.Make() should have returned an error, got nil")
	}
}
