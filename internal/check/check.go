// Package check provides utilities for detecting check conditions
// in a standard chess position.
//
// The package contains helpers for evaluating whether the current
// player is in check based on the board state.
package check

import (
	"slices"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
)

func IsCheck(board chess.Board) bool {
	_, kingPosition := board.Squares().FindPiece(piece.NotationKing, board.Turn())

	return board.IsSquareAttacked(kingPosition)
}

func IsKingPseudoAttacked(board chess.Board, color chess.Color) bool {
	_, kingPosition := board.Squares().FindPiece(piece.NotationKing, color)

	for position, piece := range board.Squares().Iter() {
		if piece == nil || piece.Color() == color {
			continue
		}

		pseudoMoves := piece.PseudoMoves(position, board.Squares())
		if slices.Contains(pseudoMoves, kingPosition) {
			return true
		}
	}

	return false
}
