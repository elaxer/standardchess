package standardchess

import "github.com/elaxer/chess"

type BoardPlayer struct {
	board  chess.Board
	cursor uint16
}

func NewBoardPlayer(board chess.Board) *BoardPlayer {
	player := &BoardPlayer{board: board}
	player.End()

	return player
}

func (p *BoardPlayer) Board() chess.Board {
	cursor := min(p.cursor, p.moveHistoryLen())

	copy, err := NewBoardFromMoves(p.moveHistory()[:cursor])
	must(err)

	return copy
}

func (p *BoardPlayer) Cursor() uint16 {
	return p.cursor
}

func (p *BoardPlayer) Reset() {
	p.cursor = 0
}

func (p *BoardPlayer) Prev() (ok bool) {
	return p.GoTo(p.cursor - 1)
}

func (p *BoardPlayer) GoTo(n uint16) (ok bool) {
	if n > p.moveHistoryLen() {
		return false
	}

	p.cursor = n

	return true
}

func (p *BoardPlayer) Next() (ok bool) {
	return p.GoTo(p.cursor + 1)
}

func (p *BoardPlayer) End() {
	p.cursor = p.moveHistoryLen()
}

func (p *BoardPlayer) moveHistory() []chess.Move {
	moveHistory := make([]chess.Move, 0, p.moveHistoryLen())
	for _, move := range p.board.MoveHistory() {
		moveHistory = append(moveHistory, move.Move())
	}

	return moveHistory
}

func (p *BoardPlayer) moveHistoryLen() uint16 {
	//nolint:gosec
	return uint16(len(p.board.MoveHistory()))
}
