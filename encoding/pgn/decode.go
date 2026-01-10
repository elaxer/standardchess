package pgn

import (
	"regexp"

	"github.com/elaxer/chess"
	"github.com/elaxer/rgx"
)

var (
	regexpMoves = regexp.MustCompile(
		"([NBKRQ]?[a-h]?[1-8]?x?[a-h][1-8](?:=[NBRQ])?)|([0Oo]-[0Oo](-[0Oo])?)",
	)
	regexpHeaders = regexp.MustCompile(`\[(?P<name>[\w]+)\s+"(?P<value>[^"]*)"\]`)
)

// Decode decodes a PGN string into headers and moves.
// It returns a slice of Header structs and a slice of chess.Move structs.
// If there is an error during decoding, it returns an error.
// The PGN string should match the regular expressions defined in headersRegexp and movesRegexp.
func Decode(pgn string) ([]Header, []chess.Move, error) {
	headers, _ := decodeHeaders(pgn)

	return headers, decodeMoves(pgn), nil
}

func decodeHeaders(pgn string) ([]Header, error) {
	headers := make([]Header, 0)

	data, err := rgx.Groups(regexpHeaders, pgn)
	if err != nil {
		return nil, err
	}

	for _, match := range data {
		headers = append(headers, NewHeader(match["name"], match["value"]))
	}

	return headers, nil
}

func decodeMoves(pgn string) []chess.Move {
	moves := make([]chess.Move, 0, 100)
	data := regexpMoves.FindAllString(pgn, -1)

	for _, move := range data {
		moves = append(moves, chess.StringMove(move))
	}

	return moves
}
