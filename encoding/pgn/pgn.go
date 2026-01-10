// Package pgn provides functionality to encode/decode chess games in the Portable Game Notation (PGN) format.
// It includes encoding headers, moves, and results into a PGN string.
// It also provides a way to decode PGN strings into headers and moves.
package pgn

import "github.com/elaxer/chess"

func MoveHistoryToStrings(moveResults []chess.MoveResult) []string {
	strSlice := make([]string, 0, len(moveResults))
	for _, move := range moveResults {
		strSlice = append(strSlice, move.String())
	}

	return strSlice
}
