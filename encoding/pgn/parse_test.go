package pgn_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/elaxer/standardchess/encoding/pgn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	reader := strings.NewReader(pgnStr)

	i := 1
	for pgn, err := range pgn.Parse(reader) {
		t.Run("No. "+strconv.Itoa(i), func(t *testing.T) {
			require.NoError(t, err)
			require.Equal(t, expected[i-1], pgn.Format(0))
		})
		i++
	}

	assert.Equal(t, len(expected), i-1)
}

//nolint:decorder
var expected = [...]string{
	//nolint:lll
	`[Event "Rated Blitz game"]
[Site "https://lichess.org/l68uct88"]
[White "Alik"]
[Black "ioi"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:44:47"]
[WhiteElo "1534"]
[BlackElo "1299"]
[WhiteRatingDiff "-19"]
[BlackRatingDiff "+20"]
[ECO "D00"]
[Opening "Queen's Pawn Game: Mason Attack"]
[TimeControl "360+0"]
[Termination "Normal"]

1. d4 d5 2. Bf4 Nf6 3. e3 e6 4. Nc3 a6 5. Nh3 Ne4 6. Nxe4 dxe4 7. a3 Bd6 8. Bg5 f6 9. Qh5+ g6 10. Qh6 fxg5 11. Nxg5 Bf8 12. Qh4 Nc6 13. Bc4 b5 14. Bxe6 Be7 15. Qxe4 Bxe6 16. Nxe6 Qd7 17. Nc5 Qf5 18. Qxc6+ Kf7 19. Qd5+ Qxd5 0-1`,

	//nolint:lll
	`[Event "Rated Classical game"]
[Site "https://lichess.org/5psccdm0"]
[White "jajce3"]
[Black "jeh-1"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:44:31"]
[WhiteElo "1684"]
[BlackElo "1799"]
[WhiteRatingDiff "+15"]
[BlackRatingDiff "-17"]
[ECO "D10"]
[Opening "Slav Defense: Exchange Variation"]
[TimeControl "300+15"]
[Termination "Normal"]

1. d4 d5 2. c4 c6 3. cxd5 cxd5 4. Nc3 e6 5. a3 Nc6 6. e3 a6 7. b4 b5 8. Qb3 Bd7 9. g3 Nf6 10. Bg2 Bd6 11. Nge2 O-O 12. O-O Re8 13. Bd2 Qb6 14. Na2 e5 15. Bc3 e4 16. Bd2 a5 17. Nec3 axb4 18. Nxd5 Nxd5 19. Qxd5 Ne7 20. Qxe4 bxa3 21. Bb4 Nc6 22. Qb1 Rac8 23. Bxd6 Na5 24. Bc5 Qa6 25. Nb4 Qf6 26. Nd5 Qd8 27. Bb6 Qg5 28. Bxa5 Be6 29. Nb6 Rb8 30. Rxa3 Bc4 31. Nxc4 bxc4 32. Qe1 Rec8 33. Rc3 Qxa5 34. Qc1 Qa4 35. Bd5 Rb4 36. Qc2 Qxc2 37. Rxc2 c3 38. Ra1 Rbb8 39. Kf1 g6 40. Ke2 Rb2 41. Rxb2 cxb2 42. Rb1 Rc2+ 43. Kd1 Rxf2 44. h3 Kg7 45. Bb3 Rh2 46. h4 Rg2 47. Bc2 Rxg3 48. e4 Rh3 49. d5 Rxh4 50. Rxb2 h5 51. d6 Kf6 52. d7 Ke7 53. Rb7 Rh1+ 54. Kd2 Ra1 55. Bd3 Ra8 56. Bb5 f6 57. Rc7 Kd8 58. Rc6 Rb8 59. Ba4 g5 60. Rxf6 Ke7 61. Rh6 Rb4 62. Bc6 Rd4+ 63. Ke3 Rc4 64. Rh8 Rxc6 65. d8=Q+ 1-0`,

	//nolint:lll
	`[Event "Rated Bullet game"]
[Site "https://lichess.org/0s6fmzyx"]
[White "SkyBang"]
[Black "KAINVALDOMERO"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:10"]
[WhiteElo "1682"]
[BlackElo "1242"]
[WhiteRatingDiff "-145"]
[BlackRatingDiff "+186"]
[ECO "C11"]
[Opening "French Defense: Classical Variation, Steinitz Variation"]
[TimeControl "120+0"]
[Termination "Normal"]

1. d4 d5 2. Nc3 e6 3. e4 Nf6 4. e5 Nfd7 5. Nf3 f6 6. exf6 Nxf6 7. Bg5 Be7 8. Ne5 h6 9. Bxf6 Bxf6 10. Qd3 O-O 11. Ng6 Re8 12. Be2 Nc6 13. Bh5 Bxd4 14. Nf4 Qg5 15. g3 e5 16. Bxe8 exf4 17. O-O-O fxg3+ 18. Kb1 Bxc3 19. bxc3 gxf2 20. h4 Qe7 21. Bxc6 bxc6 22. Qf3 Rb8+ 23. Ka1 Be6 24. Qxf2 Qa3 25. Rb1 d4 26. Rb3 Rxb3 27. cxb3 Bxb3 28. Qxd4 Qxa2# 0-1`,

	//nolint:lll
	`[Event "Rated Blitz game"]
[Site "https://lichess.org/1qx160x0"]
[White "Artem84"]
[Black "zine1971"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:05"]
[WhiteElo "1645"]
[BlackElo "1437"]
[WhiteRatingDiff "-16"]
[BlackRatingDiff "+44"]
[ECO "D00"]
[Opening "Queen's Pawn Game #2"]
[TimeControl "300+0"]
[Termination "Normal"]

1. d4 d5 2. e3 e6 3. Nf3 h6 4. c4 c6 5. Nc3 Bb4 6. Bd2 Nf6 7. a3 Bxc3 8. Bxc3 Ne4 9. Qc2 O-O 10. Nd2 Nd6 11. c5 Nf5 12. Bd3 Ne7 13. O-O f5 14. b4 Ng6 15. Nf3 Nh4 16. Nxh4 Qxh4 17. g3 Qh3 18. Rfe1 Nd7 19. Bf1 Qh5 20. Qe2 Qg6 21. Bg2 Nf6 22. b5 Bd7 23. a4 Rae8 24. Bf1 Ne4 25. Qc2 h5 26. bxc6 bxc6 27. Rab1 h4 28. Rb7 hxg3 29. fxg3 Bc8 30. Rxa7 Nxg3 31. hxg3 Qxg3+ 32. Qg2 Qh4 33. Qh2 Qg5+ 34. Qg2 Qf6 35. Re2 g5 36. Qh2 Rf7 37. Rxf7 Qxf7 38. Qe5 Qh5 39. Rh2 Qg4+ 40. Rg2 Qh4 41. Rh2 Qg4+ 42. Bg2 Qd1+ 43. Bf1 Qg4+ 44. Rg2 Qh4 45. Qf6 g4 46. Rh2 Qg3+ 47. Rg2 Qxe3+ 48. Kh1 Qh3+ 49. Rh2 Qxf1# 0-1`,

	//nolint:lll
	`[Event "Rated Blitz tournament https://lichess.org/tournament/wlvsn5j1"]
[Site "https://lichess.org/2wf8ta39"]
[White "neymardossantos"]
[Black "andresneno"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:18"]
[WhiteElo "1635"]
[BlackElo "1718"]
[WhiteRatingDiff "-9"]
[BlackRatingDiff "+9"]
[ECO "C23"]
[Opening "Bishop's Opening"]
[TimeControl "300+1"]
[Termination "Normal"]

1. e4 e5 2. Bc4 Nc6 3. Nc3 Bc5 4. Nf3 d6 5. d3 h6 6. a3 a6 7. Be3 Nd4 8. b4 Ba7 9. O-O Bg4 10. h3 Bxf3 11. gxf3 Qf6 12. Ne2 Nxf3+ 13. Kh1 Qh4 14. Ng1 Nxg1 15. Kxg1 Qxh3 16. Bxa7 Rxa7 17. f4 Qg3+ 18. Kh1 exf4 19. Qe2 Nf6 20. Rf2 Ng4 21. Rg2 Qh3+ 22. Kg1 f3 0-1`,

	//nolint:lll
	`[Event "Rated Blitz tournament https://lichess.org/tournament/wlvsn5j1"]
[Site "https://lichess.org/wzfkda1w"]
[White "Federico"]
[Black "AqibNabi01"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:18"]
[WhiteElo "1874"]
[BlackElo "1433"]
[WhiteRatingDiff "+2"]
[BlackRatingDiff "-12"]
[ECO "C00"]
[Opening "French Defense: Normal Variation"]
[TimeControl "300+1"]
[Termination "Normal"]

1. d4 e6 2. e4 Qh4 3. Nc3 Nf6 4. Bd3 Ng4 5. g3 Qf6 6. Qxg4 e5 7. dxe5 Qxe5 8. Nf3 d6 9. Qxc8+ Ke7 10. Qxc7+ Ke6 11. Nxe5 dxe5 1-0`,

	//nolint:lll
	`[Event "Rated Blitz game"]
[Site "https://lichess.org/m91p5vcs"]
[White "macface"]
[Black "deadpawn"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:25"]
[WhiteElo "1703"]
[BlackElo "1369"]
[WhiteRatingDiff "+3"]
[BlackRatingDiff "-21"]
[ECO "A43"]
[Opening "Benoni Defense: Benoni Gambit Accepted"]
[TimeControl "300+0"]
[Termination "Time forfeit"]

1. d4 c5 2. dxc5 Qa5+ 3. c3 Qxc5 4. g3 Nc6 5. Bg2 Nf6 6. e3 g6 7. Ne2 Bg7 8. O-O h5 9. Nd2 g5 10. Nd4 d6 11. Nd2b3 Qb6 12. Nxc6 bxc6 13. Nd4 Rb8 14. Bxc6+ Bd7 15. Bg2 Bg4 16. f3 Bd7 17. Kh1 g4 18. f4 h4 19. gxh4 Rxh4 20. Qe1 Rh5 21. e4 Rh7 22. f5 Bh6 23. Bxh6 Rxh6 24. Qg3 Qxb2 25. Rab1 Qxc3 26. Rxb8+ Bc8 27. Rxc8+ Qxc8 28. e5 dxe5 29. Qxe5 Qc3 30. Qb8+ Kd7 31. Bc6+ Qxc6+ 32. Nxc6 Kxc6 33. Qxa7 Ne4 34. Qa8+ Kc5 35. Qxe4 Rd6 36. Qxg4 1-0`,

	//nolint:lll
	`[Event "Rated Classical game"]
[Site "https://lichess.org/cu3hedz6"]
[White "SAIRAM"]
[Black "muratk"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:52"]
[WhiteElo "1278"]
[BlackElo "1505"]
[WhiteRatingDiff "-5"]
[BlackRatingDiff "+5"]
[ECO "A00"]
[Opening "Van't Kruijs Opening"]
[TimeControl "300+5"]
[Termination "Normal"]

1. e3 d5 2. a3 e5 3. b4 Nc6 4. Bb2 d4 5. exd4 exd4 6. Bb5 Qe7+ 7. Ne2 Kd8 8. Bxc6 bxc6 9. Bxd4 Ba6 10. d3 Nf6 11. Bxf6 gxf6 12. O-O Rg8 13. g3 Bb7 14. Nbc3 c5 15. bxc5 Qxc5 16. Na4 Qc6 17. c4 Qg2# 0-1`,

	`[Event "Rated Blitz game"]
[Site "https://lichess.org/mm1fckgq"]
[White "mchelken"]
[Black "rasmussenesq"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:58"]
[WhiteElo "1618"]
[BlackElo "1677"]
[WhiteRatingDiff "-9"]
[BlackRatingDiff "+9"]
[ECO "B40"]
[Opening "Sicilian Defense: French Variation"]
[TimeControl "180+0"]
[Termination "Normal"]

1. e4 c5 2. Nf3 e6 3. Bc4 a6 4. Nc3 b5 5. Bb3 c4 6. Nxb5 axb5 7. Bxc4 bxc4 8. d4 cxd3 9. Qxd3 Ba6 10. Qa3 Bxa3 0-1`,

	//nolint:lll
	`[Event "Rated Bullet game"]
[Site "https://lichess.org/io1kkxaj"]
[White "Panevis"]
[Black "Arthzu"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:46:22"]
[WhiteElo "2036"]
[BlackElo "1892"]
[WhiteRatingDiff "+9"]
[BlackRatingDiff "-7"]
[ECO "B10"]
[Opening "Caro-Kann Defense"]
[TimeControl "0+1"]
[Termination "Normal"]

1. e4 c6 2. Nf3 d6 3. d4 g6 4. Bd3 Bg7 5. O-O Nd7 6. Re1 Ngf6 7. c3 e5 8. Bg5 h6 9. Bh4 g5 10. Bg3 O-O 11. Nbd2 Nh5 12. Qe2 Nf4 13. Bxf4 gxf4 14. g3 Qf6 15. Nc4 Kh8 16. Kh1 Rg8 17. dxe5 Nxe5 18. Ncxe5 dxe5 19. Bc4 Bg4 20. Rf1 fxg3 21. fxg3 Bxf3+ 22. Rxf3 Qg5 23. Raf1 f6 24. Bxg8 Kxg8 25. Rf5 Qg6 26. Qf3 Kh7 27. Rf2 Rd8 28. h4 Rd1+ 29. Qxd1 1-0`,

	//nolint:lll
	`[Event "Rated Bullet game"]
[Site "https://lichess.org/qpwvedsb"]
[White "dvorak"]
[Black "nichiren1967"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:46:55"]
[WhiteElo "1559"]
[BlackElo "1770"]
[WhiteRatingDiff "+19"]
[BlackRatingDiff "-17"]
[ECO "C41"]
[Opening "Philidor Defense"]
[TimeControl "120+0"]
[Termination "Normal"]

1. e4 e5 2. Nf3 d6 3. d4 f6 4. dxe5 fxe5 5. Bc4 Qe7 6. Ng5 Nh6 7. O-O Nc6 8. f4 exf4 9. Bxf4 Bd7 10. Nc3 O-O-O 11. Nd5 Qe8 12. b4 Ne5 13. Bxe5 Qxe5 14. b5 Qxg5 15. b6 axb6 16. a4 Bc6 17. a5 bxa5 18. Rxa5 Bxd5 19. Bxd5 c6 20. Bxc6 bxc6 21. Rxg5 g6 22. Qd3 Be7 23. Rg3 Kd7 24. Qa6 Rc8 25. Rh3 Ng4 26. Rhf3 Rc7 27. Rf7 h5 28. h3 Ne3 29. Rf1f3 Nxc2 30. e5 d5 31. Rf3f6 Re8 32. Rd6# 1-0`,

	//nolint:lll
	`[Event "Rated Classical game"]
[Site "https://lichess.org/jw1mk5dp"]
[White "vitellium"]
[Black "jmcd"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:43:55"]
[WhiteElo "1706"]
[BlackElo "1569"]
[WhiteRatingDiff "+31"]
[BlackRatingDiff "-7"]
[ECO "D10"]
[Opening "Slav Defense"]
[TimeControl "600+4"]
[Termination "Normal"]

1. d4 d5 2. c4 c6 3. e3 e6 4. a3 Nf6 5. Nc3 Be7 6. Nf3 O-O 7. Bd3 b6 8. Qc2 c5 9. O-O cxd4 10. Nxd4 e5 11. Nf5 e4 12. Nxe7+ Qxe7 13. Be2 Bf5 14. cxd5 Qe5 15. h3 Nxd5 16. Nxd5 Qxd5 17. Rd1 Qe5 18. b4 Qxa1 19. Bb2 Qa2 20. Bc4 Qxc4 21. Qxc4 Nd7 22. Qd4 Nf6 23. g4 Be6 24. g5 Nh5 25. h4 f6 26. Qxe4 Rae8 27. Qb7 Rf7 28. Qg2 fxg5 29. Qxg5 Rf5 30. Qg2 Ref8 31. f4 Bb3 32. Rd7 Bf7 33. Rxa7 Re8 34. Be5 Nf6 35. Bxf6 Rxf6 36. h5 Bxh5 37. Qxg7# 1-0`,

	//nolint:lll
	`[Event "Rated Blitz game"]
[Site "https://lichess.org/y7vlt0wb"]
[White "pablotorre"]
[Black "cipri4"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:46:56"]
[WhiteElo "1765"]
[BlackElo "1777"]
[WhiteRatingDiff "-11"]
[BlackRatingDiff "+10"]
[ECO "B06"]
[Opening "Robatsch (Modern) Defense"]
[TimeControl "180+0"]
[Termination "Normal"]

1. e4 g6 2. d4 Bg7 3. c4 b6 4. Nc3 Bb7 5. f4 d6 6. Nf3 e6 7. Bd3 h6 8. O-O Ne7 9. Qc2 O-O 10. Be3 f5 11. exf5 Nxf5 12. Rae1 Bxf3 13. Rxf3 Nxe3 14. Rfxe3 Bxd4 15. Bxg6 Bxe3+ 16. Rxe3 e5 17. Rg3 Kh8 18. Ne4 Nd7 19. f5 Nf6 20. Nf2 Rg8 21. Qe2 Qe7 22. Qe3 Rxg6 23. Rxg6 Ng8 24. Ng4 Qh4 25. f6 Rf8 26. Nxh6 Rxf6 27. Rxg8+ Kh7 28. Qd3+ e4 29. Rg4 Qf2+ 30. Kh1 Qe1+ 0-1`,

	`[Event "Rated Blitz game"]
[Site "https://lichess.org/u160ism9"]
[White "gonsar"]
[Black "isildur1"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:47:09"]
[WhiteElo "1500"]
[BlackElo "1486"]
[WhiteRatingDiff "+168"]
[BlackRatingDiff "-8"]
[ECO "C20"]
[Opening "English Opening: The Whale"]
[TimeControl "300+0"]
[Termination "Time forfeit"]

1. e4 e5 2. c4 f5 3. Nc3 Bb4 4. d3 Bxc3+ 5. bxc3 fxe4 6. Qh5+ g6 7. Qxe5+ Ne7 8. Qxh8+ 1-0`,

	//nolint:lll
	`[Event "Rated Blitz game"]
[Site "https://lichess.org/z1fng3zs"]
[White "rasmussenesq"]
[Black "barriosgb2"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:47:05"]
[WhiteElo "1686"]
[BlackElo "1750"]
[WhiteRatingDiff "+13"]
[BlackRatingDiff "-13"]
[ECO "A11"]
[Opening "English Opening: Caro-Kann Defensive System"]
[TimeControl "120+2"]
[Termination "Normal"]

1. c4 c6 2. b3 d5 3. Bb2 d4 4. d3 c5 5. g3 e5 6. Bg2 Be7 7. Nf3 Nf6 8. Nxe5 Nc6 9. Nxc6 bxc6 10. O-O Rb8 11. Re1 Bb7 12. e4 dxe3 13. Rxe3 O-O 14. Nd2 Bd6 15. Qc2 Bc8 16. Rae1 Bf5 17. Be4 Nxe4 18. dxe4 Bh3 19. Qc3 f6 20. Qc2 Be5 21. Nf3 Qe7 22. Bxe5 fxe5 23. Qc3 Qf6 24. Qxe5 Qg6 25. Qxc5 Bg4 26. e5 Qh5 27. Nh4 Qf7 28. Re4 g5 29. Ng2 Bf3 30. Re4e2 Bxe2 31. Rxe2 Qf3 32. Kf1 Rbd8 33. e6 Rd1+ 34. Re1 Rxe1+ 35. Nxe1 Qh1+ 36. Ke2 Qe4+ 37. Qe3 Qg4+ 38. Kf1 Qh3+ 39. Ng2 Qxh2 40. Qxg5+ Kh8 41. e7 Re8 42. Qf6+ Kg8 43. Nf4 Qh1+ 44. Ke2 Qe4+ 45. Kd2 Rxe7 46. Qg5+ Rg7 47. Qd8+ Kf7 48. Qd7+ Kf6 49. Nh5+ Kg5 50. Qxg7+ Kxh5 51. g4+ Qxg4 52. Qxg4+ Kxg4 53. Ke3 h5 54. Ke2 h4 55. Kf1 h3 56. Kg1 Kf3 57. a4 a6 58. b4 Ke4 59. b5 axb5 60. cxb5 c5 61. b6 c4 62. b7 c3 63. b8=Q c2 64. Qe8+ Kf3 65. Qc6+ Kg4 66. Qxc2 h2+ 67. Kxh2 1-0`,

	//nolint:lll
	`[Event "Rated Classical game"]
[Site "https://lichess.org/edfcqcr2"]
[White "heribert21"]
[Black "netsah08"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:47:19"]
[WhiteElo "1500"]
[BlackElo "1959"]
[WhiteRatingDiff "-40"]
[BlackRatingDiff "+5"]
[ECO "A45"]
[Opening "Indian Game"]
[TimeControl "600+10"]
[Termination "Normal"]

1. d4 Nf6 2. Bf4 g6 3. e3 Bg7 4. Nf3 d6 5. Be2 O-O 6. h3 Nbd7 7. c4 c6 8. Nc3 Re8 9. O-O e5 10. dxe5 dxe5 11. Bg3 Nh5 12. Bh2 f5 13. Qb3 e4 14. Ne1 Nhf6 15. Rd1 Qb6 16. Nc2 Nc5 17. Qxb6 axb6 18. Bd6 Be6 19. Bxc5 bxc5 20. Rd6 Nd7 21. Rfd1 Ne5 22. f4 Nxc4 23. Bxc4 Bxc4 24. b3 Bd3 25. Ne1 Bxc3 26. Nxd3 exd3 27. Rd6xd3 Bf6 28. a4 Rad8 29. Rxd8 Rxd8 30. Rxd8+ Bxd8 31. Kf2 Kf7 32. Kf3 Ke6 33. e4 fxe4+ 34. Kxe4 Ba5 35. g4 b5 36. axb5 cxb5 37. h4 c4 38. bxc4 bxc4 39. Kd4 c3 40. Kd3 Kd5 41. f5 gxf5 42. g5 f4 43. h5 Ke6 44. g6 hxg6 0-1`,

	`[Event "Rated Bullet game"]
[Site "https://lichess.org/l0ntar7q"]
[White "Arthzu"]
[Black "Panevis"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:47:49"]
[WhiteElo "1885"]
[BlackElo "2045"]
[WhiteRatingDiff "-6"]
[BlackRatingDiff "+8"]
[ECO "B06"]
[Opening "Modern Defense"]
[TimeControl "0+1"]
[Termination "Time forfeit"]

1. e4 g6 0-1`,

	//nolint:lll
	`[Event "Rated Blitz game"]
[Site "https://lichess.org/ubyu6pee"]
[White "Tortfeasor"]
[Black "Communist"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:47:57"]
[WhiteElo "1736"]
[BlackElo "1746"]
[WhiteRatingDiff "-27"]
[BlackRatingDiff "+10"]
[ECO "B36"]
[Opening "Sicilian Defense: Accelerated Dragon"]
[TimeControl "300+0"]
[Termination "Normal"]

1. e4 c5 2. Nf3 Nc6 3. d4 cxd4 4. Nxd4 g6 5. c3 Bg7 6. Nxc6 bxc6 7. Be3 Nf6 8. f3 Rb8 9. Qc2 Qa5 10. Nd2 Ba6 11. Nb3 Qa4 12. Bxa6 Qxa6 13. Kf2 O-O 14. Rhd1 d5 15. exd5 Nxd5 16. Bd4 Bxd4+ 17. Rxd4 Qb6 18. Ke2 e5 19. Rdd1 Qb5+ 20. Qd3 Nf4+ 0-1`,

	//nolint:lll
	`[Event "Rated Classical game"]
[Site "https://lichess.org/qlbh49a4"]
[White "olegis"]
[Black "The_Black_Rider"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:48:26"]
[WhiteElo "1560"]
[BlackElo "1541"]
[WhiteRatingDiff "-23"]
[BlackRatingDiff "+68"]
[ECO "B13"]
[Opening "Caro-Kann Defense: Exchange Variation"]
[TimeControl "420+8"]
[Termination "Normal"]

1. d4 c6 2. e4 d5 3. exd5 Qxd5 4. Nc3 Qa5 5. Bd2 Qc7 6. Qf3 Nf6 7. Bc4 e6 8. O-O-O b5 9. Bb3 Bb7 10. Qe3 Be7 11. f3 a5 12. a3 b4 13. Ne4 Nbd7 14. c3 bxa3 15. bxa3 c5 16. Kb2 Nxe4 17. fxe4 cxd4 18. cxd4 a4 19. Rc1 Qb6 20. Rc3 axb3 21. Rxb3 Qc6 22. Bb4 Bxb4 23. Rxb4 Qxe4 24. Qb3 Rb8 25. Nf3 Qe2+ 26. Qc2 Qxc2+ 27. Kxc2 Be4+ 0-1`,

	//nolint:lll
	`[Event "Rated Blitz game"]
[Site "https://lichess.org/oz6n3gdp"]
[White "airtsart"]
[Black "CHIQUITIN"]
[Result "1/2-1/2"]
[UTCDate "2013.01.31"]
[UTCTime "21:48:24"]
[WhiteElo "1863"]
[BlackElo "1758"]
[WhiteRatingDiff "-4"]
[BlackRatingDiff "+3"]
[ECO "C46"]
[Opening "Three Knights Opening"]
[TimeControl "180+0"]
[Termination "Normal"]

1. e4 e5 2. Nc3 Nc6 3. Nf3 Bc5 4. Nxe5 Nxe5 5. d4 Bxd4 6. Qxd4 d6 7. Bd3 c5 8. Bb5+ Bd7 9. Qxd6 Bxb5 10. Qxe5+ Qe7 11. Qxe7+ Nxe7 12. Nxb5 O-O 13. Be3 a6 14. Bxc5 axb5 15. Bxe7 Rfe8 16. Bc5 Rxe4+ 17. Be3 f5 18. O-O f4 19. Bd2 Rd8 20. Bc3 b4 21. Rad1 Rxd1 22. Rxd1 bxc3 23. bxc3 Re2 24. Rc1 Kf7 25. Kf1 Rd2 26. Ke1 Rd6 27. Rb1 b6 28. Rb4 Rc6 29. Rxf4+ Ke6 30. Rf3 g5 31. Kd2 g4 32. Rg3 h5 33. Re3+ Kd5 34. Rd3+ Kc4 35. Rd4+ Kb5 36. Kd3 Rf6 37. Rd5+ Ka4 38. Rxh5 Rxf2 39. h3 Rxg2 40. Rh4 Rg3+ 41. Kd2 Rxh3 42. Rxg4+ Ka3 43. Rg6 Rh2+ 44. Kd3 Rh3+ 45. Kc4 Rh4+ 46. Kb5 Kxa2 47. Rxb6 Kb2 48. c4 Kxc2 49. c5 Rh5 50. Rc6 Kb3 51. Kb6 Ka4 52. Rc8 Kb4 53. c6 Rb5+ 54. Kc7 Kc5 55. Kd7 Kb6 56. Rb8+ Kc5 57. Rxb5+ Kxb5 58. c7 Kb4 59. c8=Q Kb5 60. Qc6+ Kb4 61. Kd6 Ka3 62. Qc5+ Kb2 63. Qc4 Kb1 64. Qc3 Ka2 65. Kc5 Kb1 66. Kb4 Ka2 67. Ka4 Kb1 68. Kb4 Ka2 69. Kc4 Kb1 70. Kd3 Ka2 71. Kd2 Kb1 72. Kd1 Ka2 73. Qc2+ Ka3 74. Kd2 Kb4 75. Kd3 Kb5 76. Kd4 Kb6 77. Qc5+ Ka6 78. Kd5 Kb7 79. Qc6+ Ka7 80. Kc5 Kb8 81. Kb6 1/2-1/2`,
}

//nolint:decorder,lll
const pgnStr = `[Event "Rated Blitz game"]
[Site "https://lichess.org/l68uct88"]
[White "Alik"]
[Black "ioi"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:44:47"]
[WhiteElo "1534"]
[BlackElo "1299"]
[WhiteRatingDiff "-19"]
[BlackRatingDiff "+20"]
[ECO "D00"]
[Opening "Queen's Pawn Game: Mason Attack"]
[TimeControl "360+0"]
[Termination "Normal"]

1. d4 d5 2. Bf4 Nf6 3. e3 e6 4. Nc3 a6 5. Nh3 Ne4 6. Nxe4 dxe4
7. a3 Bd6 8. Bg5 f6 9. Qh5+ g6 10. Qh6 fxg5 11. Nxg5
Bf8 12. Qh4 Nc6 13. Bc4 b5 14. Bxe6 Be7 15.
Qxe4 Bxe6 16. Nxe6 Qd7 17. Nc5 Qf5 18. Qxc6+ Kf7 19. Qd5+ Qxd5
0-1

[Event "Rated Classical game"]
[Site "https://lichess.org/5psccdm0"]
[White "jajce3"]
[Black "jeh-1"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:44:31"]
[WhiteElo "1684"]
[BlackElo "1799"]
[WhiteRatingDiff "+15"]
[BlackRatingDiff "-17"]
[ECO "D10"]
[Opening "Slav Defense: Exchange Variation"]
[TimeControl "300+15"]
[Termination "Normal"]

1. d4 d5 2. c4 c6 3. cxd5 cxd5 4. Nc3 e6 5. a3 Nc6 6. e3 a6 7. b4 b5
8. Qb3 Bd7 9. g3 Nf6 10. Bg2 Bd6 11. Nge2 O-O 12. O-O Re8 13. Bd2 Qb6 14. Na2 e5 15. Bc3 e4 16. Bd2 a5
17. Nec3 axb4 18. Nxd5 Nxd5 19. Qxd5 Ne7 20. Qxe4 bxa3 21. Bb4 Nc6 22. Qb1 Rac8 23. Bxd6 Na5 24. Bc5 Qa6 25. Nb4 Qf6 26. Nd5 Qd8 27.
Bb6 Qg5 28. Bxa5 Be6 29. Nb6 Rb8 30. Rxa3 Bc4 31. Nxc4 bxc4 32. Qe1 Rec8 33. Rc3 Qxa5 34. Qc1 Qa4 35. Bd5 Rb4 36. Qc2 Qxc2 37. Rxc2 c3 38.
Ra1 Rbb8 39. Kf1 g6 40. Ke2 Rb2 41. Rxb2 cxb2 42. Rb1 Rc2+ 43. Kd1 Rxf2 44. h3 Kg7 45. Bb3 Rh2 46. h4 Rg2 47. Bc2 Rxg3 48. e4 Rh3 49. d5 Rxh4 50. Rxb2 h5 51. d6 Kf6 52. d7 Ke7 53. Rb7 Rh1+ 54. Kd2 Ra1 55. Bd3 Ra8 56. Bb5 f6 57. Rc7 Kd8 58.
Rc6 Rb8 59. Ba4 g5 60. Rxf6 Ke7 61. Rh6 Rb4 62. Bc6 Rd4+ 63. Ke3 Rc4 64. Rh8 Rxc6 65. d8=Q+ 1-0

[Event "Rated Bullet game"]
[Site "https://lichess.org/0s6fmzyx"]
[White "SkyBang"]
[Black "KAINVALDOMERO"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:10"]
[WhiteElo "1682"]
[BlackElo "1242"]
[WhiteRatingDiff "-145"]
[BlackRatingDiff "+186"]
[ECO "C11"]
[Opening "French Defense: Classical Variation, Steinitz Variation"]
[TimeControl "120+0"]
[Termination "Normal"]

1. d4 d5 2. Nc3 e6 3. e4 Nf6 4. e5 Nfd7 5. Nf3 f6 6. exf6 Nxf6 7. Bg5 Be7 8. Ne5 h6 9. Bxf6 Bxf6 10.
Qd3 O-O 11. Ng6 Re8 12. Be2 Nc6 13. Bh5 Bxd4 14. Nf4 Qg5 15. g3 e5 16. Bxe8 exf4 17. O-O-O fxg3+
18. Kb1 Bxc3 19. bxc3 gxf2 20. h4 Qe7 21. Bxc6 bxc6 22. Qf3 Rb8+ 23. Ka1 Be6 24. Qxf2 Qa3 25. Rb1 d4 26. Rb3 Rxb3 27. cxb3 Bxb3 28. Qxd4 Qxa2# 0-1

[Event "Rated Blitz game"]
[Site "https://lichess.org/1qx160x0"]
[White "Artem84"]
[Black "zine1971"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:05"]
[WhiteElo "1645"]
[BlackElo "1437"]
[WhiteRatingDiff "-16"]
[BlackRatingDiff "+44"]
[ECO "D00"]
[Opening "Queen's Pawn Game #2"]
[TimeControl "300+0"]
[Termination "Normal"]

1. d4 d5 2. e3 e6 3. Nf3 h6 4. c4 c6 5. Nc3 Bb4 6. Bd2 Nf6 7. a3 Bxc3 8. Bxc3 Ne4 9. Qc2 O-O 10. Nd2 Nd6 11. c5 Nf5 12. Bd3 Ne7 13. O-O f5 14. b4 Ng6 15. Nf3 Nh4 16. Nxh4 Qxh4 17. g3 Qh3 18. Rfe1 Nd7 19. Bf1 Qh5 20. Qe2 Qg6 21. Bg2 Nf6 22. b5 Bd7 23. a4 Rae8 24. Bf1 Ne4 25. Qc2 h5 26. bxc6 bxc6 27. Rab1 h4 28. Rb7 hxg3 29. fxg3 Bc8 30. Rxa7 Nxg3 31. hxg3 Qxg3+ 32. Qg2 Qh4 33. Qh2 Qg5+ 34. Qg2 Qf6 35. Re2 g5 36. Qh2 Rf7 37. Rxf7 Qxf7 38. Qe5 Qh5 39. Rh2 Qg4+ 40. Rg2 Qh4 41. Rh2 Qg4+ 42. Bg2 Qd1+ 43. Bf1 Qg4+ 44. Rg2 Qh4 45. Qf6 g4 46. Rh2 Qg3+ 47. Rg2 Qxe3+ 48. Kh1 Qh3+ 49. Rh2 Qxf1# 0-1

[Event "Rated Blitz tournament https://lichess.org/tournament/wlvsn5j1"]
[Site "https://lichess.org/2wf8ta39"]
[White "neymardossantos"]
[Black "andresneno"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:18"]
[WhiteElo "1635"]
[BlackElo "1718"]
[WhiteRatingDiff "-9"]
[BlackRatingDiff "+9"]
[ECO "C23"]
[Opening "Bishop's Opening"]
[TimeControl "300+1"]
[Termination "Normal"]

1. e4 e5 2. Bc4 Nc6 3. Nc3 Bc5 4. Nf3 d6 5. d3 h6 6. a3 a6 7. Be3 Nd4 8. b4 Ba7 9. O-O Bg4 10. h3 Bxf3 11. gxf3 Qf6 12. Ne2 Nxf3+ 13. Kh1 Qh4 14. Ng1 Nxg1 15. Kxg1 Qxh3 16. Bxa7 Rxa7 17. f4 Qg3+ 18. Kh1 exf4 19. Qe2 Nf6 20. Rf2 Ng4 21. Rg2 Qh3+ 22. Kg1 f3 0-1

[Event "Rated Blitz tournament https://lichess.org/tournament/wlvsn5j1"]
[Site "https://lichess.org/wzfkda1w"]
[White "Federico"]
[Black "AqibNabi01"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:18"]
[WhiteElo "1874"]
[BlackElo "1433"]
[WhiteRatingDiff "+2"]
[BlackRatingDiff "-12"]
[ECO "C00"]
[Opening "French Defense: Normal Variation"]
[TimeControl "300+1"]
[Termination "Normal"]

1. d4 e6 2. e4 Qh4 3. Nc3 Nf6 4. Bd3 Ng4 5. g3 Qf6 6. Qxg4 e5 7. dxe5 Qxe5 8. Nf3 d6 9. Qxc8+ Ke7 10. Qxc7+ Ke6 11. Nxe5 dxe5 1-0

[Event "Rated Blitz game"]
[Site "https://lichess.org/m91p5vcs"]
[White "macface"]
[Black "deadpawn"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:25"]
[WhiteElo "1703"]
[BlackElo "1369"]
[WhiteRatingDiff "+3"]
[BlackRatingDiff "-21"]
[ECO "A43"]
[Opening "Benoni Defense: Benoni Gambit Accepted"]
[TimeControl "300+0"]
[Termination "Time forfeit"]

1. d4 c5 2. dxc5 Qa5+ 3. c3 Qxc5 4. g3 Nc6 5. Bg2 Nf6 6. e3 g6 7. Ne2 Bg7 8. O-O h5 9. Nd2 g5 10. Nd4 d6 11. Nd2b3 Qb6 12. Nxc6 bxc6 13. Nd4 Rb8 14. Bxc6+ Bd7 15. Bg2 Bg4 16. f3 Bd7 17. Kh1 g4 18. f4 h4 19. gxh4 Rxh4 20. Qe1 Rh5 21. e4 Rh7 22. f5 Bh6 23. Bxh6 Rxh6 24. Qg3 Qxb2 25. Rab1 Qxc3 26. Rxb8+ Bc8 27. Rxc8+ Qxc8 28. e5 dxe5 29. Qxe5 Qc3 30. Qb8+ Kd7 31. Bc6+ Qxc6+ 32. Nxc6 Kxc6 33. Qxa7 Ne4 34. Qa8+ Kc5 35. Qxe4 Rd6 36. Qxg4 1-0

[Event "Rated Classical game"]
[Site "https://lichess.org/cu3hedz6"]
[White "SAIRAM"]
[Black "muratk"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:52"]
[WhiteElo "1278"]
[BlackElo "1505"]
[WhiteRatingDiff "-5"]
[BlackRatingDiff "+5"]
[ECO "A00"]
[Opening "Van't Kruijs Opening"]
[TimeControl "300+5"]
[Termination "Normal"]

1. e3 d5 2. a3 e5 3. b4 Nc6 4. Bb2 d4 5. exd4 exd4 6. Bb5 Qe7+ 7. Ne2 Kd8 8. Bxc6 bxc6 9. Bxd4 Ba6 10. d3 Nf6 11. Bxf6 gxf6 12. O-O Rg8 13. g3 Bb7 14. Nbc3 c5 15. bxc5 Qxc5 16. Na4 Qc6 17. c4 Qg2# 0-1

[Event "Rated Blitz game"]
[Site "https://lichess.org/mm1fckgq"]
[White "mchelken"]
[Black "rasmussenesq"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:45:58"]
[WhiteElo "1618"]
[BlackElo "1677"]
[WhiteRatingDiff "-9"]
[BlackRatingDiff "+9"]
[ECO "B40"]
[Opening "Sicilian Defense: French Variation"]
[TimeControl "180+0"]
[Termination "Normal"]

1. e4 c5 2. Nf3 e6 3. Bc4 a6 4. Nc3 b5 5. Bb3 c4 6. Nxb5 axb5 7. Bxc4 bxc4 8. d4 cxd3 9. Qxd3 Ba6 10. Qa3 Bxa3 0-1

[Event "Rated Bullet game"]
[Site "https://lichess.org/io1kkxaj"]
[White "Panevis"]
[Black "Arthzu"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:46:22"]
[WhiteElo "2036"]
[BlackElo "1892"]
[WhiteRatingDiff "+9"]
[BlackRatingDiff "-7"]
[ECO "B10"]
[Opening "Caro-Kann Defense"]
[TimeControl "0+1"]
[Termination "Normal"]

1. e4 c6 2. Nf3 d6 3. d4 g6 4. Bd3 Bg7 5. O-O Nd7 6. Re1 Ngf6 7. c3 e5 8. Bg5 h6 9. Bh4 g5 10. Bg3 O-O 11. Nbd2 Nh5 12. Qe2 Nf4 13. Bxf4 gxf4 14. g3 Qf6 15. Nc4 Kh8 16. Kh1 Rg8 17. dxe5 Nxe5 18. Ncxe5 dxe5 19. Bc4 Bg4 20. Rf1 fxg3 21. fxg3 Bxf3+ 22. Rxf3 Qg5 23. Raf1 f6 24. Bxg8 Kxg8 25. Rf5 Qg6 26. Qf3 Kh7 27. Rf2 Rd8 28. h4 Rd1+ 29. Qxd1 1-0

[Event "Rated Bullet game"]
[Site "https://lichess.org/qpwvedsb"]
[White "dvorak"]
[Black "nichiren1967"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:46:55"]
[WhiteElo "1559"]
[BlackElo "1770"]
[WhiteRatingDiff "+19"]
[BlackRatingDiff "-17"]
[ECO "C41"]
[Opening "Philidor Defense"]
[TimeControl "120+0"]
[Termination "Normal"]

1. e4 e5 2. Nf3 d6 3. d4 f6 4. dxe5 fxe5 5. Bc4 Qe7 6. Ng5 Nh6 7. O-O Nc6 8. f4 exf4 9. Bxf4 Bd7 10. Nc3 O-O-O 11. Nd5 Qe8 12. b4 Ne5 13. Bxe5 Qxe5 14. b5 Qxg5 15. b6 axb6 16. a4 Bc6 17. a5 bxa5 18. Rxa5 Bxd5 19. Bxd5 c6 20. Bxc6 bxc6 21. Rxg5 g6 22. Qd3 Be7 23. Rg3 Kd7 24. Qa6 Rc8 25. Rh3 Ng4 26. Rhf3 Rc7 27. Rf7 h5 28. h3 Ne3 29. Rf1f3 Nxc2 30. e5 d5 31. Rf3f6 Re8 32. Rd6# 1-0

[Event "Rated Classical game"]
[Site "https://lichess.org/jw1mk5dp"]
[White "vitellium"]
[Black "jmcd"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:43:55"]
[WhiteElo "1706"]
[BlackElo "1569"]
[WhiteRatingDiff "+31"]
[BlackRatingDiff "-7"]
[ECO "D10"]
[Opening "Slav Defense"]
[TimeControl "600+4"]
[Termination "Normal"]

1. d4 d5 2. c4 c6 3. e3 e6 4. a3 Nf6 5. Nc3 Be7 6. Nf3 O-O 7. Bd3 b6 8. Qc2 c5 9. O-O cxd4 10. Nxd4 e5 11. Nf5 e4 12. Nxe7+ Qxe7 13. Be2 Bf5 14. cxd5 Qe5 15. h3 Nxd5 16. Nxd5 Qxd5 17. Rd1 Qe5 18. b4 Qxa1 19. Bb2 Qa2 20. Bc4 Qxc4 21. Qxc4 Nd7 22. Qd4 Nf6 23. g4 Be6 24. g5 Nh5 25. h4 f6 26. Qxe4 Rae8 27. Qb7 Rf7 28. Qg2 fxg5 29. Qxg5 Rf5 30. Qg2 Ref8 31. f4 Bb3 32. Rd7 Bf7 33. Rxa7 Re8 34. Be5 Nf6 35. Bxf6 Rxf6 36. h5 Bxh5 37. Qxg7# 1-0

[Event "Rated Blitz game"]
[Site "https://lichess.org/y7vlt0wb"]
[White "pablotorre"]
[Black "cipri4"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:46:56"]
[WhiteElo "1765"]
[BlackElo "1777"]
[WhiteRatingDiff "-11"]
[BlackRatingDiff "+10"]
[ECO "B06"]
[Opening "Robatsch (Modern) Defense"]
[TimeControl "180+0"]
[Termination "Normal"]

1. e4 g6 2. d4 Bg7 3. c4 b6 4. Nc3 Bb7 5. f4 d6 6. Nf3 e6 7. Bd3 h6 8. O-O Ne7 9. Qc2 O-O 10. Be3 f5 11. exf5 Nxf5 12. Rae1 Bxf3 13. Rxf3 Nxe3 14. Rfxe3 Bxd4 15. Bxg6 Bxe3+ 16. Rxe3 e5 17. Rg3 Kh8 18. Ne4 Nd7 19. f5 Nf6 20. Nf2 Rg8 21. Qe2 Qe7 22. Qe3 Rxg6 23. Rxg6 Ng8 24. Ng4 Qh4 25. f6 Rf8 26. Nxh6 Rxf6 27. Rxg8+ Kh7 28. Qd3+ e4 29. Rg4 Qf2+ 30. Kh1 Qe1+ 0-1

[Event "Rated Blitz game"]
[Site "https://lichess.org/u160ism9"]
[White "gonsar"]
[Black "isildur1"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:47:09"]
[WhiteElo "1500"]
[BlackElo "1486"]
[WhiteRatingDiff "+168"]
[BlackRatingDiff "-8"]
[ECO "C20"]
[Opening "English Opening: The Whale"]
[TimeControl "300+0"]
[Termination "Time forfeit"]

1. e4 e5 2. c4 f5 3. Nc3 Bb4 4. d3 Bxc3+ 5. bxc3 fxe4 6. Qh5+ g6 7. Qxe5+ Ne7 8. Qxh8+ 1-0

[Event "Rated Blitz game"]
[Site "https://lichess.org/z1fng3zs"]
[White "rasmussenesq"]
[Black "barriosgb2"]
[Result "1-0"]
[UTCDate "2013.01.31"]
[UTCTime "21:47:05"]
[WhiteElo "1686"]
[BlackElo "1750"]
[WhiteRatingDiff "+13"]
[BlackRatingDiff "-13"]
[ECO "A11"]
[Opening "English Opening: Caro-Kann Defensive System"]
[TimeControl "120+2"]
[Termination "Normal"]

1. c4 c6 2. b3 d5 3. Bb2 d4 4. d3 c5 5. g3 e5 6. Bg2 Be7 7. Nf3 Nf6 8. Nxe5 Nc6 9. Nxc6 bxc6 10. O-O Rb8 11. Re1 Bb7 12. e4 dxe3 13. Rxe3 O-O 14. Nd2 Bd6 15. Qc2 Bc8 16. Rae1 Bf5 17. Be4 Nxe4 18. dxe4 Bh3 19. Qc3 f6 20. Qc2 Be5 21. Nf3 Qe7 22. Bxe5 fxe5 23. Qc3 Qf6 24. Qxe5 Qg6 25. Qxc5 Bg4 26. e5 Qh5 27. Nh4 Qf7 28. Re4 g5 29. Ng2 Bf3 30. Re4e2 Bxe2 31. Rxe2 Qf3 32. Kf1 Rbd8 33. e6 Rd1+ 34. Re1 Rxe1+ 35. Nxe1 Qh1+ 36. Ke2 Qe4+ 37. Qe3 Qg4+ 38. Kf1 Qh3+ 39. Ng2 Qxh2 40. Qxg5+ Kh8 41. e7 Re8 42. Qf6+ Kg8 43. Nf4 Qh1+ 44. Ke2 Qe4+ 45. Kd2 Rxe7 46. Qg5+ Rg7 47. Qd8+ Kf7 48. Qd7+ Kf6 49. Nh5+ Kg5 50. Qxg7+ Kxh5 51. g4+ Qxg4 52. Qxg4+ Kxg4 53. Ke3 h5 54. Ke2 h4 55. Kf1 h3 56. Kg1 Kf3 57. a4 a6 58. b4 Ke4 59. b5 axb5 60. cxb5 c5 61. b6 c4 62. b7 c3 63. b8=Q c2 64. Qe8+ Kf3 65. Qc6+ Kg4 66. Qxc2 h2+ 67. Kxh2 1-0

[Event "Rated Classical game"]
[Site "https://lichess.org/edfcqcr2"]
[White "heribert21"]
[Black "netsah08"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:47:19"]
[WhiteElo "1500"]
[BlackElo "1959"]
[WhiteRatingDiff "-40"]
[BlackRatingDiff "+5"]
[ECO "A45"]
[Opening "Indian Game"]
[TimeControl "600+10"]
[Termination "Normal"]

1. d4 Nf6 2. Bf4 g6 3. e3 Bg7 4. Nf3 d6 5. Be2 O-O 6. h3 Nbd7 7. c4 c6 8. Nc3 Re8 9. O-O e5 10. dxe5 dxe5 11. Bg3 Nh5 12. Bh2 f5 13. Qb3 e4 14. Ne1 Nhf6 15. Rd1 Qb6 16. Nc2 Nc5 17. Qxb6 axb6 18. Bd6 Be6 19. Bxc5 bxc5 20. Rd6 Nd7 21. Rfd1 Ne5 22. f4 Nxc4 23. Bxc4 Bxc4 24. b3 Bd3 25. Ne1 Bxc3 26. Nxd3 exd3 27. Rd6xd3 Bf6 28. a4 Rad8 29. Rxd8 Rxd8 30. Rxd8+ Bxd8 31. Kf2 Kf7 32. Kf3 Ke6 33. e4 fxe4+ 34. Kxe4 Ba5 35. g4 b5 36. axb5 cxb5 37. h4 c4 38. bxc4 bxc4 39. Kd4 c3 40. Kd3 Kd5 41. f5 gxf5 42. g5 f4 43. h5 Ke6 44. g6 hxg6 0-1

[Event "Rated Bullet game"]
[Site "https://lichess.org/l0ntar7q"]
[White "Arthzu"]
[Black "Panevis"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:47:49"]
[WhiteElo "1885"]
[BlackElo "2045"]
[WhiteRatingDiff "-6"]
[BlackRatingDiff "+8"]
[ECO "B06"]
[Opening "Modern Defense"]
[TimeControl "0+1"]
[Termination "Time forfeit"]

1. e4 g6 0-1

[Event "Rated Blitz game"]
[Site "https://lichess.org/ubyu6pee"]
[White "Tortfeasor"]
[Black "Communist"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:47:57"]
[WhiteElo "1736"]
[BlackElo "1746"]
[WhiteRatingDiff "-27"]
[BlackRatingDiff "+10"]
[ECO "B36"]
[Opening "Sicilian Defense: Accelerated Dragon"]
[TimeControl "300+0"]
[Termination "Normal"]

1. e4 c5 2. Nf3 Nc6 3. d4 cxd4 4. Nxd4 g6 5. c3 Bg7 6. Nxc6 bxc6 7. Be3 Nf6 8. f3 Rb8 9. Qc2 Qa5 10. Nd2 Ba6 11. Nb3 Qa4 12. Bxa6 Qxa6 13. Kf2 O-O 14. Rhd1 d5 15. exd5 Nxd5 16. Bd4 Bxd4+ 17. Rxd4 Qb6 18. Ke2 e5 19. Rdd1 Qb5+ 20. Qd3 Nf4+ 0-1

[Event "Rated Classical game"]
[Site "https://lichess.org/qlbh49a4"]
[White "olegis"]
[Black "The_Black_Rider"]
[Result "0-1"]
[UTCDate "2013.01.31"]
[UTCTime "21:48:26"]
[WhiteElo "1560"]
[BlackElo "1541"]
[WhiteRatingDiff "-23"]
[BlackRatingDiff "+68"]
[ECO "B13"]
[Opening "Caro-Kann Defense: Exchange Variation"]
[TimeControl "420+8"]
[Termination "Normal"]

1. d4 c6 2. e4 d5 3. exd5 Qxd5 4. Nc3 Qa5 5. Bd2 Qc7 6. Qf3 Nf6 7. Bc4 e6 8. O-O-O b5 9. Bb3 Bb7 10. Qe3 Be7 11. f3 a5 12. a3 b4 13. Ne4 Nbd7 14. c3 bxa3 15. bxa3 c5 16. Kb2 Nxe4 17. fxe4 cxd4 18. cxd4 a4 19. Rc1 Qb6 20. Rc3 axb3 21. Rxb3 Qc6 22. Bb4 Bxb4 23. Rxb4 Qxe4 24. Qb3 Rb8 25. Nf3 Qe2+ 26. Qc2 Qxc2+ 27. Kxc2 Be4+ 0-1

[Event "Rated Blitz game"]
[Site "https://lichess.org/oz6n3gdp"]
[White "airtsart"]
[Black "CHIQUITIN"]
[Result "1/2-1/2"]
[UTCDate "2013.01.31"]
[UTCTime "21:48:24"]
[WhiteElo "1863"]
[BlackElo "1758"]
[WhiteRatingDiff "-4"]
[BlackRatingDiff "+3"]
[ECO "C46"]
[Opening "Three Knights Opening"]
[TimeControl "180+0"]
[Termination "Normal"]

1. e4 e5 2. Nc3 Nc6 3. Nf3 Bc5 4. Nxe5 Nxe5 5. d4 Bxd4 6. Qxd4 d6 7. Bd3 c5 8. Bb5+ Bd7 9. Qxd6 Bxb5 10. Qxe5+ Qe7 11. Qxe7+ Nxe7 12. Nxb5 O-O 13. Be3 a6 14. Bxc5 axb5 15. Bxe7 Rfe8 16. Bc5 Rxe4+ 17. Be3 f5 18. O-O f4 19. Bd2 Rd8 20. Bc3 b4 21. Rad1 Rxd1 22. Rxd1 bxc3 23. bxc3 Re2 24. Rc1 Kf7 25. Kf1 Rd2 26. Ke1 Rd6 27. Rb1 b6 28. Rb4 Rc6 29. Rxf4+ Ke6 30. Rf3 g5 31. Kd2 g4 32. Rg3 h5 33. Re3+ Kd5 34. Rd3+ Kc4 35. Rd4+ Kb5 36. Kd3 Rf6 37. Rd5+ Ka4 38. Rxh5 Rxf2 39. h3 Rxg2 40. Rh4 Rg3+ 41. Kd2 Rxh3 42. Rxg4+ Ka3 43. Rg6 Rh2+ 44. Kd3 Rh3+ 45. Kc4 Rh4+ 46. Kb5 Kxa2 47. Rxb6 Kb2 48. c4 Kxc2 49. c5 Rh5 50. Rc6 Kb3 51. Kb6 Ka4 52. Rc8 Kb4 53. c6 Rb5+ 54. Kc7 Kc5 55. Kd7 Kb6 56. Rb8+ Kc5 57. Rxb5+ Kxb5 58. c7 Kb4 59. c8=Q Kb5 60. Qc6+ Kb4 61. Kd6 Ka3 62. Qc5+ Kb2 63. Qc4 Kb1 64. Qc3 Ka2 65. Kc5 Kb1 66. Kb4 Ka2 67. Ka4 Kb1 68. Kb4 Ka2 69. Kc4 Kb1 70. Kd3 Ka2 71. Kd2 Kb1 72. Kd1 Ka2 73. Qc2+ Ka3 74. Kd2 Kb4 75. Kd3 Kb5 76. Kd4 Kb6 77. Qc5+ Ka6 78. Kd5 Kb7 79. Qc6+ Ka7 80. Kc5 Kb8 81. Kb6 1/2-1/2

`
