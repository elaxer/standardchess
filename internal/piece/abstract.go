package piece

import "github.com/elaxer/chess"

type abstract struct {
	color   chess.Color
	isMoved bool
}

func (p *abstract) Color() chess.Color {
	return p.color
}

func (p *abstract) IsMoved() bool {
	return p.isMoved
}

func (p *abstract) SetIsMoved(isMoved bool) {
	p.isMoved = isMoved
}

func (p *abstract) canMove(squarePiece chess.Piece, pieceColor chess.Color) bool {
	return squarePiece == nil || pieceColor != squarePiece.Color()
}
