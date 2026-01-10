package pgn_test

import (
	"testing"

	"github.com/elaxer/standardchess/encoding/pgn"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	moves := []string{
		"e4", "c6",
		"d4", "d5",
		"e5", "Bf5",
		"Nc3", "e6",
		"g4", "Bg6",
		"Nge2", "c5",
		"Be3", "Ne7",
		"f4", "h5",
		"f5", "exf5",
		"g5", "Nbc6",
		"Nf4", "a6",
		"Bg2", "cxd4",
		"Bxd4", "Nxd4",
		"Qxd4", "Nc6",
		"Qf2", "Bb4",
		"O-O-O", "Bxc3",
		"bxc3", "Qa5",
		"Rxd5", "Qxc3",
		"Qc5", "Qxc5",
		"Rxc5", "O-O",
		"Bxc6", "bxc6",
		"Rd1", "Rab8",
		"c4", "Rfd8",
		"Rd6", "Kf8",
		"Rcxc6", "Rdc8",
		"Kc2", "h4",
		"Rxc8+", "Rxc8",
		"Kc3", "a5",
		"Ra6", "Rb8",
		"Rxa5", "Rb1",
		"c5", "Re1",
		"Ra8+", "Ke7",
		"Ra7+", "Ke8",
		"Nd3", "Re3",
		"Kd2", "Rh3",
		"c6", "Rxh2+",
		"Ke3", "Rc2",
		"e6", "h3",
		"Nb4", "f4+",
		"Kd4", "h2",
		"Ra8+", "Ke7",
		"Rh8", "Rd2+",
		"Kc5", "Be4",
		"c7", "Bb7",
		"Kb6", "Bc8",
		"Rxc8", "h1=Q",
		"Re8+", "Kxe8",
		"c8=Q+", "Ke7",
		"Nc6+", "Qxc6+",
		"Qxc6", "Rd6",
	}

	pgn := pgn.Encode([]pgn.Header{
		pgn.NewHeader("Event", "Saint Louis Rapid 2017"),
		pgn.NewHeader("Site", "Saint Louis USA"),
		pgn.NewHeader("Date", "2017.08.14"),
		pgn.NewHeader("Round", "?"),
		pgn.NewHeader("White", "Garry Kasparov"),
		pgn.NewHeader("Black", "Navara, David"),
		pgn.NewHeader("Result", "*"),
		pgn.NewHeader("TimeControl", ""),
		pgn.NewHeader("Link", "https://www.chess.com/games/view/14842105"),
	}, moves, "*")
	expectedPGN := `[Event "Saint Louis Rapid 2017"]
[Site "Saint Louis USA"]
[Date "2017.08.14"]
[Round "?"]
[White "Garry Kasparov"]
[Black "Navara, David"]
[Result "*"]
[TimeControl ""]
[Link "https://www.chess.com/games/view/14842105"]

1. e4 c6 2. d4 d5 3. e5 Bf5 4. Nc3 e6 5. g4 Bg6 6. Nge2 c5 7. Be3 Ne7 8. f4 h5
9. f5 exf5 10. g5 Nbc6 11. Nf4 a6 12. Bg2 cxd4 13. Bxd4 Nxd4 14. Qxd4 Nc6 15.
Qf2 Bb4 16. O-O-O Bxc3 17. bxc3 Qa5 18. Rxd5 Qxc3 19. Qc5 Qxc5 20. Rxc5 O-O 21.
Bxc6 bxc6 22. Rd1 Rab8 23. c4 Rfd8 24. Rd6 Kf8 25. Rcxc6 Rdc8 26. Kc2 h4 27.
Rxc8+ Rxc8 28. Kc3 a5 29. Ra6 Rb8 30. Rxa5 Rb1 31. c5 Re1 32. Ra8+ Ke7 33. Ra7+
Ke8 34. Nd3 Re3 35. Kd2 Rh3 36. c6 Rxh2+ 37. Ke3 Rc2 38. e6 h3 39. Nb4 f4+ 40.
Kd4 h2 41. Ra8+ Ke7 42. Rh8 Rd2+ 43. Kc5 Be4 44. c7 Bb7 45. Kb6 Bc8 46. Rxc8
h1=Q 47. Re8+ Kxe8 48. c8=Q+ Ke7 49. Nc6+ Qxc6+ 50. Qxc6 Rd6 *`

	assert.Equal(t, expectedPGN, pgn)
}

func TestEncodeHeaders(t *testing.T) {
	headers := []pgn.Header{
		pgn.NewHeader("my", "header"),
		pgn.NewHeader("anoth", "er header"),
		pgn.NewHeader("foo", "bar"),
		pgn.NewHeader("", "empty"),
		pgn.NewHeader("", ""),
	}

	const expected = `[my "header"]
[anoth "er header"]
[foo "bar"]
[ "empty"]
[ ""]`

	assert.Equal(t, expected, pgn.EncodeHeaders(headers))
}
