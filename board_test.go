package standardchess_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_board_MakeMove_CaptureAddsToCapturedPieces(t *testing.T) {
	b := standardtest.DecodeFEN("r3k2r/ppp2ppp/B1n2n1B/3pp2Q/3PP2q/b1N2N1b/PPP2PPP/R3K2R")
	require.Equal(t, 0, len(b.CapturedPieces()))
	cases := []struct {
		move          string
		expectCapture bool
	}{
		{"Qh4", true},
		{"Bb2", true},
		{"Nd5", true},
		{"0-0", false},
	}

	capturedCount := 0
	for _, c := range cases {
		res, err := b.MakeMove(chess.StringMove(c.move))
		require.NoErrorf(t, err, "MakeMove failed (%s)", c.move)

		if !c.expectCapture {
			require.Nilf(
				t,
				res.CapturedPiece(),
				"expected no capture for move %s, but CapturedPiece is %v",
				c.move,
				res.CapturedPiece(),
			)
			require.Equal(t, capturedCount, len(b.CapturedPieces()))

			continue
		}

		require.NotNilf(
			t,
			res.CapturedPiece(),
			"expected move %s to capture a piece, but CapturedPiece is nil",
			c.move,
		)
		capturedCount++

		require.Equal(t, capturedCount, len(b.CapturedPieces()))
		require.Equalf(
			t,
			b.CapturedPieces()[capturedCount-1],
			res.CapturedPiece(),
			"after %s captured piece in board does not match move result",
			c.move,
		)
	}
}

func TestMakeMoveWithUndo(t *testing.T) {
	tests := []struct {
		name    string
		pgnStr  string
		wantFEN string
	}{
		{
			"https://www.chess.com/games/view/14842105",
			`1. e4 c6 2. d4 d5 3. e5 Bf5 4. Nc3 e6 5. g4 Bg6 6. Nge2 c5 7. Be3 Ne7 8. f4 h5
9. f5 exf5 10. g5 Nbc6 11. Nf4 a6 12. Bg2 cxd4 13. Bxd4 Nxd4 14. Qxd4 Nc6 15.
Qf2 Bb4 16. O-O-O Bxc3 17. bxc3 Qa5 18. Rxd5 Qxc3 19. Qc5 Qxc5 20. Rxc5 O-O 21.
Bxc6 bxc6 22. Rd1 Rab8 23. c4 Rfd8 24. Rd6 Kf8 25. Rcxc6 Rdc8 26. Kc2 h4 27.
Rxc8+ Rxc8 28. Kc3 a5 29. Ra6 Rb8 30. Rxa5 Rb1 31. c5 Re1 32. Ra8+ Ke7 33. Ra7+
Ke8 34. Nd3 Re3 35. Kd2 Rh3 36. c6 Rxh2+ 37. Ke3 Rc2 38. e6 h3 39. Nb4 f4+ 40.
Kd4 h2 41. Ra8+ Ke7 42. Rh8 Rd2+ 43. Kc5 Be4 44. c7 Bb7 45. Kb6 Bc8 46. Rxc8
h1=Q 47. Re8+ Kxe8 48. c8=Q+ Ke7 49. Nc6+ Qxc6+ 50. Qxc6 Rd6 0-1`,
			"8/4kpp1/1KQrP3/6P1/5p2/8/P7/8 w - - 1 51",
		},
		{
			"https://www.chess.com/analysis/collection/chess-com-immortal-games-YAv9iF5p/PhnXZTvAi/analysis",
			`1. e4 d6 2. d4 Nf6 3. Nc3 g6 4. f4 Bg7 5. a3 O-O 6. Nf3 Bg4 7. h3 Bxf3 8. Qxf3
Nc6 9. Be3 e5 10. dxe5 dxe5 11. f5 Nd4 12. Qf2 Nd7 13. g4 c6 14. O-O-O b5 15. g5
f6 16. h4 a5 17. h5 fxg5 18. hxg6 h6 19. Ne2 c5 20. Nc3 b4 21. Bc4+ Kh8 22.
Rxh6+ Bxh6 23. Qh2 Kg7 24. Qxh6+ Kxh6 25. Rh1+ 1-0`,
			"r2q1r2/3n4/6Pk/p1p1pPp1/1pBnP3/P1N1B3/1PP5/2K4R b - - 1 25",
		},
		{
			"https://www.chess.com/game/live/138534055125",
			`1. e4 e5 2. Bb5 c6 3. Nc3 cxb5 4. Nxb5 Nf6 5. f3 d5 6. exd5 Nxd5 7. c4 a6 8.
Nd6+ Bxd6 9. cxd5 Bf5 10. d4 exd4 11. Qe2+ Qe7 12. Qxe7+ Bxe7 13. Bg5 Bxg5 14.
Nh3 Bh4+ 15. g3 Bg5 16. Nxg5 O-O 17. b4 Nc6 18. b5 Ne5 19. b6 Nxf3+ 20. Kf2 Nxg5
21. Rac1 Ne4+ 22. Kf3 Nd2+ 23. Kf4 Bb1 24. Rxb1 Nxb1 25. Rxb1 h6 26. g4 g5+ 27.
Kf3 d3 28. Rd1 Rad8 29. Rxd3 a5 30. d6 Rfe8 31. d7 Re1 32. a3 Rh1 33. h4 Rxh4
34. Rd6 h5 35. gxh5 Rxh5 36. Kg4 Rh4+ 37. Kxg5 Ra4 38. Kf6 Rxa3 39. Re6 Kf8 40.
Re7 Rf3+ 41. Ke5 a4 42. Re8+ Rxe8+ 43. dxe8=Q+ Kxe8 44. Kd6 a3 45. Kc7 Rb3 46.
Kxb7 a2 47. Kc8 a1=Q 48. b7 Qc3+ 49. Kb8 Qe5+ 50. Ka7 Qd5 51. b8=Q+ Rxb8 52.
Kxb8 Qc6 53. Ka7 Kd7 54. Kb8 Qc7+ 55. Ka8 Kc6 1/2-1/2`,
			"K7/2q2p2/2k5/8/8/8/8/8 w - - 7 56",
		},
		{
			"https://lichess.org/hND1G90l",
			`1. d4 Nf6 2. Nc3 c5 3. d5 g6 4. e4 Bg7 5. e5 Ng8 6. f4 d6 7. Nf3 Bg4 8. Bb5+ Kf8 9. e6 fxe6
10. Ng5 Qa5 11. Qxg4 Bxc3+ 12. bxc3 Qxc3+ 13. Bd2 Qxa1+ 14. Ke2 Qxh1 15. Qxe6 Qxg2+ 16. Kd3 Nh6
17. Qc8+ Kg7 18. Bc3+ e5 19. dxe6# 1-0`,
			"rnQ4r/pp4kp/3pP1pn/1Bp3N1/5P2/2BK4/P1P3qP/8 b - - 0 19",
		},
		{
			"https://www.chess.com/game/live/122246752662",
			`1. c4 g5 2. Nc3 g4 3. g3 e5 4. c5 d5 5. Qa4+ c6 6. e3 b5 7. Qd1 Bxc5 8. Be2 d4
9. Nb1 dxe3 10. dxe3 Qg5 11. Bd2 Nf6 12. b4 Bd6 13. Bd3 e4 14. Be2 Na6 15. Nc3
Nxb4 16. h4 gxh3 17. Nxh3 Qh6 18. O-O Bxh3 19. Re1 Ng4 20. Bxg4 Bxg4 21. Nxe4
Bxd1 22. Bxb4 Bf3 23. Nxd6+ Kd7 24. Red1 Qh1# 0-1`,
			"r6r/p2k1p1p/2pN4/1p6/1B6/4PbP1/P4P2/R2R2Kq w - - 3 25",
		},
		{
			"https://www.chess.com/game/live/104676098479",
			`1. d4 Nc6 2. Nf3 e5 3. dxe5 g5 4. Nxg5 Nxe5 5. e3 h6 6. Nf3 Bb4+ 7. Bd2 d6 8.
Bxb4 Nxf3+ 9. gxf3 Nf6 10. c3 Nd5 11. Ba3 Rg8 12. Bb5+ Bd7 13. Bxd7+ Kxd7 14. e4
Nf4 15. c4 Ne6 16. Qd5 Rg5 17. Qd3 Qf6 18. Nc3 Nd4 19. O-O-O Nxf3 20. c5 Ne5 21.
Qb5+ c6 22. Qf1 Ke7 23. Rg1 b5 24. cxb6 axb6 25. Rxd6 Qxd6 26. Bxd6+ Kxd6 27.
Qd1+ Ke7 28. Na4 Rxa4 29. Qxa4 Rxg1+ 30. Kc2 b5 31. Qa7+ Nd7 32. f3 h5 33. Qxg1
Nc5 34. Qxc5+ Kf6 35. Qxh5 b4 36. f4 c5 37. e5+ Ke6 38. Kb3 Kd5 39. a3 Kd4 40.
axb4 cxb4 41. Qf5 Ke3 42. Kxb4 Kf2 43. Ka4 Kg2 44. b3 Kxh2 45. b4 Kg1 46. b5 Kf2
47. Ka5 Ke3 48. b6 Kd4 49. Ka6 Kc5 50. b7 Kc6 51. Ka7 Kb5 52. b8=Q+ Kc6 53. Qxf7
Kc5 54. Qd8 Kc6 55. Qfc7+ Kb5 56. Qb6+ Kc4 57. Qc6+ Kb4 58. Qdd5 Ka3 59. Qb6 Ka4
60. Qda5# 1-0`,
			"8/K7/1Q6/Q3P3/k4P2/8/8/8 b - - 14 60"},
	}

	initFEN := standardtest.EncodeFEN(standardchess.NewBoard())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := standardchess.NewBoard()

			for _, move := range standardtest.MovesFromPGN(tt.pgnStr) {
				result, err := board.MakeMove(move)
				require.NotNil(t, result)
				require.NoError(t, err)
			}

			require.Equal(t, tt.wantFEN, standardtest.EncodeFEN(board))

			for i := range board.MoveHistory() {
				result, err := board.UndoLastMove()
				require.NoErrorf(t, err, "No %d", i+1)
				require.NotNil(t, result, "No %d", i+1)
			}

			afterFEN := standardtest.EncodeFEN(board)
			assert.Equal(t, initFEN, afterFEN)
		})
	}
}

func BenchmarkNewBoard(b *testing.B) {
	for range b.N {
		standardchess.NewBoard()
	}
}

func BenchmarkNewBoardFromMoves(b *testing.B) {
	moves := standardtest.MovesFromPGN(
		`1. e4 e5 2. Bb5 c6 3. Nc3 cxb5 4. Nxb5 Nf6 5. f3 d5 6. exd5 Nxd5 7. c4 a6 8.
Nd6+ Bxd6 9. cxd5 Bf5 10. d4 exd4 11. Qe2+ Qe7 12. Qxe7+ Bxe7 13. Bg5 Bxg5 14.
Nh3 Bh4+ 15. g3 Bg5 16. Nxg5 O-O 17. b4 Nc6 18. b5 Ne5 19. b6 Nxf3+ 20. Kf2 Nxg5
21. Rac1 Ne4+ 22. Kf3 Nd2+ 23. Kf4 Bb1 24. Rxb1 Nxb1 25. Rxb1 h6 26. g4 g5+ 27.
Kf3 d3 28. Rd1 Rad8 29. Rxd3 a5 30. d6 Rfe8 31. d7 Re1 32. a3 Rh1 33. h4 Rxh4
34. Rd6 h5 35. gxh5 Rxh5 36. Kg4 Rh4+ 37. Kxg5 Ra4 38. Kf6 Rxa3 39. Re6 Kf8 40.
Re7 Rf3+ 41. Ke5 a4 42. Re8+ Rxe8+ 43. dxe8=Q+ Kxe8 44. Kd6 a3 45. Kc7 Rb3 46.
Kxb7 a2 47. Kc8 a1=Q 48. b7 Qc3+ 49. Kb8 Qe5+ 50. Ka7 Qd5 51. b8=Q+ Rxb8 52.
Kxb8 Qc6 53. Ka7 Kd7 54. Kb8 Qc7+ 55. Ka8 Kc6 1/2-1/2`,
	)

	b.ResetTimer()
	for range b.N {
		_, _ = standardchess.NewBoardFromMoves(moves)
	}
}
