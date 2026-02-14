package rule_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/encoding/pgn"
	"github.com/elaxer/standardchess/internal/rule"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/elaxer/standardchess/internal/state"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFiftyMoves(t *testing.T) {
	tests := []struct {
		name   string // description of this test case
		fenStr string
		pgnStr string
		want   chess.State
	}{
		{
			"50_moves",
			"kq6/8/8/8/8/8/8/6QK w",
			`1. Qf1 Qc8 2. Qe1 Qd8 3. Qd1 Qe8 4. Qc1 Qf8 5. Qb1 Qg8
6. Qb2 Qg7 7. Qc2 Qf7 8. Qd2 Qe7 9. Qe2 Qf7 10. Qf2 Qg7
11. Qh2 Qg6 12. Qg3 Qf6 13. Qe3 Qe6 14. Qd3 Qd6
15. Qc3 Qc6+ 16. Kg1 Qb6+ 17. Kh1 Qa6
18. Qb3 Qb6 19. Qa3+ Qa6 20. Qb3 Qc8
21. Qc4 Qc5 22. Qd4 Qe5 23. Qf4 Qf5
24. Qg4 Qg5 25. Qh4 Qh5 1/2-1/2`,
			state.FiftyMoves,
		},
		{
			"49_moves",
			"kq6/8/8/8/8/8/8/6QK w",
			`1. Qf1 Qc8 2. Qe1 Qd8 3. Qd1 Qe8 4. Qc1 Qf8 5. Qb1 Qg8
6. Qb2 Qg7 7. Qc2 Qf7 8. Qd2 Qe7 9. Qe2 Qf7 10. Qf2 Qg7
11. Qh2 Qg6 12. Qg3 Qf6 13. Qe3 Qe6 14. Qd3 Qd6
15. Qc3 Qc6+ 16. Kg1 Qb6+ 17. Kh1 Qa6
18. Qb3 Qb6 19. Qa3+ Qa6 20. Qb3 Qc8
21. Qc4 Qc5 22. Qd4 Qe5 23. Qf4 Qf5
24. Qg4 Qg5 25. Qh4 *`,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := standardtest.DecodeFEN(tt.fenStr)
			pgn, err := pgn.FromString(tt.pgnStr)
			require.NoError(t, err)

			for _, move := range pgn.Moves() {
				_, err := board.MakeMove(move)
				require.NoError(t, err)
			}

			got := rule.FiftyMoves(board)
			assert.Equal(t, tt.want, got)
		})
	}
}
