// Package promotion contains code for validating,
// executing, and canceling pawn moves
// to the farthest horizontal row of the board with promotion.
package promotion

import (
	"errors"
	"fmt"
	"regexp"
	"slices"

	"github.com/elaxer/chess"
	"github.com/elaxer/rgx"
	"github.com/elaxer/standardchess/internal/move/piecemove"
	"github.com/elaxer/standardchess/internal/piece"
)

var ErrMoveValidation = errors.New("promotion move validation error")

var regexpPromotion = regexp.MustCompile(
	"(?P<from>[a-p]?(1[0-6]|[1-9])?)x?(?P<to>[a-p](1[0-6]|[1-9]))=(?P<promoted_piece>[QBNR])[#+]?$",
)

var allowedNotations = []string{
	piece.NotationQueen,
	piece.NotationRook,
	piece.NotationBishop,
	piece.NotationKnight,
}

type Move struct {
	piecemove.PieceMove

	PromotedPieceNotation string `json:"promoted_piece_notation"`
}

func NewMove(from, to chess.Position, promotedPieceNotation string) *Move {
	return &Move{
		piecemove.NewPieceMove(from, to),
		promotedPieceNotation,
	}
}

func MoveFromString(notation string) (*Move, error) {
	data, err := rgx.Group(regexpPromotion, notation)
	if err != nil {
		return nil, err
	}

	return NewMove(
		chess.PositionFromString(data["from"]),
		chess.PositionFromString(data["to"]),
		data["promoted_piece"],
	), nil
}

func (m *Move) Validate() error {
	if err := m.PieceMove.Validate(); err != nil {
		return err
	}
	if !slices.Contains(allowedNotations, m.PromotedPieceNotation) {
		return fmt.Errorf("%w: wrong new promoted piece notation", ErrMoveValidation)
	}

	return nil
}

func (m *Move) String() string {
	return m.From.String() + m.To.String() + "=" + m.PromotedPieceNotation
}
