package rule

import "github.com/elaxer/chess"

// Rule is a function that checks the state of the board
type Rule func(board chess.Board, side chess.Side) chess.State
