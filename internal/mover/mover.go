// Package mover contains code for executing and canceling any type of move.
package mover

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/castling"
	"github.com/elaxer/standardchess/internal/move/enpassant"
	"github.com/elaxer/standardchess/internal/move/normal"
	"github.com/elaxer/standardchess/internal/move/promotion"
	"github.com/elaxer/standardchess/internal/piece"
)

var (
	Err         = errors.New("mover error")
	ErrMakeMove = fmt.Errorf("%w: cannot make move", Err)
	ErrUndoMove = fmt.Errorf("%w: cannot undo move", Err)
)

func MakeMove(move chess.Move, board chess.Board) (chess.MoveResult, error) {
	str := move.String()

	if move, err := normal.MoveFromString(str); err == nil {
		isPawn := move.PieceNotation == piece.NotationPawn
		if enpassant.CanEnPassant(board) && isPawn &&
			move.To == enpassant.EnPassantPosition(board) {
			return enpassant.MakeMove(enpassant.NewEnPassant(move.From, move.To), board)
		}

		return normal.MakeMove(move, board)
	}
	if move, err := promotion.MoveFromString(str); err == nil {
		return promotion.MakeMove(move, board)
	}
	if move, err := castling.TypeFromString(str); err == nil {
		return castling.MakeMove(move, board)
	}

	return nil, fmt.Errorf("%w: invalid move \"%s\"", ErrMakeMove, str)
}

func UndoMove(move chess.MoveResult, board chess.Board) error {
	switch move := move.(type) {
	case *normal.MoveResult:
		return normal.UndoMove(move, board)
	case *promotion.MoveResult:
		return promotion.UndoPromotion(move, board)
	case *enpassant.MoveResult:
		return enpassant.UndoMove(move, board)
	case *castling.MoveResult:
		return castling.UndoMove(move, board)
	default:
		return fmt.Errorf("%w: unknown move to undo", ErrUndoMove)
	}
}
