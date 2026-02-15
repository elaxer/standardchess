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
	Err = errors.New("mover error")

	ErrMakeMove            = fmt.Errorf("%w: cannot make move", Err)
	errMakeMoveUnknownType = fmt.Errorf("%w: unknown move type", ErrMakeMove)

	ErrUndoMove            = fmt.Errorf("%w: cannot undo move", Err)
	errUndoMoveUnknownType = fmt.Errorf("%w: unknown move type", ErrUndoMove)
)

func MakeMove(moveStr string, board chess.Board) (chess.Move, error) {
	move, err := moveFromUCI(moveStr, board)
	if err != nil {
		move, err = moveFromNotation(moveStr)
	}

	if err != nil {
		return nil, err
	}

	switch move := move.(type) {
	case *normal.Move:
		isPawn := move.PieceNotation == piece.NotationPawn
		if enpassant.CanEnPassant(board) && isPawn &&
			move.To == enpassant.EnPassantTargetSquare(board) {
			return enpassant.MakeMove(enpassant.NewEnPassant(move.From, move.To), board)
		}

		return normal.MakeMove(move, board)

	case *promotion.Move:
		return promotion.MakeMove(move, board)
	case castling.CastlingType:
		return castling.MakeMove(move, board)
	}

	return nil, errMakeMoveUnknownType
}

func moveFromUCI(uciStr string, board chess.Board) (any, error) {
	if castlingType, err := castling.FromUCI(uciStr, board.Squares(), board.Turn()); err == nil {
		return castlingType, nil
	}
	if move, err := promotion.MoveFromUCI(uciStr); err == nil {
		return move, nil
	}
	if move, err := normal.MoveFromUCI(uciStr, board.Squares()); err == nil {
		return move, nil
	}

	return nil, fmt.Errorf("%w: invalid move \"%s\"", ErrMakeMove, uciStr)
}

func moveFromNotation(notation string) (any, error) {
	if move, err := normal.MoveFromNotation(notation); err == nil {
		return move, nil
	}
	if move, err := promotion.MoveFromNotation(notation); err == nil {
		return move, nil
	}
	if castlingType, err := castling.FromNotation(notation); err == nil {
		return castlingType, nil
	}

	return nil, fmt.Errorf("%w: invalid move \"%s\"", ErrMakeMove, notation)
}

func UndoMove(move chess.Move, board chess.Board) error {
	switch move := move.(type) {
	case *normal.MoveResult:
		return normal.UndoMove(move, board)
	case *promotion.MoveResult:
		return promotion.UndoMove(move, board)
	case *enpassant.MoveResult:
		return enpassant.UndoMove(move, board)
	case *castling.MoveResult:
		return castling.UndoMove(move, board)
	default:
		return errUndoMoveUnknownType
	}
}
