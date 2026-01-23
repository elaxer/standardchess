package pgn

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/state"
)

const (
	ResultInProcess Result = "*"
	ResultWinWhite  Result = "1-0"
	ResultWinBlack  Result = "0-1"
	ResultDraw      Result = "1/2-1/2"
)

// Result represents the outcome of a chess game in PGN notation.
type Result string

// ResultFromBoard determines the Result of a game given a chess.Board state.
//
// It returns:
//   - ResultWinWhite if White has won
//   - ResultWinBlack if Black has won
//   - ResultDraw if the game is drawn
//   - ResultInProcess if the game is still ongoing
func ResultFromBoard(board chess.Board) Result {
	if !board.State().Type().IsTerminal() {
		return ResultInProcess
	}
	if board.State() == state.Checkmate {
		if board.Turn().IsWhite() {
			return ResultWinBlack
		}

		return ResultWinWhite
	}

	return ResultDraw
}

func (r Result) IsInProcess() bool {
	return r == ResultInProcess
}

func (r Result) IsWinWhite() bool {
	return r == ResultWinWhite
}

func (r Result) IsWinBlack() bool {
	return r == ResultWinBlack
}

func (r Result) IsDraw() bool {
	return r == ResultDraw
}
