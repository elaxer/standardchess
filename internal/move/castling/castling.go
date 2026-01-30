package castling

import "github.com/elaxer/chess"

func KingInitPosition(side chess.Color) chess.Position {
	return chess.NewPosition(chess.FileE, rank(side))
}

func KingCastledPosition(castling CastlingType, side chess.Color) chess.Position {
	if castling.IsLong() {
		return chess.NewPosition(chess.FileC, rank(side))
	}

	return chess.NewPosition(chess.FileG, rank(side))
}

func RookInitPosition(castling CastlingType, side chess.Color) chess.Position {
	if castling.IsLong() {
		return chess.NewPosition(chess.FileA, rank(side))
	}

	return chess.NewPosition(chess.FileH, rank(side))
}

func RookCastledPosition(castling CastlingType, side chess.Color) chess.Position {
	if castling.IsLong() {
		return chess.NewPosition(chess.FileD, rank(side))
	}

	return chess.NewPosition(chess.FileF, rank(side))
}

func rank(color chess.Color) chess.Rank {
	if color.IsBlack() {
		return chess.Rank8
	}

	return chess.Rank1
}
