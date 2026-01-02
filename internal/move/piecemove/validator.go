// validator содержит валидаторы для проверки возможности выполнения ходов в шахматах.
package piecemove

import (
	"errors"
	"fmt"
	"slices"

	"github.com/elaxer/chess"
)

var ErrValidation = errors.New("piece move validation error")

func ValidateMove(move PieceMove, movingPieceNotation string, board chess.Board) error {
	piece, err := board.Squares().FindByPosition(move.From)
	if err != nil || piece == nil {
		return fmt.Errorf("%w: no piece at square", ErrValidation)
	}

	if piece.Side() != board.Turn() {
		return fmt.Errorf("%w: wrong side", ErrValidation)
	}
	if piece.Notation() != movingPieceNotation {
		return fmt.Errorf("%w: moving piece doesn't found at this position", ErrValidation)
	}
	if !slices.Contains(board.LegalMoves(piece), move.To) {
		return fmt.Errorf("%w: piece doesn't have such move", ErrValidation)
	}

	return nil
}
