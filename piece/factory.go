package piece

import (
	"errors"
	"strings"
	"unicode"

	"github.com/elaxer/chess"
)

var AllNotations = []string{
	NotationPawn,
	NotationRook,
	NotationKnight,
	NotationBishop,
	NotationQueen,
	NotationKing,
}

type factory struct{}

func NewFactory() *factory {
	return &factory{}
}

func (f *factory) CreateFromNotation(notation string, side chess.Side) (chess.Piece, error) {
	switch notation {
	case NotationPawn:
		return NewPawn(side), nil
	case NotationRook:
		return NewRook(side), nil
	case NotationKnight:
		return NewKnight(side), nil
	case NotationBishop:
		return NewBishop(side), nil
	case NotationQueen:
		return NewQueen(side), nil
	case NotationKing:
		return NewKing(side), nil
	}

	return nil, errors.New("todo")
}

func (f *factory) CreateFromString(str string) (chess.Piece, error) {
	if len(str) != 1 {
		return nil, errors.New("todo")
	}

	side := chess.SideWhite
	if unicode.IsLower([]rune(str)[0]) {
		side = chess.SideBlack
	}

	notation := strings.ToUpper(str)
	if notation == "P" {
		notation = NotationPawn
	}

	return f.CreateFromNotation(notation, side)
}
