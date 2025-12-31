package mover

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	mv "github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/resolver"
	"github.com/elaxer/standardchess/internal/move/result"
	"github.com/elaxer/standardchess/internal/move/validator"
)

var (
	Err         = errors.New("mover error")
	ErrMakeMove = fmt.Errorf("%w: cannot make move", Err)
	ErrUndoMove = fmt.Errorf("%w: cannot undo move", Err)
)

func MakeMove(move chess.Move, board chess.Board) (chess.MoveResult, error) {
	str := move.String()

	if move, err := mv.NormalFromString(str); err == nil {
		return MakeNormal(move, board)
	}
	if move, err := mv.PromotionFromString(str); err == nil {
		return MakePromotion(move, board)
	}
	if move, err := mv.CastlingFromString(str); err == nil {
		return MakeCastling(move, board)
	}

	return nil, fmt.Errorf("%w: invalid move \"%s\"", ErrMakeMove, str)
}

func UndoMove(move chess.MoveResult, board chess.Board) error {
	switch moveResult := move.(type) {
	case *result.Normal:
		return UndoNormal(moveResult, board)
	case *result.Promotion:
		return UndoPromotion(moveResult, board)
	case *result.EnPassant:
		return nil
	case *result.Castling:
		return UndoCastling(moveResult, board)
	default:
		return fmt.Errorf("%w: unknown move to undo", ErrUndoMove)
	}
}

func movePiece(move mv.PieceMove, movingPieceNotation string, board chess.Board) (result.PieceMove, error) {
	fullFrom, err := resolver.ResolveFrom(move, movingPieceNotation, board, board.Turn())
	if err != nil {
		return result.PieceMove{}, err
	}
	pieceMove := mv.NewPieceMove(fullFrom, move.To)

	shortenedFrom, err := resolver.UnresolveFrom(pieceMove, board)
	if err != nil {
		return result.PieceMove{}, err
	}

	if err := validator.ValidatePieceMove(pieceMove, movingPieceNotation, board); err != nil {
		return result.PieceMove{}, err
	}

	capturedPiece, err := board.Squares().MovePiece(pieceMove.From, pieceMove.To)
	if err != nil {
		return result.PieceMove{}, err
	}

	piece, err := board.Squares().FindByPosition(pieceMove.To)
	if err != nil {
		return result.PieceMove{}, err
	}

	wasMoved := piece.IsMoved()

	piece.SetIsMoved(true)

	return result.PieceMove{
		WasMoved:      wasMoved,
		FromFull:      fullFrom,
		FromShortened: shortenedFrom,
		Captured:      capturedPiece,
	}, nil
}

func newAbstractResult(board chess.Board) result.Abstract {
	return result.Abstract{
		MoveSide: board.Turn(),
		NewState: board.State(!board.Turn()),
	}
}
