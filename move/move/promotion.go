package move

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess/position"
	"github.com/elaxer/rgx"
	"github.com/elaxer/standardchess/piece"
	validation "github.com/go-ozzo/ozzo-validation"
)

var RegexpPromotion = regexp.MustCompile("(?P<from>[a-p]?(1[0-6]|[1-9])?)x?(?P<to>[a-p](1[0-6]|[1-9]))=(?P<promoted_piece>[QBNR])[#+]?$")

// Promotion представляет ход с превращением пешки в другую фигуру.
// В шахматной нотации он записывается как "e8=Q" или "e7=R+".
type Promotion struct {
	Piece
	PromotedPieceNotation string `json:"promoted_piece_notation"`
}

func NewPromotion(from, to position.Position, promotedPieceNotation string) *Promotion {
	return &Promotion{
		NewPiece(from, to),
		promotedPieceNotation,
	}
}

func PromotionFromString(notation string) (*Promotion, error) {
	data, err := rgx.Group(RegexpPromotion, notation)
	if err != nil {
		return nil, err
	}

	return NewPromotion(
		position.FromString(data["from"]),
		position.FromString(data["to"]),
		data["promoted_piece"],
	), nil
}

func (m *Promotion) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Piece),
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
