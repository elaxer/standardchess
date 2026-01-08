package enpassant

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/piecemove"
	"github.com/elaxer/standardchess/internal/move/resolver"
	"github.com/elaxer/standardchess/internal/move/result"
	"github.com/elaxer/standardchess/internal/piece"
)

func MakeMove(move *Move, board chess.Board) (*MoveResult, error) {
	if err := move.Validate(); err != nil {
		return nil, err
	}

	fullFrom, err := resolver.ResolveFrom(move.From, move.To, piece.NotationPawn, board, board.Turn())
	if err != nil {
		return nil, err
	}
	if err := ValidateMove(fullFrom, move.To, board); err != nil {
		return nil, err
	}
	shortenedFrom, err := resolver.UnresolveFrom(fullFrom, move.To, board)
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

	return &MoveResult{
		PieceMoveResult: piecemove.PieceMoveResult{
			Abstract:      &result.Abstract{MoveSide: board.Turn()},
			WasMoved:      true,
			FromFull:      fullFrom,
			FromShortened: shortenedFrom,
			Captured:      capturedPawn,
		},
		InputMove: *move,
	}, nil
}

func UndoMove(move *MoveResult, board chess.Board) error {
	if err := move.Validate(); err != nil {
		return err
	}

	if _, err := board.Squares().MovePiece(move.InputMove.To, move.FromFull); err != nil {
		return err
	}

	capturedPiecePosition := chess.NewPosition(move.InputMove.To.File, enPassantRank(move.MoveSide))
	if err := board.Squares().PlacePiece(move.Captured, capturedPiecePosition); err != nil {
		return err
	}

	return nil
}
