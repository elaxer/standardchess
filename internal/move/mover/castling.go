package mover

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/result"
	"github.com/elaxer/standardchess/internal/move/validator"
	"github.com/elaxer/standardchess/internal/piece"
)

type Castling struct {
}

func (m *Castling) Make(castlingType move.Castling, board chess.Board) (chess.MoveResult, error) {
	if err := validator.ValidateCastlingMove(castlingType, board.Turn(), board, true); err != nil {
		return nil, err
	}

	fileDir := fileDirection(castlingType)

	king, kingPosition := board.Squares().FindPiece(piece.NotationKing, board.Turn())
	rook, rookPosition, err := m.getRook(fileDir, board.Squares(), kingPosition)
	if err != nil {
		return nil, err
	}

	rank := kingPosition.Rank
	// todo: изменить на конкретные позиции (для шахмат 960)
	kingNewPosition := chess.NewPosition(kingPosition.File+fileDir*2, rank)
	rookNewPosition := chess.NewPosition(kingPosition.File+fileDir, rank)

	if _, err := board.Squares().MovePiece(kingPosition, kingNewPosition); err != nil {
		return nil, err
	}
	if _, err := board.Squares().MovePiece(rookPosition, rookNewPosition); err != nil {
		return nil, err
	}

	king.MarkMoved()
	rook.MarkMoved()

	return &result.Castling{Abstract: newAbstractResult(board), Castling: castlingType}, nil
}

func (m *Castling) getRook(direction chess.File, squares *chess.Squares, kingPosition chess.Position) (chess.Piece, chess.Position, error) {
	for position, p := range squares.IterByDirection(kingPosition, chess.NewPosition(direction, 0)) {
		if p != nil && p.Notation() == piece.NotationRook {
			return p, position, nil
		}
	}

	return nil, chess.NewPositionEmpty(), fmt.Errorf("%w: rook wasn't found", validator.ErrCastling)
}

func fileDirection(castlingType move.Castling) chess.File {
	return map[move.Castling]chess.File{
		move.CastlingShort: 1,
		move.CastlingLong:  -1,
	}[castlingType]
}
