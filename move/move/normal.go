package move

import (
	"fmt"
	"regexp"

	"github.com/elaxer/chess/position"
	"github.com/elaxer/rgx"
	"github.com/elaxer/standardchess/piece"
	validation "github.com/go-ozzo/ozzo-validation"
)

var RegexpNormal = regexp.MustCompile("^(?P<piece>[KQBNR])?(?P<from>[a-p]?(1[0-6]|[1-9])?)x?(?P<to>[a-p](1[0-6]|[1-9]))[#+]?$")

// Normal представляет обычный ход фигурой в шахматах.
type Normal struct {
	Piece
	// PieceNotation обозначает фигуру, которая делает ход.
	PieceNotation string `json:"piece_notation"`
}

func NewNormal(from, to position.Position, pieceNotation string) *Normal {
	return &Normal{NewPiece(from, to), pieceNotation}
}

// NormalFromString создает новый ход из шахматной нотации.
func NormalFromString(notation string) (*Normal, error) {
	data, err := rgx.Group(RegexpNormal, notation)
	if err != nil {
		return nil, err
	}

	return NewNormal(
		position.FromString(data["from"]),
		position.FromString(data["to"]),
		data["piece"],
	), nil
}

func (m *Normal) Validate() error {
	pieceNotations := make([]any, 0, len(piece.AllNotations))
	for _, notation := range piece.AllNotations {
		pieceNotations = append(pieceNotations, notation)
	}

	return validation.ValidateStruct(
		m,
		validation.Field(&m.Piece),
		validation.Field(&m.PieceNotation, validation.In(pieceNotations...)),
	)
}

func (m *Normal) String() string {
	return fmt.Sprintf("%s%s%s", m.PieceNotation, m.From, m.To)
}
