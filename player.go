package standardchess

import "github.com/elaxer/chess"

type BoardPlayer struct {
	moveHistory []chess.Move
	cursor      int
}

func NewBoardPlayer(board chess.Board) *BoardPlayer {
	moveHistory := make([]chess.Move, 0, len(board.MoveHistory()))
	for _, move := range board.MoveHistory() {
		moveHistory = append(moveHistory, move.Move())
	}

	return &BoardPlayer{moveHistory: moveHistory, cursor: len(moveHistory) - 1}
}

func (p *BoardPlayer) Board() chess.Board {
	copy, err := NewBoardFromMoves(p.moveHistory[:p.cursor])
	must(err)

	return copy
}

func (p *BoardPlayer) Reset() {
	p.cursor = 0
}

func (p *BoardPlayer) Prev() (ok bool) {
	return p.GoTo(p.cursor - 1)
}

func (p *BoardPlayer) GoTo(n int) (ok bool) {
	if n < 0 || n >= len(p.moveHistory) {
		return false
	}

	p.cursor = n

	return true
}

func (p *BoardPlayer) Next() (ok bool) {
	return p.GoTo(p.cursor + 1)
}

func (p *BoardPlayer) End() {
	p.cursor = len(p.moveHistory) - 1
}
