package rule_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/rule"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/elaxer/standardchess/internal/state"
	"github.com/stretchr/testify/assert"
)

func TestInsufficientMaterial(t *testing.T) {
	tests := []struct {
		name   string
		fenStr string
		want   chess.State
	}{
		{
			"k_vs_k",
			"8/2k5/8/8/8/3K4/8/8 w - - 1 1",
			state.InsufficientMaterial,
		},
		{
			"k+b_vs_k",
			"8/2kb4/8/8/8/3K4/8/8 w - - 1 1",
			state.InsufficientMaterial,
		},
		{
			"k_vs_k+n",
			"8/2k5/8/8/8/3KN3/8/8 w - - 1 1",
			state.InsufficientMaterial,
		},
		{
			"k+b_vs_k+n",
			"8/2kb4/8/8/8/3KN3/8/8 w - - 1 1",
			state.InsufficientMaterial,
		},
		{
			"k+b_vs_k+b",
			"8/2kb4/8/8/8/3KB3/8/8 w - - 1 1",
			state.InsufficientMaterial,
		},
		{
			"k+n_vs_k+n",
			"8/2kn4/8/8/8/3KN3/8/8 w - - 1 1",
			state.InsufficientMaterial,
		},

		{
			"k+q_vs_k",
			"8/2kq4/8/8/8/3K4/8/8 w - - 1 1",
			nil,
		},
		{
			"k_vs_k+p",
			"8/2k5/8/8/8/3KP3/8/8 w - - 1 1",
			nil,
		},
		{
			"k+r_vs_k",
			"8/2kr4/8/8/8/3K4/8/8 w - - 1 1",
			nil,
		},
		{
			"k+2n_vs_k",
			"8/2knn3/8/8/8/3K4/8/8 w - - 1 1",
			nil,
		},
		{
			"k+3b_vs_k",
			"8/2kbbb2/8/8/8/3K4/8/8 w - - 1 1",
			nil,
		},
		{
			"k+2b_vs_kb",
			"8/2kbb3/8/8/8/3KB3/8/8 w - - 1 1",
			nil,
		},
		{
			"k+b+n_vs_k",
			"8/2kbn3/8/8/8/3K4/8/8 w - - 1 1",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rule.InsufficientMaterial(standardtest.DecodeFEN(tt.fenStr))
			assert.Equal(t, tt.want, got)
		})
	}
}
