package mover

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	mv "github.com/elaxer/standardchess/internal/move/move"
)

var Err = errors.New("mover error")

func MakeMove(move chess.Move, board chess.Board) (chess.MoveResult, error) {
	return makeMoveFromString(move.String(), board)
}

func makeMoveFromString(string string, board chess.Board) (chess.MoveResult, error) {
	if move, err := mv.NormalFromString(string); err == nil {
		return MakeNormal(move, board)
	}
	if move, err := mv.PromotionFromString(string); err == nil {
		return MakePromotion(move, board)
	}
	if move, err := mv.CastlingFromString(string); err == nil {
		return MakeCastling(move, board)
	}

	return nil, fmt.Errorf("%w: invalid move \"%s\"", Err, string)
}
