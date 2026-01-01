package enpassant

import (
	"math"

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

	if normalMove.InputMove.To.Rank-normalMove.FromFull.Rank != chess.Rank(math.Abs(2)) {
		return chess.NewPositionEmpty()
	}

	return chess.NewPosition(
		normalMove.InputMove.To.File,
		normalMove.FromFull.Rank+piece.PawnRankDirection(!board.Turn()),
	)
}

func AttackingPawns(board chess.Board) [2]chess.Piece {
	var pawns [2]chess.Piece

	enPassantPosition := EnPassantPosition(board)
	if enPassantPosition.IsEmpty() {
		return pawns
	}

	rank := enPassantRank(board.Turn())
	pawnPositions := [2]chess.Position{
		chess.NewPosition(enPassantPosition.File-1, rank),
		chess.NewPosition(enPassantPosition.File+1, rank),
	}
	for i, pawnPosition := range pawnPositions {
		if p, err := board.Squares().FindByPosition(pawnPosition); err != nil {
			if p.Notation() == piece.NotationPawn && p.Side() == board.Turn() {
				pawns[i] = p
			}
		}
	}

	return pawns
}

func enPassantRank(side chess.Side) chess.Rank {
	if side.IsBlack() {
		return chess.Rank5
	}

	return chess.Rank4
}
