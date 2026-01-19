// Package pgn provides functionality to encode/decode chess games in the Portable Game Notation (PGN) format.
// It includes encoding headers, moves, and results into a PGN string.
// It also provides a way to decode PGN strings into headers and moves.
package pgn

import (
	"errors"
	"regexp"
	"strings"

	"github.com/elaxer/chess"
	"github.com/elaxer/rgx"
)

var ErrDecode = errors.New("error decoding PGN string")

var (
	regexpMoves = regexp.MustCompile(
		`(([NBKRQ]?[a-h]?[1-8]?x?[a-h][1-8](?:=[NBRQ])?)|([0Oo]-[0Oo](-[0Oo])?))(\+|\#)?`,
	)
	regexpHeaders = regexp.MustCompile(`\[(?P<name>[\w]+)\s+"(?P<value>[^"]*)"\]`)
	regexpResult  = regexp.MustCompile(`((1-0)|(0-1)|(1/2-1/2)|\*)\z`)
)

type PGN struct {
	headers []Header
	moves   []chess.Move
	result  Result
}

func NewPGN(headers []Header, moves []chess.Move, result Result) PGN {
	return PGN{headers, moves, result}
}

func (p PGN) Headers() []Header {
	return p.headers
}

func (p PGN) Moves() []chess.Move {
	return p.moves
}

func (p PGN) Result() Result {
	return p.result
}

func (p PGN) String() string {
	var pgnStr strings.Builder
	pgnStr.WriteString(encodeHeaders(p.headers) + "\n\n")

	movesStr := wrapText(encodeMoves(p.moves), 79)
	pgnStr.WriteString(movesStr)

	return pgnStr.String() + " " + string(p.result)
}

// FromString decodes a PGN string into headers, moves and result.
// It returns a slice of Header structs and a slice of chess.Move structs.
// If there is an error during decoding, it returns an error.
// The PGN string should match the regular expressions defined in headersRegexp and movesRegexp.
func FromString(pgnStr string) (PGN, error) {
	moves, err := decodeMoves(pgnStr)
	if err != nil {
		return PGN{}, err
	}
	result, err := decodeResult(pgnStr)
	if err != nil {
		return PGN{}, err
	}

	return PGN{decodeHeaders(pgnStr), moves, result}, nil
}

func decodeHeaders(pgnStr string) []Header {
	headers := make([]Header, 0)

	data, err := rgx.Groups(regexpHeaders, pgnStr)
	if err != nil {
		return headers
	}

	for _, match := range data {
		headers = append(headers, NewHeader(match["name"], match["value"]))
	}

	return headers
}

func decodeMoves(pgnStr string) ([]chess.Move, error) {
	moves := make([]chess.Move, 0, 100)
	data := regexpMoves.FindAllString(pgnStr, -1)
	if len(data) == 0 {
		return nil, ErrDecode
	}

	for _, move := range data {
		moves = append(moves, chess.StringMove(move))
	}

	return moves, nil
}

func decodeResult(pgnStr string) (Result, error) {
	result := Result(regexpResult.FindString(pgnStr))
	if result == "" {
		return "", ErrDecode
	}

	return result, nil
}
