// Package fen contains functions for encoding the board into FEN encoding and vice versa.
package fen

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/rgx"
	"github.com/elaxer/standardchess/internal/move/castling"
)

var regexpFEN = regexp.MustCompile(
	`^(?P<placement>(((1[0-6]|[1-9])|[PNBRQKpnbrqk])+/){5,15}((1[0-6]|[1-9])|[PNBRQKpnbrqk])+)\s` +
		`(?P<turn>[wb])\s(?P<castlings>-|(K?Q?k?q?))\s(?P<enpassant>-|([a-p](1[0-6]|[1-9])))\s` +
		`(?P<halfmove_clock>\d+)\s(?P<move_number>\d+)$`,
)

type FEN struct {
	placement       string
	turn            chess.Color
	castlings       map[chess.Color]map[castling.CastlingType]bool
	enPassantSquare chess.Position
	halfmoveClock   int
	moveNumber      int
}

func FromString(str string) (FEN, error) {
	data, err := rgx.Group(regexpFEN, str)
	if err != nil {
		return FEN{}, err
	}

	halfmoveClock, err := strconv.Atoi(data["halfmove_clock"])
	if err != nil {
		panic(err)
	}
	moveNumber, err := strconv.Atoi(data["move_number"])
	if err != nil {
		panic(err)
	}

	return FEN{
		placement: data["placement"],
		turn:      color(data["turn"]),
		castlings: map[chess.Color]map[castling.CastlingType]bool{
			chess.ColorWhite: {
				castling.TypeShort: strings.Contains(data["castlings"], "K"),
				castling.TypeLong:  strings.Contains(data["castlings"], "Q"),
			},
			chess.ColorBlack: {
				castling.TypeShort: strings.Contains(data["castlings"], "k"),
				castling.TypeLong:  strings.Contains(data["castlings"], "q"),
			},
		},
		enPassantSquare: chess.PositionFromString(data["enpassant"]),
		halfmoveClock:   halfmoveClock,
		moveNumber:      moveNumber,
	}, nil
}

func (f FEN) Placement() string {
	return f.placement
}

func (f FEN) Turn() chess.Color {
	return f.turn
}

func (f FEN) Castlings(side chess.Color) (short, long bool) {
	return f.castlings[side][castling.TypeShort], f.castlings[side][castling.TypeLong]
}

func (f FEN) EnPassantSquare() chess.Position {
	return f.enPassantSquare
}

func (f FEN) HalfmoveClock() int {
	return f.halfmoveClock
}

func (f FEN) MoveNumber() int {
	return f.moveNumber
}

func (f FEN) String() string {
	castlings := ""
	if f.castlings[chess.ColorWhite][castling.TypeShort] {
		castlings += "K"
	}
	if f.castlings[chess.ColorWhite][castling.TypeLong] {
		castlings += "Q"
	}
	if f.castlings[chess.ColorBlack][castling.TypeShort] {
		castlings += "k"
	}
	if f.castlings[chess.ColorBlack][castling.TypeLong] {
		castlings += "q"
	}
	if castlings == "" {
		castlings = "-"
	}

	enPassantSquare := "-"
	if f.enPassantSquare.IsFull() {
		enPassantSquare = f.EnPassantSquare().String()
	}

	str := f.placement + " " + f.turn.String() + " " + castlings + " " + enPassantSquare + " "

	return str + strconv.Itoa(f.halfmoveClock) + " " + strconv.Itoa(f.moveNumber)
}

func color(str string) chess.Color {
	switch strings.ToLower(str) {
	case "w", "":
		return chess.ColorWhite
	default:
		return chess.ColorBlack
	}
}
