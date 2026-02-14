package rule

import (
	"slices"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/state"
)

var strongMaterial = []string{piece.NotationQueen, piece.NotationRook, piece.NotationPawn}

func InsufficientMaterial(board chess.Board) chess.State {
	whitePiecesCount := 0
	blackPiecesCount := 0
	for _, p := range board.Squares().Iter() {
		if p == nil {
			continue
		}
		if slices.Contains(strongMaterial, p.Notation()) {
			return nil
		}

		switch p.Color() {
		case chess.ColorWhite:
			whitePiecesCount++
		case chess.ColorBlack:
			blackPiecesCount++
		}
	}

	if whitePiecesCount > 2 || blackPiecesCount > 2 {
		return nil
	}

	return state.InsufficientMaterial
}
