package rule_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/internal/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFivefoldRepetition_Rule(t *testing.T) {
	tests := []struct {
		name  string
		moves []string
		want  chess.State
	}{
		{
			"repetition",
			[]string{
				"d4", "d5",
				"e4", "e5",
				"dxe5", "dxe4",

				"Qd4", "Qd5",
				"Qd3", "Qd6",

				"Qd4", "Qd5",
				"Qd3", "Qd6",

				"Qd4", "Qd5",
				"Qd3", "Qd6",

				"Qd4", "Qd5",
				"Qd3", "Qd6",

				"Qd4", "Qd5",
			},
			state.FivefoldRepetition,
		},
		{
			"no_repetition",
			[]string{
				"d4", "d5",
				"e4", "e5",
				"dxe5", "dxe4",

				"Qd4", "Qd5",
				"Qd3", "Qd6",

				"Qd4", "Qd5",
				"Qd3", "Qd6",

				"Qd4", "Qd5",
				"Qd3", "Qd6",

				"Qd4", "Qd5",
				"Qd3", "Qd6",

				"Qd4",
			},
			chess.StateClear,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board, err := standardchess.NewBoardFromMoves(tt.moves)
			require.NoError(t, err)

			assert.Equal(t, tt.want, board.State())
		})
	}
}

func TestFivefoldRepetition_Rule_WithUndo(t *testing.T) {
	moves := []string{
		"d4", "d5",
		"e4", "e5",
		"dxe5", "dxe4",

		"Qd4", "Qd5",
		"Qd3", "Qd6",

		"Qd4", "Qd5",
		"Qd3", "Qd6",

		"Qd4", "Qd5",
		"Qd3", "Qd6",

		"Qd4", "Qd5",
		"Qd3", "Qd6",

		"Qd4", "Qd5",
	}

	board, err := standardchess.NewBoardFromMoves(moves)
	require.NoError(t, err)

	require.Equal(t, state.FivefoldRepetition, board.State())

	_, err = board.UndoLastMove()
	require.NoError(t, err)

	assert.Equal(t, chess.StateClear, board.State())
}
