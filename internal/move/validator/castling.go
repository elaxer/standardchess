package validator

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/piece"
)

var ErrCastling = fmt.Errorf("%w: castling validation error", Err)

func ValidateCastlingMove(castlingType move.Castling, side chess.Side, board chess.Board, validateObstacle bool) error {
	king, kingPosition := board.Squares().FindPiece(piece.NotationKing, side)
	if king == nil {
		return fmt.Errorf("%w: the king wasn't found", ErrCastling)
	}
	if king.IsMoved() {
		return fmt.Errorf("%w: the king already has been moved", ErrCastling)
	}
	if !board.State(side).Type().IsClear() {
		return fmt.Errorf("%w: the king is under threat", ErrCastling)
	}

	rook, err := board.Squares().FindByPosition(castlingRookPosition(castlingType, kingPosition.Rank))
	if err != nil {
		return err
	}
	if rook == nil || rook.Notation() != piece.NotationRook {
		return fmt.Errorf("%w: the rook wasn't found", ErrCastling)
	}
	if rook.IsMoved() {
		return fmt.Errorf("%w: the rook already has been moved", ErrCastling)
	}

	direction := fileDirection(castlingType)

	if validateObstacle {
		if err := castlingValidateObstacle(direction, board.Squares(), kingPosition, rook); err != nil {
			return err
		}
	}

	positions := mapset.NewSet(
		chess.NewPosition(kingPosition.File+direction, kingPosition.Rank),
		chess.NewPosition(kingPosition.File+direction*2, kingPosition.Rank),
	)
	if board.Moves(!side).Intersect(positions).Cardinality() > 0 {
		return fmt.Errorf("%w: castling squares are under threat", ErrCastling)
	}

	return nil
}

func castlingValidateObstacle(direction chess.File, squares *chess.Squares, kingPosition chess.Position, castlingRook chess.Piece) error {
	for _, piece := range squares.IterByDirection(kingPosition, chess.NewPosition(direction, 0)) {
		if piece != nil && piece != castlingRook {
			return fmt.Errorf("%w: an obstacle", Err)
		}
	}

	return nil
}

func fileDirection(castlingType move.Castling) chess.File {
	return map[move.Castling]chess.File{
		move.CastlingShort: 1,
		move.CastlingLong:  -1,
	}[castlingType]
}

func castlingRookPosition(castlingType move.Castling, rank chess.Rank) chess.Position {
	return map[move.Castling]chess.Position{
		move.CastlingShort: chess.NewPosition(chess.FileH, rank),
		move.CastlingLong:  chess.NewPosition(chess.FileA, rank),
	}[castlingType]
}
