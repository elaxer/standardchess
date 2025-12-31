package move

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess"
	"github.com/elaxer/rgx"
	"github.com/elaxer/standardchess/internal/piece"
	validation "github.com/go-ozzo/ozzo-validation"
)

var regexpPromotion = regexp.MustCompile("(?P<from>[a-p]?(1[0-6]|[1-9])?)x?(?P<to>[a-p](1[0-6]|[1-9]))=(?P<promoted_piece>[QBNR])[#+]?$")

// Promotion представляет ход с превращением пешки в другую фигуру.
// В шахматной нотации он записывается как "e8=Q" или "e7=R+".
type Promotion struct {
	PieceMove

	PromotedPieceNotation string `json:"promoted_piece_notation"`
}

func NewPromotion(from, to chess.Position, promotedPieceNotation string) *Promotion {
	return &Promotion{
		NewPieceMove(from, to),
		promotedPieceNotation,
	}
}

func PromotionFromString(notation string) (*Promotion, error) {
	data, err := rgx.Group(regexpPromotion, notation)
	if err != nil {
		return nil, err
	}

	return NewPromotion(
		chess.PositionFromString(data["from"]),
		chess.PositionFromString(data["to"]),
		data["promoted_piece"],
	), nil
}

func (m *Promotion) Validate() error {
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

func (m *Promotion) String() string {
	return fmt.Sprintf("%s%s=%s", m.From, m.To, m.PromotedPieceNotation)
}
