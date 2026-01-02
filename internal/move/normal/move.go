package normal

import (
	"regexp"

	"github.com/elaxer/chess"
	"github.com/elaxer/rgx"
	"github.com/elaxer/standardchess/internal/move/piecemove"
	"github.com/elaxer/standardchess/internal/piece"
	validation "github.com/go-ozzo/ozzo-validation"
)

var regexpNormal = regexp.MustCompile("^(?P<piece>[KQBNR])?(?P<from>[a-p]?(1[0-6]|[1-9])?)x?(?P<to>[a-p](1[0-6]|[1-9]))[#+]?$")

// Move представляет обычный ход фигурой в шахматах.
type Move struct {
	piecemove.PieceMove

	// PieceNotation обозначает фигуру, которая делает ход.
	PieceNotation string `json:"piece_notation"`
}

func NewMove(from, to chess.Position, pieceNotation string) *Move {
	return &Move{piecemove.NewPieceMove(from, to), pieceNotation}
}

// MoveFromString создает новый ход из шахматной нотации.
func MoveFromString(notation string) (*Move, error) {
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

func (m *Move) Validate() error {
	pieceNotations := make([]any, 0, len(piece.AllNotations))
	for _, notation := range piece.AllNotations {
		pieceNotations = append(pieceNotations, notation)
	}

	return validation.ValidateStruct(
		m,
		validation.Field(&m.PieceMove),
		validation.Field(&m.PieceNotation, validation.In(pieceNotations...)),
	)
}

func (m *Move) String() string {
	return m.PieceNotation + m.From.String() + m.To.String()
}
