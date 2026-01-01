package promotion

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess"
	"github.com/elaxer/rgx"
	"github.com/elaxer/standardchess/internal/move/piecemove"
	"github.com/elaxer/standardchess/internal/piece"
	validation "github.com/go-ozzo/ozzo-validation"
)

var regexpPromotion = regexp.MustCompile("(?P<from>[a-p]?(1[0-6]|[1-9])?)x?(?P<to>[a-p](1[0-6]|[1-9]))=(?P<promoted_piece>[QBNR])[#+]?$")

// Move представляет ход с превращением пешки в другую фигуру.
// В шахматной нотации он записывается как "e8=Q" или "e7=R+".
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
	return validation.ValidateStruct(
		m,
		validation.Field(&m.PieceMove),
		validation.Field(
			&m.PromotedPieceNotation,
			validation.Required,
			validation.In(piece.NotationQueen, piece.NotationRook, piece.NotationBishop, piece.NotationKnight),
		),
	)
}

func (m *Move) String() string {
	return fmt.Sprintf("%s%s=%s", m.PieceMove.From, m.PieceMove.To, m.PromotedPieceNotation)
}
