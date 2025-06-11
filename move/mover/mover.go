package mover

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	mv "github.com/elaxer/standardchess/move/move"
)

var Err = errors.New("mover error")

var (
	normalMover    = new(Normal)
	castlingMover  = new(Castling)
	promotionMover = new(Promotion)
)

func MakeMove(move chess.Move, board chess.Board) (chess.MoveResult, error) {
	return makeMoveFromString(move.String(), board)
}

func makeMoveFromString(string string, board chess.Board) (chess.MoveResult, error) {
	if move, err := mv.NormalFromString(string); err == nil {
		return normalMover.Make(move, board)
	}
	if move, err := mv.PromotionFromString(string); err == nil {
		return promotionMover.Make(move, board)
	}
	if move, err := mv.CastlingFromString(string); err == nil {
		return castlingMover.Make(move, board)
	}

	return nil, fmt.Errorf("%w: invalid move \"%s\"", Err, string)
}
