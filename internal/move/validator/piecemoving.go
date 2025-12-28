// validator содержит валидаторы для проверки возможности выполнения ходов в шахматах.
package validator

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
)

var (
	Err              = errors.New("validation error")
	ErrEmptySquare   = fmt.Errorf("%w: no piece at square", Err)
	ErrInvalidString = fmt.Errorf("%w: invalid move string", Err)

	ErrMoving = fmt.Errorf("%w: moving validation error", Err)
)

func ValidatePieceMove(move move.Piece, movingPieceNotation string, board chess.Board) error {
	piece, err := board.Squares().FindByPosition(move.From)
	if err != nil || piece == nil {
		return ErrEmptySquare
	}

	if piece.Side() != board.Turn() {
		return fmt.Errorf("%w: wrong side", ErrMoving)
	}
	if piece.Notation() != movingPieceNotation {
		return fmt.Errorf("%w: moving piece doesn't found at this position", ErrMoving)
	}
	if !board.LegalMoves(piece).ContainsOne(move.To) {
		return fmt.Errorf("%w: piece doesn't have such move", ErrMoving)
	}

	return nil
}
