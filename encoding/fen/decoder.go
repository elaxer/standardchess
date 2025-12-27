package fen

import (
	"errors"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
	"github.com/elaxer/rgx"
	"github.com/elaxer/standardchess/piece"
)

// ErrDecoding is returned when there is an error decoding a FEN string.
var ErrDecoding = errors.New("error decoding FEN string")

var errSquaresNumOverflowed = fmt.Errorf(
	"%w: number of squares in one row is overflowed",
	ErrDecoding,
)

var regexpFEN = regexp.MustCompile(
	`(?i)^(?P<placement>(((1[0-6]|[1-9])|[PNBRQK])+/){5,15}((1[0-6]|[1-9])|[PNBRQK])+)\s?(?P<side>[WB])?`,
)

// Decoder decodes a FEN string into a chess board.
// It uses a board factory to create the board and a piece factory to create pieces.
type Decoder struct {
	boardFactory chess.BoardFactory
}

func NewDecoder(boardFactory chess.BoardFactory) *Decoder {
	return &Decoder{boardFactory}
}

// Decode decodes a FEN string into a chess board.
// The FEN string should match the regular expression defined in Regexp.
// It returns an error if the FEN string is invalid or if there is an error creating the board or pieces.
func (d *Decoder) Decode(fen string) (chess.Board, error) {
	if d.boardFactory == nil {
		return nil, fmt.Errorf("%w: board factory must be provided", ErrDecoding)
	}

	data, err := rgx.Group(regexpFEN, fen)
	if err != nil {
		return nil, err
	}

	placement := make(map[position.Position]chess.Piece, 256)

	rows := strings.Split(data["placement"], "/")
	if len(rows) > int(position.RankMax) {
		return nil, fmt.Errorf("%w: number of rows are overflowed", ErrDecoding)
	}
	slices.Reverse(rows)
	for i, row := range rows {
		//nolint:gosec
		rowPlacement, err := d.placementFromRow(row, position.Rank(i+1))
		if err != nil {
			return nil, err
		}

		maps.Copy(placement, rowPlacement)
	}

	return d.boardFactory.Create(d.side(data["side"]), placement)
}

func (d *Decoder) placementFromRow(
	row string,
	rank position.Rank,
) (map[position.Position]chess.Piece, error) {
	placement := make(map[position.Position]chess.Piece, 16)

	rowRunes := []rune(row)
	if len(rowRunes) > int(position.FileMax) {
		return nil, errSquaresNumOverflowed
	}

	pos := position.New(position.FileMin, rank)
	for i, char := range rowRunes {
		if i+1 < len(rowRunes) && d.isArabDigit(rowRunes[i]) && d.isArabDigit(rowRunes[i+1]) {
			emptySquaresLen, _ := strconv.Atoi(string(rowRunes[i : i+2]))
			//nolint:gosec
			pos.File += position.File(emptySquaresLen)

			continue
		}
		if d.isArabDigit(char) {
			pos.File += position.File(char - '0')

			continue
		}

		piece, err := d.createPiece(char)
		if err != nil {
			return nil, err
		}

		placement[pos] = piece
		pos.File++
	}

	if len(placement) > int(position.FileMax) {
		return nil, errSquaresNumOverflowed
	}

	return placement, nil
}

func (d *Decoder) isArabDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func (d *Decoder) side(str string) chess.Side {
	switch strings.ToLower(str) {
	case "w", "":
		return chess.SideWhite
	default:
		return chess.SideBlack
	}
}

func (d *Decoder) createPiece(char rune) (chess.Piece, error) {
	side := chess.SideWhite
	if char >= 'a' && char <= 'z' {
		side = chess.SideBlack
	}

	notation := strings.ToUpper(string(char))
	if notation == "P" {
		notation = piece.NotationPawn
	}

	return piece.New(notation, side)
}
