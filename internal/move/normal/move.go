// Package normal contains the structure of the move, the result of the move,
// and the logic for executing and cancelling a standard move on a chessboard.
package normal

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

var ErrMoveValidation = errors.New("normal move validation error")

var regexpNormal = regexp.MustCompile(
	"^(?P<piece>[KQBNR])?(?P<from>[a-p]?(1[0-6]|[1-9])?)x?(?P<to>[a-p](1[0-6]|[1-9]))[#+]?$",
)

type Move struct {
	piecemove.PieceMove

	PieceNotation string `json:"piece_notation"`
}

func NewMove(from, to chess.Position, pieceNotation string) *Move {
	return &Move{piecemove.NewPieceMove(from, to), pieceNotation}
}

func MoveFromNotation(notation string) (*Move, error) {
	data, err := rgx.Group(regexpNormal, notation)
	if err != nil {
		return nil, err
	}

	return NewMove(
		chess.PositionFromString(data["from"]),
		chess.PositionFromString(data["to"]),
		data["piece"],
	), nil
}

func MoveFromUCI(uciStr string, squares *chess.Squares) (*Move, error) {
	uciMove, err := uci.MoveFromString(uciStr)
	if err != nil {
		return nil, err
	}

	piece, err := squares.FindByPosition(uciMove.From)
	if err != nil {
		return nil, err
	}
	if piece == nil {
		return nil, fmt.Errorf("%w: no piece at this from position", ErrMoveValidation)
	}

	return NewMove(uciMove.From, uciMove.To, piece.Notation()), nil
}

func (m *Move) String() string {
	return m.PieceNotation + m.From.String() + m.To.String()
}

func (m *Move) validate() error {
	if err := m.Validate(); err != nil {
		return err
	}
	if !slices.Contains(piece.AllNotations, m.PieceNotation) {
		return fmt.Errorf("%w: wrong piece notation", ErrMoveValidation)
	}

	return nil
}
