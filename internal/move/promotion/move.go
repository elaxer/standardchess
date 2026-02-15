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
	"github.com/elaxer/standardchess/internal/uci"
)

var (
	ErrMoveValidation = errors.New("promotion move validation error")

	errMoveValidationWrongPromoted = fmt.Errorf("%w: wrong promoted piece notation", ErrMoveValidation)
)

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

func MoveFromNotation(notation string) (*Move, error) {
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

func MoveFromUCI(uciStr string) (*Move, error) {
	uciMove, err := uci.MoveFromString(uciStr)
	if err != nil {
		return nil, err
	}
	if !slices.Contains(allowedNotations, uciMove.PromotedPieceNotation) {
		return nil, errMoveValidationWrongPromoted
	}

	return NewMove(uciMove.From, uciMove.To, uciMove.PromotedPieceNotation), nil
}

func (m *Move) String() string {
	return m.From.String() + m.To.String() + "=" + m.PromotedPieceNotation
}

func (m *Move) validate() error {
	if err := m.Validate(); err != nil {
		return err
	}
	if !slices.Contains(allowedNotations, m.PromotedPieceNotation) {
		return errMoveValidationWrongPromoted
	}

	return nil
}
