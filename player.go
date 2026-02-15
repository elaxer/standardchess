package standardchess

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/player"
)

var copyBoardFunc = func(board chess.Board, cursor int) chess.Board {
	moveHistory := make([]string, 0, len(board.MoveHistory()))
	for _, move := range board.MoveHistory() {
		moveHistory = append(moveHistory, move.Input())
	}

	copy, err := NewBoardFromMoves(moveHistory[:cursor])
	if err != nil {
		panic(err)
	}

	return copy
}

func NewBoardPlayer(board chess.Board) *player.BoardPlayer {
	return player.NewBoardPlayer(board, copyBoardFunc)
}
