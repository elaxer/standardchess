package mover

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
	"github.com/elaxer/standardchess/move/move"
	"github.com/elaxer/standardchess/move/result"
	"github.com/elaxer/standardchess/move/validator"
	"github.com/elaxer/standardchess/piece"
)

type Castling struct {
}

func (m *Castling) Make(castlingType move.Castling, board chess.Board) (chess.MoveResult, error) {
	if err := validator.ValidateCastlingMove(castlingType, board.Turn(), board, true); err != nil {
		return nil, err
	}

	direction := fileDirection(castlingType)

	_, kingPosition := board.Squares().FindPiece(piece.NotationKing, board.Turn())
	rookPosition, _ := m.rookPosition(direction, board.Squares(), kingPosition)

	rank := kingPosition.Rank
	kingNewPosition := position.New(kingPosition.File+direction*2, rank)
	rookNewPosition := position.New(kingPosition.File+direction, rank)

	if _, err := board.Squares().MovePiece(kingPosition, kingNewPosition); err != nil {
		return nil, err
	}
	if _, err := board.Squares().MovePiece(rookPosition, rookNewPosition); err != nil {
		return nil, err
	}

	king, err := board.Squares().FindByPosition(kingNewPosition)
	if err != nil {
		return nil, err
	}

	rook, err := board.Squares().FindByPosition(rookNewPosition)
	if err != nil {
		return nil, err
	}

	king.MarkMoved()
	rook.MarkMoved()

	return &result.Castling{Abstract: newAbstractResult(board), Castling: castlingType}, nil
}

func (m *Castling) rookPosition(direction position.File, squares *chess.Squares, kingPosition position.Position) (position.Position, error) {
	for position, p := range squares.IterByDirection(kingPosition, position.New(direction, 0)) {
		if p != nil && p.Notation() == piece.NotationRook {
			return position, nil
		}
	}

	return position.NewEmpty(), fmt.Errorf("%w: rook wasn't found", validator.ErrCastling)
}

func fileDirection(castlingType move.Castling) position.File {
	return map[move.Castling]position.File{
		move.CastlingShort: 1,
		move.CastlingLong:  -1,
	}[castlingType]
}
