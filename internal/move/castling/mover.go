package castling

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/result"
	"github.com/elaxer/standardchess/internal/piece"
)

var ErrUndoMove = errors.New("cannot undo castling move")

func MakeMove(castlingType CastlingType, board chess.Board) (chess.MoveResult, error) {
	if err := ValidateMove(castlingType, board.Turn(), board, true); err != nil {
		return nil, err
	}

	king, kingPosition := board.Squares().FindPiece(piece.NotationKing, board.Turn())
	rook, rookPosition, _, err := getRook(fileDirection(castlingType), board.Turn(), board.Squares(), kingPosition)
	if err != nil {
		return nil, err
	}

	if err := board.Squares().PlacePiece(nil, kingPosition); err != nil {
		return nil, err
	}
	if err := board.Squares().PlacePiece(nil, rookPosition); err != nil {
		return nil, err
	}

	kingNewPosition, rookNewPosition := pickPositions(castlingType, kingPosition.Rank)

	if err := board.Squares().PlacePiece(king, kingNewPosition); err != nil {
		return nil, err
	}
	if err := board.Squares().PlacePiece(rook, rookNewPosition); err != nil {
		return nil, err
	}

	king.SetIsMoved(true)
	rook.SetIsMoved(true)

	return &MoveResult{
		Abstract:         result.NewAbstract(board),
		CastlingType:     castlingType,
		InitKingPosition: kingPosition,
		InitRookPosition: rookPosition,
	}, nil
}

func UndoMove(move *MoveResult, board chess.Board) error {
	if err := move.Validate(); err != nil {
		return err
	}

	rank := chess.Rank1
	if move.Side().IsBlack() {
		rank = chess.Rank8
	}

	kingPosition, rookPosition := pickPositions(move.CastlingType, rank)

	king, err := board.Squares().FindByPosition(kingPosition)
	if err != nil {
		return err
	}
	if king == nil {
		return fmt.Errorf("%w: cannot find the king", ErrUndoMove)
	}

	rook, err := board.Squares().FindByPosition(rookPosition)
	if err != nil {
		return err
	}
	if rook == nil {
		return fmt.Errorf("%w: cannot find the rook", ErrUndoMove)
	}

	if err := board.Squares().PlacePiece(nil, kingPosition); err != nil {
		return err
	}
	if err := board.Squares().PlacePiece(nil, rookPosition); err != nil {
		return err
	}

	king.SetIsMoved(false)
	if err := board.Squares().PlacePiece(king, move.InitKingPosition); err != nil {
		return err
	}

	rook.SetIsMoved(false)

	return board.Squares().PlacePiece(rook, move.InitRookPosition)
}

func pickPositions(castlingType CastlingType, rank chess.Rank) (kingPosition, rookPosition chess.Position) {
	if castlingType.IsLong() {
		return chess.NewPosition(kingFileAfterLongCastling, rank), chess.NewPosition(rookFileAfterLongCastling, rank)
	}

	return chess.NewPosition(kingFileAfterShortCastling, rank), chess.NewPosition(rookFileAfterShortCastling, rank)
}
