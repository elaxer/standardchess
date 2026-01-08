package enpassant

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/normal"
	"github.com/elaxer/standardchess/internal/piece"
)

func CanEnPassant(board chess.Board) bool {
	return !EnPassantPosition(board).IsEmpty()
}

func EnPassantPosition(board chess.Board) chess.Position {
	if len(board.MoveHistory()) == 0 {
		return chess.NewPositionEmpty()
	}

	lastMove := board.MoveHistory()[len(board.MoveHistory())-1]
	normalMove, ok := lastMove.(*normal.MoveResult)
	if !ok || normalMove.InputMove.PieceNotation != piece.NotationPawn {
		return chess.NewPositionEmpty()
	}

	if absRank(normalMove.InputMove.To.Rank-normalMove.FromFull.Rank) != 2 {
		return chess.NewPositionEmpty()
	}

	return chess.NewPosition(
		normalMove.InputMove.To.File,
		normalMove.FromFull.Rank+piece.PawnRankDirection(!board.Turn()),
	)
}

func enPassantRank(side chess.Color) chess.Rank {
	if side.IsBlack() {
		return chess.Rank4
	}

	return chess.Rank5
}

func absRank(rank chess.Rank) chess.Rank {
	if rank < 0 {
		return -rank
	}

	return rank
}
