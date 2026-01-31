package pgn

import (
	"fmt"
	"strings"

	"github.com/elaxer/chess"
)

func Encode(headers Headers, board chess.Board, result Result) PGN {
	moves := make([]string, 0, len(board.MoveHistory()))
	for _, move := range board.MoveHistory() {
		moves = append(moves, move.String())
	}

	return NewPGN(headers, moves, result)
}

func encodeHeaders(headers Headers) string {
	headerStrings := make([]string, 0, len(headers))
	for _, header := range headers {
		headerStrings = append(headerStrings, header.String())
	}

	return strings.Join(headerStrings, "\n")
}

func encodeMoves(moves []string) string {
	var str strings.Builder
	currentMoveNumber := 0
	for i, move := range moves {
		if moveNumber := (i + 2) / 2; moveNumber != currentMoveNumber {
			currentMoveNumber = moveNumber
			fmt.Fprintf(&str, "%d. ", currentMoveNumber)
		}

		if i == len(moves)-1 {
			str.WriteString(move)

			break
		}

		str.WriteString(move + " ")
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
			result.WriteRune(' ')
			lineLen++
		}

		result.WriteString(word)
		lineLen += len(word)
	}

	return result.String()
}
