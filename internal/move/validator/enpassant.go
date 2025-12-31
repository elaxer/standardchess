package validator

import (
	"fmt"
	"math"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/result"
	"github.com/elaxer/standardchess/internal/piece"
)

var ErrEnPassant = fmt.Errorf("%w: en passant move validation error", Err)

func ValidateEnPassant(from, to chess.Position, board chess.Board) error {
	movesCount := len(board.MoveHistory())
	if movesCount == 0 {
		return fmt.Errorf("%w: no moves made on the board", ErrEnPassant)
	}

	lastMove := board.MoveHistory()[movesCount-1]
	normalMove, ok := lastMove.(*result.Normal)
	if !ok {
		return fmt.Errorf("%w: the last move is not a normal move", ErrEnPassant)
	}
	if to != chess.NewPosition(normalMove.InputMove.To.File, normalMove.InputMove.To.Rank+piece.PawnRankDirection(board.Turn())) {
		return fmt.Errorf("%w: wrong move coordinates", ErrEnPassant)
	}
	if normalMove.InputMove.PieceNotation != piece.NotationPawn {
		return fmt.Errorf("%w: the last move was not made by a pawn", ErrEnPassant)
	}
	if normalMove.FromFull.Rank-normalMove.InputMove.To.Rank != chess.Rank(math.Abs(2)) {
		return fmt.Errorf("%w: the pawn moved only one square", ErrEnPassant)
	}

	p, err := board.Squares().FindByPosition(from)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrEnPassant, err)
	}
	if p == nil || p.Notation() != piece.NotationPawn {
		return fmt.Errorf("%w: the moving pawn wasn't found", ErrEnPassant)
	}

	if from.Rank != enPassantRank(board.Turn()) {
		return fmt.Errorf("%w: the pawn is in the wrong rank", ErrEnPassant)
	}
	if from.File-normalMove.InputMove.To.File != chess.File(math.Abs(1)) {
		return fmt.Errorf("%w: the pawn is in the wrong file", ErrEnPassant)
	}

	return nil
}

func enPassantRank(side chess.Side) chess.Rank {
	if side.IsBlack() {
		return chess.Rank5
	}

	return chess.Rank4
}
