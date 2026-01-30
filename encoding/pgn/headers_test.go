package pgn_test

import (
	"testing"

	"github.com/elaxer/standardchess/encoding/pgn"
	"github.com/stretchr/testify/assert"
)

func TestHeaders_Get(t *testing.T) {
	headers := pgn.Headers{
		pgn.NewHeader("Event", "It (open)"),
		pgn.NewHeader("Site", "Sevilla (Spain)"),
		pgn.NewHeader("Date", "1992.??.??"),
		pgn.NewHeader("Round", "?"),
		pgn.NewHeader("White", "Gonzalez Raul"),
		pgn.NewHeader("Black", "Mikhail Tal"),
		pgn.NewHeader("Result", "0-1"),
		pgn.NewHeader("TimeControl", ""),
		pgn.NewHeader("Link", "https://www.chess.com/games/view/4082964"),
	}

	site, ok := headers.Get("Site")
	assert.Equal(t, pgn.NewHeader("Site", "Sevilla (Spain)"), site)
	assert.True(t, ok)

	result, ok := headers.Get("Result")
	assert.Equal(t, pgn.NewHeader("Result", "0-1"), result)
	assert.True(t, ok)

	unknown, ok := headers.Get("Unknown")
	assert.Equal(t, pgn.Header{}, unknown)
	assert.False(t, ok)
}
