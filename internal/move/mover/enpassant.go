package mover

import (
	"github.com/elaxer/chess"
	mv "github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/resolver"
	"github.com/elaxer/standardchess/internal/move/result"
	"github.com/elaxer/standardchess/internal/move/validator"
	"github.com/elaxer/standardchess/internal/piece"
)

func MakeEnPassant(move *mv.EnPassant, board chess.Board) (*result.EnPassant, error) {
	if err := move.Validate(); err != nil {
		return nil, err
	}

	fullFrom, err := resolver.ResolveFrom(move.PieceMove, piece.NotationPawn, board, board.Turn())
	if err != nil {
		return nil, err
	}
	if err := validator.ValidateEnPassant(fullFrom, move.To, board); err != nil {
		return nil, err
	}
	shortenedFrom, err := resolver.UnresolveFrom(mv.NewPieceMove(fullFrom, move.To), board)
	if err != nil {
		return nil, err
	}

	if _, err := board.Squares().MovePiece(fullFrom, move.To); err != nil {
		return nil, err
	}

	capturedPawnPosition := chess.NewPosition(move.To.File, move.To.Rank-piece.PawnRankDirection(board.Turn()))
	capturedPawn, err := board.Squares().FindByPosition(capturedPawnPosition)
	if err != nil {
		return nil, err
	}

	if err := board.Squares().PlacePiece(nil, capturedPawnPosition); err != nil {
		return nil, err
	}

	return &result.EnPassant{
		PieceMove: result.PieceMove{
			Abstract:      newAbstractResult(board),
			WasMoved:      true,
			FromFull:      fullFrom,
			FromShortened: shortenedFrom,
			Captured:      capturedPawn,
		},
		InputMove: *move,
	}, nil
}

func UndoEnPassant(move *result.EnPassant, board chess.Board) error {
	if err := move.Validate(); err != nil {
		return err
	}

	if _, err := board.Squares().MovePiece(move.InputMove.To, move.FromFull); err != nil {
		return nil
	}

	capturedPiecePosition := chess.NewPosition(move.InputMove.To.File, enPassantRank(move.MoveSide))
	if err := board.Squares().PlacePiece(move.Captured, capturedPiecePosition); err != nil {
		return err
	}

	return nil
}

func enPassantRank(side chess.Side) chess.Rank {
	if side.IsBlack() {
		return chess.Rank5
	}

	return chess.Rank4
}
