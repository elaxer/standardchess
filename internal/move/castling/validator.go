package castling

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
)

var ErrValidation = errors.New("castling move validation error")

var (
	kingFileAfterShortCastling = chess.FileG
	rookFileAfterShortCastling = chess.FileF

	kingFileAfterLongCastling = chess.FileC
	rookFileAfterLongCastling = chess.FileD
)

func ValidateMove(castlingType CastlingType, board chess.Board) error {
	return validateMove(castlingType, board.Turn(), board, true)
}

func ValidateMoveWithObstacle(
	castlingType CastlingType,
	side chess.Color,
	board chess.Board,
) error {
	return validateMove(castlingType, side, board, false)
}

func validateMove(
	castlingType CastlingType,
	side chess.Color,
	board chess.Board,
	validateObstacle bool,
) error {
	king, kingPosition := board.Squares().FindPiece(piece.NotationKing, side)
	if err := validateKing(king, side, board); err != nil {
		return err
	}
	fileDir := fileDirection(castlingType)

	rook, _, hasObstacle, err := getRook(fileDir, side, board.Squares(), kingPosition)
	if err != nil {
		return err
	}
	if rook.IsMoved() {
		return fmt.Errorf("%w: the rook already has been moved", ErrValidation)
	}

	if validateObstacle && hasObstacle {
		return fmt.Errorf("%w: an obstacle", ErrValidation)
	}

	kingNewPosition, rookNewPosition := pickPositions(castlingType, kingPosition.Rank)

	if side == board.Turn() &&
		(board.IsSquareAttacked(kingNewPosition) || board.IsSquareAttacked(rookNewPosition)) {
		return fmt.Errorf("%w: castling squares are under threat", ErrValidation)

	}

	return nil
}

func validateKing(king chess.Piece,
	side chess.Color,
	board chess.Board) error {
	if king == nil {
		return fmt.Errorf("%w: the king wasn't found", ErrValidation)
	}
	if king.IsMoved() {
		return fmt.Errorf("%w: the king already has been moved", ErrValidation)
	}
	if side == board.Turn() && !board.State().Type().IsClear() {
		return fmt.Errorf("%w: the king is under threat", ErrValidation)
	}

	return nil
}

func getRook(
	fileDir chess.File,
	side chess.Color,
	squares *chess.Squares,
	kingPosition chess.Position,
) (chess.Piece, chess.Position, bool, error) {
	hasObstacle := false
	for position, p := range squares.IterByDirection(kingPosition, chess.NewPosition(fileDir, 0)) {
		if p == nil {
			continue
		}
		if p.Color() != side || p.Notation() != piece.NotationRook {
			hasObstacle = true

			continue
		}

		return p, position, hasObstacle, nil
	}

	return nil, chess.NewPositionEmpty(), hasObstacle, fmt.Errorf(
		"%w: rook wasn't found",
		ErrValidation,
	)
}

func fileDirection(castlingType CastlingType) chess.File {
	if castlingType.IsLong() {
		return -1
	}

	return 1
}
