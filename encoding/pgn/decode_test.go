package pgn_test

import (
	"slices"
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/chesstest"
	"github.com/elaxer/standardchess/encoding/pgn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	type args struct {
		pgn string
	}
	tests := []struct {
		name        string
		args        args
		wantHeaders []pgn.Header
		wantMoves   []chess.Move
		wantErr     bool
	}{
		{
			"only_headers",
			args{`[Event "It (open)"]
[Site "Sevilla (Spain)"]
[Date "1992.??.??"]
[Round "?"]
[White "Gonzalez Raul"]
[Black "Mikhail Tal"]
[Result "0-1"]
[TimeControl ""]
[Link "https://www.chess.com/games/view/4082964"]`},
			[]pgn.Header{
				pgn.NewHeader("Event", "It (open)"),
				pgn.NewHeader("Site", "Sevilla (Spain)"),
				pgn.NewHeader("Date", "1992.??.??"),
				pgn.NewHeader("Round", "?"),
				pgn.NewHeader("White", "Gonzalez Raul"),
				pgn.NewHeader("Black", "Mikhail Tal"),
				pgn.NewHeader("Result", "0-1"),
				pgn.NewHeader("TimeControl", ""),
				pgn.NewHeader("Link", "https://www.chess.com/games/view/4082964"),
			},
			[]chess.Move{},
			false,
		},
		{
			"only_moves",
			args{`1. e4 c5 2. Nf3 Nc6 3. d4 cxd4 4. Nxd4 Qb6 5. Nb3 Nf6 6. Nc3 e6 7. Be3 Qc7 8. a3
a6 9. Be2 b5 10. O-O Be7 11. Bd3 O-O 12. f4 d6 13. Nd2 Bb7 14. Nf3 Rad8 15. Ng5
h6 16. Nh3 d5 17. e5 Ne4 18. Bxe4 dxe4 19. Qg4 Nd4 20. Bxd4 Rxd4 21. Rad1 Rxd1
22. Rxd1 Bxa3 23. Nxe4 Bxe4 24. bxa3 Qxc2 25. Nf2 Bd5 26. Nd3 h5`},
			[]pgn.Header{},
			chesstest.MoveStrings(
				"e4",
				"c5",
				"Nf3",
				"Nc6",
				"d4",
				"cxd4",
				"Nxd4",
				"Qb6",
				"Nb3",
				"Nf6",
				"Nc3",
				"e6",
				"Be3",
				"Qc7",
				"a3",
				"a6",
				"Be2",
				"b5",
				"O-O",
				"Be7",
				"Bd3",
				"O-O",
				"f4",
				"d6",
				"Nd2",
				"Bb7",
				"Nf3",
				"Rad8",
				"Ng5",
				"h6",
				"Nh3",
				"d5",
				"e5",
				"Ne4",
				"Bxe4",
				"dxe4",
				"Qg4",
				"Nd4",
				"Bxd4",
				"Rxd4",
				"Rad1",
				"Rxd1",
				"Rxd1",
				"Bxa3",
				"Nxe4",
				"Bxe4",
				"bxa3",
				"Qxc2",
				"Nf2",
				"Bd5",
				"Nd3",
				"h5",
			),
			false,
		},
		{
			"raul_vs_tal",
			args{`[Event "It (open)"]
[Site "Sevilla (Spain)"]
[Date "1992.??.??"]
[Round "?"]
[White "Gonzalez Raul"]
[Black "Mikhail Tal"]
[Result "0-1"]
[TimeControl ""]
[Link "https://www.chess.com/games/view/4082964"]

1. e4 c5 2. Nf3 Nc6 3. d4 cxd4 4. Nxd4 Qb6 5. Nb3 Nf6 6. Nc3 e6 7. Be3 Qc7 8. a3
a6 9. Be2 b5 10. O-O Be7 11. Bd3 O-O 12. f4 d6 13. Nd2 Bb7 14. Nf3 Rad8 15. Ng5
h6 16. Nh3 d5 17. e5 Ne4 18. Bxe4 dxe4 19. Qg4 Nd4 20. Bxd4 Rxd4 21. Rad1 Rxd1
22. Rxd1 Bxa3 23. Nxe4 Bxe4 24. bxa3 Qxc2 25. Nf2 Bd5 26. Nd3 h5 0-1`},
			[]pgn.Header{
				pgn.NewHeader("Event", "It (open)"),
				pgn.NewHeader("Site", "Sevilla (Spain)"),
				pgn.NewHeader("Date", "1992.??.??"),
				pgn.NewHeader("Round", "?"),
				pgn.NewHeader("White", "Gonzalez Raul"),
				pgn.NewHeader("Black", "Mikhail Tal"),
				pgn.NewHeader("Result", "0-1"),
				pgn.NewHeader("TimeControl", ""),
				pgn.NewHeader("Link", "https://www.chess.com/games/view/4082964"),
			},
			chesstest.MoveStrings(
				"e4",
				"c5",
				"Nf3",
				"Nc6",
				"d4",
				"cxd4",
				"Nxd4",
				"Qb6",
				"Nb3",
				"Nf6",
				"Nc3",
				"e6",
				"Be3",
				"Qc7",
				"a3",
				"a6",
				"Be2",
				"b5",
				"O-O",
				"Be7",
				"Bd3",
				"O-O",
				"f4",
				"d6",
				"Nd2",
				"Bb7",
				"Nf3",
				"Rad8",
				"Ng5",
				"h6",
				"Nh3",
				"d5",
				"e5",
				"Ne4",
				"Bxe4",
				"dxe4",
				"Qg4",
				"Nd4",
				"Bxd4",
				"Rxd4",
				"Rad1",
				"Rxd1",
				"Rxd1",
				"Bxa3",
				"Nxe4",
				"Bxe4",
				"bxa3",
				"Qxc2",
				"Nf2",
				"Bd5",
				"Nd3",
				"h5",
			),
			false,
		},
		{
			"carlsen_vs_ilinca",
			args{`[Event "January 07 Late 2025"]
[Site ""]
[Date "2025.01.07"]
[Round "?"]
[White "Magnus Carlsen"]
[Black "Felix Antonio Ilinca Ilinca"]
[Result "1-0"]
[TimeControl ""]
[Link "https://www.chess.com/games/view/17633485"]

1. d4 d5 2. c4 e6 3. Nf3 Nf6 4. g3 Be7 5. Bg2 O-O 6. O-O dxc4 7. Nc3 a6 8. Ne5
Ra7 9. Nxc4 b5 10. Ne5 Bb7 11. Bxb7 Rxb7 12. Be3 Nd5 13. Nxd5 Qxd5 14. Qb3 Rd8
15. Rfc1 c5 16. dxc5 Qxe5 17. c6 Nxc6 18. Rxc6 a5 19. Rd1 Rxd1+ 20. Qxd1 h6 21.
Rc8+ Kh7 22. Qc2+ Qf5 23. Qc6 Qd5 24. Qxd5 exd5 25. Bd4 Kg6 26. Kg2 f6 27. Kf3
Kf5 28. Rc6 a4 29. g4+ Kg6 30. h4 Kf7 31. h5 b4 32. Bc5 a3 33. b3 Bxc5 34. Rxc5
Ke6 35. Ke3 Kd6 36. Kd4 Ke6 37. f4 f5 38. Rc6+ Kd7 39. Rg6 fxg4 40. Rxg7+ Kc6
41. Rxg4 Re7 42. e3 Re4+ 43. Kd3 Kc5 44. Rg6 Re8 45. Rxh6 Rg8 46. Rg6 Rh8 47. h6
1-0`},
			[]pgn.Header{
				pgn.NewHeader("Event", "January 07 Late 2025"),
				pgn.NewHeader("Site", ""),
				pgn.NewHeader("Date", "2025.01.07"),
				pgn.NewHeader("Round", "?"),
				pgn.NewHeader("White", "Magnus Carlsen"),
				pgn.NewHeader("Black", "Felix Antonio Ilinca Ilinca"),
				pgn.NewHeader("Result", "1-0"),
				pgn.NewHeader("TimeControl", ""),
				pgn.NewHeader("Link", "https://www.chess.com/games/view/17633485"),
			},
			chesstest.MoveStrings(
				"d4",
				"d5",
				"c4",
				"e6",
				"Nf3",
				"Nf6",
				"g3",
				"Be7",
				"Bg2",
				"O-O",
				"O-O",
				"dxc4",
				"Nc3",
				"a6",
				"Ne5",
				"Ra7",
				"Nxc4",
				"b5",
				"Ne5",
				"Bb7",
				"Bxb7",
				"Rxb7",
				"Be3",
				"Nd5",
				"Nxd5",
				"Qxd5",
				"Qb3",
				"Rd8",
				"Rfc1",
				"c5",
				"dxc5",
				"Qxe5",
				"c6",
				"Nxc6",
				"Rxc6",
				"a5",
				"Rd1",
				"Rxd1",
				"Qxd1",
				"h6",
				"Rc8",
				"Kh7",
				"Qc2",
				"Qf5",
				"Qc6",
				"Qd5",
				"Qxd5",
				"exd5",
				"Bd4",
				"Kg6",
				"Kg2",
				"f6",
				"Kf3",
				"Kf5",
				"Rc6",
				"a4",
				"g4",
				"Kg6",
				"h4",
				"Kf7",
				"h5",
				"b4",
				"Bc5",
				"a3",
				"b3",
				"Bxc5",
				"Rxc5",
				"Ke6",
				"Ke3",
				"Kd6",
				"Kd4",
				"Ke6",
				"f4",
				"f5",
				"Rc6",
				"Kd7",
				"Rg6",
				"fxg4",
				"Rxg7",
				"Kc6",
				"Rxg4",
				"Re7",
				"e3",
				"Re4",
				"Kd3",
				"Kc5",
				"Rg6",
				"Re8",
				"Rxh6",
				"Rg8",
				"Rg6",
				"Rh8",
				"h6",
			),

			false,
		},
		{
			"tal_vs_lautier",
			args{`[Event "It"]
[Site "Barcelona (Spain)"]
[Date "1992.??.??"]
[Round "?"]
[White "Mikhail Tal"]
[Black "Joel Lautier"]
[Result "1-0"]
[TimeControl ""]
[Link "https://www.chess.com/games/view/4082949"]

1. d4 Nf6 2. Nf3 e6 3. g3 b5 4. Bg2 Bb7 5. O-O c5 6. Bg5 Qb6 7. a4 a6 8. Nc3 Ne4
9. Nxe4 Bxe4 10. axb5 Qxb5 11. Qd2 f6 12. Bf4 Qb7 13. c4 cxd4 14. Qxd4 e5 15.
Bxe5 fxe5 16. Qxe5+ Be7 17. Nd4 Bxg2 18. Nf5 Qb4 19. Kxg2 Nc6 20. Qxg7 O-O-O 21.
Rxa6 Qb7 22. Rfa1 Nb4+ 23. Kg1 Nxa6 24. Qxe7 Qb6 25. Qa3 Rhf8 26. Nd6+ Kc7 27.
Qxa6 Ra8 28. Qxb6+ Kxb6 29. Rd1 Ra2 30. Rd2 Kc6 31. f3 Rfa8 32. Nb5 R8a4 33. Rc2
Kc5 34. Nc3 Ra1+ 35. Kf2 Rxc4 36. Rd2 Ra7 37. e4 Kc6 38. Ke3 Rb7 39. Rc2 d6 40.
Kd3 Rb5 41. f4 Rbb4 42. g4 Kd7 43. g5 Ke6 44. h4 d5 45. Nxd5 Rxc2 46. Nxb4 Rxb2
47. Nc2 Rb3+ 48. Kc4 Rh3 49. Nd4+ Kf7 50. f5 Rxh4 51. Kd5 Rg4 52. Nf3 Rg3 53.
Ne5+ Kg8 54. f6 Rxg5 55. Ke6 Rg1 56. f7+ Kg7 57. Nd7 Rf1 58. f8=Q+ Rxf8 59. Nxf8
h6 60. Nd7 h5 61. Ne5 h4 62. Nf3# 1-0`},
			[]pgn.Header{
				pgn.NewHeader("Event", "It"),
				pgn.NewHeader("Site", "Barcelona (Spain)"),
				pgn.NewHeader("Date", "1992.??.??"),
				pgn.NewHeader("Round", "?"),
				pgn.NewHeader("White", "Mikhail Tal"),
				pgn.NewHeader("Black", "Joel Lautier"),
				pgn.NewHeader("Result", "1-0"),
				pgn.NewHeader("TimeControl", ""),
				pgn.NewHeader("Link", "https://www.chess.com/games/view/4082949"),
			},
			chesstest.MoveStrings(
				"d4",
				"Nf6",
				"Nf3",
				"e6",
				"g3",
				"b5",
				"Bg2",
				"Bb7",
				"O-O",
				"c5",
				"Bg5",
				"Qb6",
				"a4",
				"a6",
				"Nc3",
				"Ne4",
				"Nxe4",
				"Bxe4",
				"axb5",
				"Qxb5",
				"Qd2",
				"f6",
				"Bf4",
				"Qb7",
				"c4",
				"cxd4",
				"Qxd4",
				"e5",
				"Bxe5",
				"fxe5",
				"Qxe5",
				"Be7",
				"Nd4",
				"Bxg2",
				"Nf5",
				"Qb4",
				"Kxg2",
				"Nc6",
				"Qxg7",
				"O-O-O",
				"Rxa6",
				"Qb7",
				"Rfa1",
				"Nb4",
				"Kg1",
				"Nxa6",
				"Qxe7",
				"Qb6",
				"Qa3",
				"Rhf8",
				"Nd6",
				"Kc7",
				"Qxa6",
				"Ra8",
				"Qxb6",
				"Kxb6",
				"Rd1",
				"Ra2",
				"Rd2",
				"Kc6",
				"f3",
				"Rfa8",
				"Nb5",
				"R8a4",
				"Rc2",
				"Kc5",
				"Nc3",
				"Ra1",
				"Kf2",
				"Rxc4",
				"Rd2",
				"Ra7",
				"e4",
				"Kc6",
				"Ke3",
				"Rb7",
				"Rc2",
				"d6",
				"Kd3",
				"Rb5",
				"f4",
				"Rbb4",
				"g4",
				"Kd7",
				"g5",
				"Ke6",
				"h4",
				"d5",
				"Nxd5",
				"Rxc2",
				"Nxb4",
				"Rxb2",
				"Nc2",
				"Rb3",
				"Kc4",
				"Rh3",
				"Nd4",
				"Kf7",
				"f5",
				"Rxh4",
				"Kd5",
				"Rg4",
				"Nf3",
				"Rg3",
				"Ne5",
				"Kg8",
				"f6",
				"Rxg5",
				"Ke6",
				"Rg1",
				"f7",
				"Kg7",
				"Nd7",
				"Rf1",
				"f8=Q",
				"Rxf8",
				"Nxf8",
				"h6",
				"Nd7",
				"h5",
				"Ne5",
				"h4",
				"Nf3",
			),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeaders, gotMoves, err := pgn.Decode(tt.args.pgn)

			require.Truef(
				t,
				(err != nil) == tt.wantErr,
				"Decode() error = %v, wantErr %v",
				err,
				tt.wantErr,
			)
			if tt.wantErr {
				return
			}

			assert.True(t, slices.Equal(gotHeaders, tt.wantHeaders))
			assert.True(t, slices.Equal(gotMoves, tt.wantMoves))
		})
	}
}
