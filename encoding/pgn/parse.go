package pgn

import (
	"bufio"
	"io"
	"iter"
	"strings"
)

// Parse reads PGN data from the provided io.Reader and returns a slice of PGN games.
// The reader can be any type implementing io.Reader, such as a file, buffer, or network stream.
// The function splits the input into individual PGN games using empty lines as separators.
// Each game is returned as a PGN object containing headers, moves, and the result.
// Returns a slice of PGN games and an error if reading or parsing fails.
func Parse(r io.Reader) iter.Seq2[PGN, error] {
	scanner := bufio.NewScanner(r)
	var pgnStr strings.Builder

	wasLineBreak := false

	return func(yield func(PGN, error) bool) {
		for scanner.Scan() {
			line := scanner.Text()

			if strings.TrimSpace(line) != "" {
				pgnStr.WriteString(line + "\n")

				continue
			}

			if pgnStr.Len() == 0 {
				continue
			}

			if !wasLineBreak {
				wasLineBreak = true
				pgnStr.WriteString("\n")

				continue
			}

			pgn, err := FromString(pgnStr.String())
			if !yield(pgn, err) {
				return
			}

			wasLineBreak = false
			pgnStr.Reset()
		}
	}
}
