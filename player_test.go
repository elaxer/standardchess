package standardchess_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess"
	"github.com/elaxer/standardchess/encoding/fen"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const initFENStr = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func TestNewBoardPlayer(t *testing.T) {
	board := standardtest.NewBoardFromPGN(`1. e4 e5 2. Nf3 Qf6 3. d4 exd4 4. Nxd4 Nc6 5. Be3 Bb4+ 6. c3 Bc5 7. Bb5 Bxd4 8.
cxd4 b6 9. e5 Qg6 10. O-O Nh6 11. d5 Ne7 12. Re1 Nhf5 13. Nc3 Nxe3 14. fxe3 Qg5
15. e6 Qh4 16. exf7+ Kxf7 17. Qf3+ Kg8 18. d6 c6 19. dxe7 Qxe7 20. Rf1 cxb5 21.
Nxb5 d5 22. Qxd5+ Be6 23. Qxa8+ Bc8 24. Qxc8+ Qf8 25. Qxf8# 1-0`)

	player := standardchess.NewBoardPlayer(board)
	assert.Equal(t, "5Qkr/p5pp/1p6/1N6/8/4P3/PP4PP/R4RK1 b - - 0 25", fen.Encode(player.Board()).String())
}

func TestBoardPlayer_GoTo(t *testing.T) {
	rewinds := []struct {
		n          uint16
		wantOK     bool
		wantFENStr string
	}{
		{
			27,
			true,
			"r1b1k2r/ppp3pp/8/Q5B1/3P4/2P5/PP3PPP/R3KB1R b KQk - 0 14",
		},
		{
			47, // last
			true,
			"r7/p1p4p/1p6/7b/3P3k/2P2PQ1/PP2r1PP/2KR3R b - - 3 24",
		},
		{
			0,
			true,
			initFENStr,
		},
		{
			1,
			true,
			"rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
		},
		{
			48,
			false,
			"",
		},
	}

	board := standardtest.NewBoardFromPGN(
		`1. e4 e5 2. Nf3 Bb4 3. Nc3 Nf6 4. Nxe5 Nxe4 5. Nxe4 f6 6. Nc4 f5 7. Ng5 Qxg5 8.
c3 Ba5 9. Nxa5 f4 10. d4 f3 11. Bxg5 d5 12. Qxf3 Nc6 13. Qxd5 Nxa5 14. Qxa5 b6
15. Qe5+ Kf7 16. Qe7+ Kg6 17. Bd3+ Kh5 18. Qxg7 Re8+ 19. Be2+ Bg4 20. Qf7+ Kxg5
21. f3 Bh5 22. O-O-O Rxe2 23. Qg7+ Kh4 24. Qg3# 1-0`)
	player := standardchess.NewBoardPlayer(board)

	for _, r := range rewinds {
		ok := player.GoTo(r.n)
		require.Equal(t, r.wantOK, ok)
		if !r.wantOK {
			continue
		}
		fen := fen.Encode(player.Board())
		require.Equal(t, r.wantFENStr, fen.String())
	}
}

func TestBoardPlayer_Reset(t *testing.T) {
	board := standardtest.NewBoardFromPGN(`1. e4 e5 2. Nf3 Bb4 3. Nc3 Nf6 4. Nxe5 Nxe4 5. Nxe4 f6 6. Nc4 f5 7. Ng5 Qxg5 8.
c3 Ba5 9. Nxa5 f4 10. d4 f3 11. Bxg5 d5 12. Qxf3 Nc6 13. Qxd5 Nxa5 14. Qxa5 b6
15. Qe5+ Kf7 16. Qe7+ Kg6 17. Bd3+ Kh5 18. Qxg7 Re8+ 19. Be2+ Bg4 20. Qf7+ Kxg5
21. f3 Bh5 22. O-O-O Rxe2 23. Qg7+ Kh4 24. Qg3# 1-0`)

	player := standardchess.NewBoardPlayer(board)
	player.Reset()

	fen := fen.Encode(player.Board()).String()
	require.Equal(t, initFENStr, fen)
}

func TestBoardPlayer_End(t *testing.T) {
	board := standardtest.NewBoardFromPGN(`1. e4 e5 2. Nf3 Nc6 3. Bd3 Nf6 4. O-O d5 5. exd5 Nxd5 6. Qe1 Bc5 7. Nxe5 O-O 8.
Nxc6 bxc6 9. c4 Nf4 10. Nc3 Bg4 11. g3 Nh3+ 12. Kg2 Qxd3 13. Rh1 Qf3+ 14. Kf1
Qxh1# 0-1`)

	player := standardchess.NewBoardPlayer(board)
	player.End()

	fen := fen.Encode(player.Board()).String()
	require.Equal(t, "r4rk1/p1p2ppp/2p5/2b5/2P3b1/2N3Pn/PP1P1P1P/R1B1QK1q w - - 0 15", fen)
}

func TestBoardPlayer_AfterNewMoves(t *testing.T) {
	board := standardchess.NewBoard()
	player := standardchess.NewBoardPlayer(board)

	_, err := board.MakeMove(chess.StringMove("e4"))
	require.NoError(t, err)
	_, err = board.MakeMove(chess.StringMove("e5"))
	require.NoError(t, err)
	_, err = board.MakeMove(chess.StringMove("Nf3"))
	require.NoError(t, err)

	player.Reset()
	require.Equal(t, initFENStr, fen.Encode(player.Board()).String())

	player.Next()
	require.Equal(t, "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1", fen.Encode(player.Board()).String())
	player.Next()
	require.Equal(t, "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2", fen.Encode(player.Board()).String())
	player.Next()
	require.Equal(t, "rnbqkbnr/pppp1ppp/8/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2", fen.Encode(player.Board()).String())

	ok := player.Next()
	require.False(t, ok)

	_, err = board.UndoLastMove()
	require.NoError(t, err)
	require.Equal(t, "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2", fen.Encode(player.Board()).String())

}

func TestBoardPlayer_Prev(t *testing.T) {
	board := standardtest.NewBoardFromPGN(`1. e4 e5 2. Nf3 Nc6 3. Bd3 Nf6 4. O-O d5 5. exd5 Nxd5 6. Qe1 Bc5 7. Nxe5 O-O 8.
Nxc6 bxc6 9. c4 Nf4 10. Nc3 Bg4 11. g3 Nh3+ 12. Kg2 Qxd3 13. Rh1 Qf3+ 14. Kf1
Qxh1# 0-1`)

	player := standardchess.NewBoardPlayer(board)
	player.Reset()

	ok := player.Prev()
	require.False(t, ok)

	require.Equal(t, initFENStr, fen.Encode(player.Board()).String())
}
