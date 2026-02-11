package rule_test

import (
	"testing"

	"github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/internal/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestThreefoldRepetition_Rule(t *testing.T) {
	tests := []struct {
		name  string
		moves []string
		want  bool
	}{
		{
			"repetition",
			[]string{
				"d4", "d5",
				"Qd3", "Qd6",
				"Qe4", "Qe5",
				"Qf4", "Qf5",
				"Qe4", "Qe5",
				"Qf4", "Qf5",
				"Qe4", "Qe5",
			},
			true,
		},
		{
			"no_repetition",
			[]string{
				"d4", "d5",
				"Qd3", "Qd6",
				"Qe4", "Qe5",
				"Qf4", "Qf5",
				"Qe4", "Qe5",
				"Qf4", "Qf5",
				"Qe4",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board, err := standardchess.NewBoardFromMoves(tt.moves)
			require.NoError(t, err)

			assert.Equal(t, tt.want, board.State() == state.ThreefoldRepetition)
		})
	}
}

func TestThreefoldRepetition_Rule_WithUndo(t *testing.T) {
	moves := []string{
		"d4", "d5",
		"Qd3", "Qd6",
		"Qe4", "Qe5",
		"Qf4", "Qf5",
		"Qe4", "Qe5",
		"Qf4", "Qf5",
		"Qe4", "Qe5",
	}

	board, err := standardchess.NewBoardFromMoves(moves)
	require.NoError(t, err)

	require.Equal(t, state.ThreefoldRepetition, board.State())

	_, err = board.UndoLastMove()
	require.NoError(t, err)

	assert.NotEqual(t, state.ThreefoldRepetition, board.State())
}
