package piece

import "github.com/elaxer/chess"

// New creates a new chess piece based on the provided notation and side.
// Returns nil if the piece cannot be created.
// See chess.PieceFactory for more details.
func New(notation string, side chess.Side) chess.Piece {
	piece, err := NewFactory().CreateFromNotation(notation, side)
	if err != nil {
		return nil
	}

	return piece
}

// FromString creates a new chess piece based on the provided string representation.
// Returns nil if the piece cannot be created.
// See chess.PieceFactory for more details.
func FromString(str string) chess.Piece {
	piece, err := NewFactory().CreateFromString(str)
	if err != nil {
		return nil
	}

	return piece
}
