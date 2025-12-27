// Package pgn provides functionality to encode/decode chess games in the Portable Game Notation (PGN) format.
// It includes encoding headers, moves, and results into a PGN string.
// It also provides a way to decode PGN strings into headers and moves.
package pgn

import (
	"fmt"
	"strings"

	"github.com/elaxer/chess"
)

// Encode encodes the given headers and board state into a PGN string.
// It formats the headers, moves, and result according to the PGN standard.
// The headers are encoded as a string with each header on a new line.
// The moves are encoded as a single line with move numbers and piece notations.
// The result is appended at the end of the PGN string.
func Encode(headers []Header, board chess.Board, result string) string {
	var pgn strings.Builder
	fmt.Fprintln(&pgn, EncodeHeaders(headers))
	fmt.Fprintln(&pgn)

	movesStr := wrapText(EncodeMoves(board.MoveHistory()), 79)
	fmt.Fprint(&pgn, movesStr)

	return pgn.String() + " " + result
}

// EncodeHeaders encodes the given headers into a PGN string.
// Each header is formatted as "[name "value"]" and joined by newlines.
func EncodeHeaders(headers []Header) string {
	headerStrings := make([]string, 0, len(headers))
	for _, header := range headers {
		headerStrings = append(headerStrings, header.String())
	}

	return strings.Join(headerStrings, "\n")
}

// EncodeMoves encodes the given moves into a PGN string.
// It formats the moves with move numbers and piece notations.
// Each move is separated by a space, and move numbers are added every two moves.
// The first move is considered move 1, and the second move is considered move 1 as well (for white and black).
// The move number is incremented for every two moves (one for white and one for black).
// The moves are formatted as "1. e4 1... e5 2. Nf3 2... Nc6" for example.
func EncodeMoves(moves []chess.MoveResult) string {
	var str strings.Builder
	currentMoveNumber := 0
	for i, move := range moves {
		if moveNumber := (i + 2) / 2; moveNumber != currentMoveNumber {
			currentMoveNumber = moveNumber
			fmt.Fprintf(&str, "%d. ", currentMoveNumber)
		}

		fmt.Fprintf(&str, "%s ", move)
	}

	return str.String()
}

func wrapText(text string, maxWidth int) string {
	if maxWidth <= 0 {
		return text
	}

	var result strings.Builder
	var lineLen int

	words := strings.Fields(text)

	for i, word := range words {
		if lineLen+len(word) > maxWidth {
			result.WriteString("\n")
			lineLen = 0
		} else if i != 0 {
			result.WriteString(" ")
			lineLen++
		}

		result.WriteString(word)
		lineLen += len(word)
	}

	return result.String()
}
