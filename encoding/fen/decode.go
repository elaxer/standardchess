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
	"github.com/elaxer/rgx"
	board "github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/internal/piece"
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

// Decode decodes a FEN string into a chess board.
// The FEN string should match the regular expression defined in Regexp.
// It returns an error if the FEN string is invalid or if there is an error creating the board or pieces.
func Decode(fen string) (chess.Board, error) {
	data, err := rgx.Group(regexpFEN, fen)
	if err != nil {
		return nil, err
	}

	placement := make(map[chess.Position]chess.Piece, 256)

	rows := strings.Split(data["placement"], "/")
	if len(rows) > int(chess.RankMax) {
		return nil, fmt.Errorf("%w: number of rows are overflowed", ErrDecoding)
	}

	var maxFile chess.File = 1
	slices.Reverse(rows)
	for i, row := range rows {
		//nolint:gosec
		rowPlacement, filesNum, err := placementFromRow(row, chess.Rank(i+1))
		if err != nil {
			return nil, err
		}

		maxFile = max(maxFile, filesNum)

		maps.Copy(placement, rowPlacement)
	}

	return board.NewBoardSized(
		side(data["side"]),
		placement,
		chess.NewPosition(maxFile-1, chess.Rank(len(rows))),
	)
}

func placementFromRow(
	row string,
	rank chess.Rank,
) (map[chess.Position]chess.Piece, chess.File, error) {
	placement := make(map[chess.Position]chess.Piece, int(chess.FileMax))

	rowRunes := []rune(row)
	if len(rowRunes) > int(chess.FileMax) {
		return nil, 0, errSquaresNumOverflowed
	}

	pos := chess.NewPosition(chess.FileMin, rank)
	for i := 0; i < len(rowRunes); i++ {
		char := rowRunes[i]
		if i+1 < len(rowRunes) && isArabDigit(rowRunes[i]) && isArabDigit(rowRunes[i+1]) {
			emptySquaresLen, _ := strconv.Atoi(string(rowRunes[i : i+2]))
			//nolint:gosec
			pos.File += chess.File(emptySquaresLen)

			i++
			continue
		}
		if isArabDigit(char) {
			pos.File += chess.File(char - '0')

			continue
		}

		piece, err := createPiece(char)
		if err != nil {
			return nil, 0, err
		}

		placement[pos] = piece
		pos.File++
	}

	if len(placement) > int(chess.FileMax) {
		return nil, 0, errSquaresNumOverflowed
	}

	return placement, pos.File, nil
}

func isArabDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func side(str string) chess.Side {
	switch strings.ToLower(str) {
	case "w", "":
		return chess.SideWhite
	default:
		return chess.SideBlack
	}
}

func createPiece(char rune) (chess.Piece, error) {
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
