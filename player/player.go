package player

import "github.com/elaxer/chess"

type CopyBoardFunc func(board chess.Board, cursor int) chess.Board

type BoardPlayer struct {
	board         chess.Board
	copyBoardFunc CopyBoardFunc
	cursor        int
}

func NewBoardPlayer(board chess.Board, copyBoardFunc CopyBoardFunc) *BoardPlayer {
	player := &BoardPlayer{board: board, copyBoardFunc: copyBoardFunc}
	player.End()

	return player
}

func (p *BoardPlayer) Snapshot() chess.Board {
	cursor := min(p.cursor, len(p.board.MoveHistory()))

	return p.copyBoardFunc(p.board, cursor)
}

func (p *BoardPlayer) Cursor() int {
	return p.cursor
}

func (p *BoardPlayer) Reset() {
	p.cursor = 0
}

func (p *BoardPlayer) Prev() (ok bool) {
	if p.cursor == 0 {
		return false
	}

	return p.GoTo(p.cursor - 1)
}

func (p *BoardPlayer) GoTo(n int) (ok bool) {
	if n > len(p.board.MoveHistory()) {
		return false
	}

	p.cursor = n

	return true
}

func (p *BoardPlayer) Next() (ok bool) {
	return p.GoTo(p.cursor + 1)
}

func (p *BoardPlayer) End() {
	p.cursor = len(p.board.MoveHistory())
}
