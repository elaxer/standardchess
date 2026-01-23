package pgn

import (
	"bufio"
	"io"
	"strings"
)

// Parse reads PGN data from the provided io.Reader and returns a slice of PGN games.
// The reader can be any type implementing io.Reader, such as a file, buffer, or network stream.
// The function splits the input into individual PGN games using empty lines as separators.
// Each game is returned as a PGN object containing headers, moves, and the result.
// Returns a slice of PGN games and an error if reading or parsing fails.
func Parse(r io.Reader) ([]PGN, error) {
	scanner := bufio.NewScanner(r)
	pgns := make([]PGN, 0)
	var pgnStr strings.Builder

	var lineBreaks uint8
	wasLineBreak := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "\n" {
			if !wasLineBreak {
				lineBreaks++
			}
			wasLineBreak = true
		} else {
			wasLineBreak = false
		}

		if lineBreaks == 2 {
			pgn, err := FromString(pgnStr.String())
			if err != nil {
				return nil, err
			}

			pgns = append(pgns, pgn)

			pgnStr.Reset()

			continue
		}

		if _, err := pgnStr.WriteString(line); err != nil {
			return nil, err
		}
	}

	return pgns, nil
}
