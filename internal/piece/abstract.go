package piece

import "github.com/elaxer/chess"

// abstract это базовая структура для шахматной фигуры.
// Она содержит базовые поля и вспомогательные методы для работы с фигурами.
type abstract struct {
	side    chess.Side
	isMoved bool
}

func (p *abstract) Side() chess.Side {
	return p.side
}

func (p *abstract) IsMoved() bool {
	return p.isMoved
}

func (p *abstract) SetIsMoved(isMoved bool) {
	p.isMoved = isMoved
}

// canMove проверяет, может ли фигура переместиться на указанную клетку.
// Если клетка существует и пуста или занята фигурой противника, то перемещение возможно.
func (p *abstract) canMove(squarePiece chess.Piece, pieceSide chess.Side) bool {
	return squarePiece == nil || pieceSide != squarePiece.Side()
}
