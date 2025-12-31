package mover

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/result"
	"github.com/elaxer/standardchess/internal/move/validator"
	"github.com/elaxer/standardchess/internal/piece"
)

var (
	kingFileAfterShortCastling = chess.FileG
	rookFileAfterShortCastling = chess.FileF

	kingFileAfterLongCastling = chess.FileC
	rookFileAfterLongCastling = chess.FileD
)

func MakeCastling(castlingType move.Castling, board chess.Board) (chess.MoveResult, error) {
	if err := validator.ValidateCastling(castlingType, board.Turn(), board, true); err != nil {
		return nil, err
	}

	king, kingPosition := board.Squares().FindPiece(piece.NotationKing, board.Turn())
	rook, rookPosition, err := getRook(fileDirection(castlingType), board.Squares(), kingPosition)
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

	return &result.Castling{
		Abstract:         newAbstractResult(board),
		Castling:         castlingType,
		InitKingPosition: kingPosition,
		InitRookPosition: rookPosition,
	}, nil
}

func UndoCastling(castling *result.Castling, board chess.Board) error {
	if err := castling.Validate(); err != nil {
		return err
	}

	rank := chess.Rank1
	if castling.Side().IsBlack() {
		rank = chess.Rank8
	}

	kingPosition, rookPosition := pickPositions(castling.Castling, rank)

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
	if err := board.Squares().PlacePiece(king, castling.InitKingPosition); err != nil {
		return err
	}

	rook.SetIsMoved(false)

	return board.Squares().PlacePiece(rook, castling.InitRookPosition)
}

func getRook(fileDir chess.File, squares *chess.Squares, kingPosition chess.Position) (chess.Piece, chess.Position, error) {
	for position, p := range squares.IterByDirection(kingPosition, chess.NewPosition(fileDir, 0)) {
		if p != nil && p.Notation() == piece.NotationRook {
			return p, position, nil
		}
	}

	return nil, chess.NewPositionEmpty(), fmt.Errorf("%w: rook wasn't found", validator.ErrCastling)
}

func fileDirection(castlingType move.Castling) chess.File {
	if castlingType.IsLong() {
		return -1
	}

	return 1
}

func pickPositions(castlingType move.Castling, rank chess.Rank) (kingPos, rookPos chess.Position) {
	kingFile := kingFileAfterShortCastling
	if castlingType == move.CastlingLong {
		kingFile = kingFileAfterLongCastling
	}
	rookFile := rookFileAfterShortCastling
	if castlingType == move.CastlingLong {
		rookFile = rookFileAfterLongCastling
	}

	return chess.NewPosition(kingFile, rank), chess.NewPosition(rookFile, rank)
}
