// Package rule contains functions that check boards for compliance with states.
package rule

import "github.com/elaxer/chess"

// Rule is a function that checks the state of the board.
type Rule func(board chess.Board) chess.State
