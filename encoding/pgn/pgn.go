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
	regexpSplit = regexp.MustCompile(`\s\s`)
	regexpMove  = regexp.MustCompile(
		`(([NBKRQ]?[a-h]?[1-8]?x?[a-h][1-8](?:=[NBRQ])?)|([0Oo]-[0Oo](-[0Oo])?))(\+|\#)?`,
	)
	regexpHeader = regexp.MustCompile(`\[(?P<name>[\w]+)\s+"(?P<value>[^"]*)"\]`)
	regexpResult = regexp.MustCompile(`((1-0)|(0-1)|(1/2-1/2)|\*)\z`)
)

// PGN represents a single chess game in PGN format.
// It contains headers, moves, and the result of the game.
type PGN struct {
	headers []Header
	moves   []chess.Move
	result  Result
}

func NewPGN(headers []Header, moves []chess.Move, result Result) PGN {
	return PGN{headers, moves, result}
}

// Headers returns the list of headers for the PGN game.
func (p PGN) Headers() []Header {
	return p.headers
}

// Moves returns the list of moves in the PGN game.
func (p PGN) Moves() []chess.Move {
	return p.moves
}

// Result returns the result of the PGN game.
func (p PGN) Result() Result {
	return p.result
}

// Format returns the PGN as a formatted string,
// wrapping move text at the specified width.
// movesWidth specifies the maximum line length for the moves section.
// Headers are included at the top, followed by the moves, then the result.
func (p PGN) Format(movesWidth int) string {
	var pgnStr strings.Builder
	pgnStr.WriteString(encodeHeaders(p.headers) + "\n\n")

	movesStr := wrapText(encodeMoves(p.moves), movesWidth)
	pgnStr.WriteString(movesStr)

	return strings.TrimSpace(pgnStr.String() + " " + string(p.result))
}

func (p PGN) String() string {
	return p.Format(79)
}

// FromString parses a single PGN game from the provided string.
// pgnStr should contain headers, moves and result.
// Headers can be omitted.
// Returns a PGN object containing headers, moves, and the result.
// Returns ErrDecode if the string does not match the expected PGN format.
func FromString(pgnStr string) (PGN, error) {
	s := regexpSplit.Split(strings.TrimSpace(pgnStr), 2)
	headerStr := ""
	movesStr := ""
	if len(s) == 1 {
		movesStr = s[0]
	} else {
		headerStr = s[0]
		movesStr = s[1]
	}

	moves, err := decodeMoves(movesStr)
	if err != nil {
		return PGN{}, err
	}
	result, err := decodeResult(movesStr)
	if err != nil {
		return PGN{}, err
	}

	return PGN{decodeHeaders(headerStr), moves, result}, nil
}

func decodeHeaders(pgnStr string) []Header {
	headers := make([]Header, 0)

	data, err := rgx.Groups(regexpHeader, pgnStr)
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
	data := regexpMove.FindAllString(pgnStr, -1)
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
