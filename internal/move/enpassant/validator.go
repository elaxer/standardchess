package enpassant

import (
	"errors"
	"fmt"
	"math"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/normal"
	"github.com/elaxer/standardchess/internal/piece"
)

var ErrValidation = errors.New("en passant move validation error")

func ValidateMove(from, to chess.Position, board chess.Board) error {
	movesCount := len(board.MoveHistory())
	if movesCount == 0 {
		return fmt.Errorf("%w: no moves made on the board", ErrValidation)
	}

	lastMove := board.MoveHistory()[movesCount-1]
	normalMove, ok := lastMove.(*normal.MoveResult)
	if !ok {
		return fmt.Errorf("%w: the last move is not a normal move", ErrValidation)
	}
	if to != chess.NewPosition(normalMove.InputMove.To.File, normalMove.InputMove.To.Rank+piece.PawnRankDirection(board.Turn())) {
		return fmt.Errorf("%w: wrong move coordinates", ErrValidation)
	}
	if normalMove.InputMove.PieceNotation != piece.NotationPawn {
		return fmt.Errorf("%w: the last move was not made by a pawn", ErrValidation)
	}
	if normalMove.FromFull.Rank-normalMove.InputMove.To.Rank != chess.Rank(math.Abs(2)) {
		return fmt.Errorf("%w: the pawn moved only one square", ErrValidation)
	}

	p, err := board.Squares().FindByPosition(from)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrValidation, err)
	}
	if p == nil || p.Notation() != piece.NotationPawn {
		return fmt.Errorf("%w: the moving pawn wasn't found", ErrValidation)
	}

	if from.Rank != enPassantRank(board.Turn()) {
		return fmt.Errorf("%w: the pawn is in the wrong rank", ErrValidation)
	}
	if from.File-normalMove.InputMove.To.File != chess.File(math.Abs(1)) {
		return fmt.Errorf("%w: the pawn is in the wrong file", ErrValidation)
	}

	return nil
}
