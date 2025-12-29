package mover

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/result"
	"github.com/elaxer/standardchess/internal/move/validator"
	"github.com/elaxer/standardchess/internal/piece"
)

var (
	kingFileAfterShortCastling = chess.FileG
	rookFileAfterShortCastling = chess.FileF

	kingFileAfterLongCastling = chess.FileC
	rookFileAfterLongCastling = chess.FileD
)

func MakeCastling(castlingType move.Castling, board chess.Board) (chess.MoveResult, error) {
	if err := validator.ValidateCastlingMove(castlingType, board.Turn(), board, true); err != nil {
		return nil, err
	}

	king, kingPosition := board.Squares().FindPiece(piece.NotationKing, board.Turn())
	rook, rookPosition, err := getRook(fileDirection(castlingType), board.Squares(), kingPosition)
	if err != nil {
		return nil, err
	}

	if err := board.Squares().PlacePiece(nil, kingPosition); err != nil {
		return nil, err
	}
	if err := board.Squares().PlacePiece(nil, rookPosition); err != nil {
		return nil, err
	}

	kingNewPos, rookNewPos := pickNewPositions(castlingType, kingPosition.Rank)

	if err := board.Squares().PlacePiece(king, kingNewPos); err != nil {
		return nil, err
	}
	if err := board.Squares().PlacePiece(rook, rookNewPos); err != nil {
		return nil, err
	}

	king.MarkMoved()
	rook.MarkMoved()

	return &result.Castling{Abstract: newAbstractResult(board), Castling: castlingType}, nil
}

func getRook(direction chess.File, squares *chess.Squares, kingPosition chess.Position) (chess.Piece, chess.Position, error) {
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

func pickNewPositions(castlingType move.Castling, rank chess.Rank) (kingPos, rookPos chess.Position) {
	kingFile := kingFileAfterShortCastling
	if castlingType == move.CastlingLong {
		kingFile = kingFileAfterLongCastling
	}
	rookFile := rookFileAfterShortCastling
	if castlingType == move.CastlingLong {
		rookFile = rookFileAfterLongCastling
	}

	return chess.NewPosition(kingFile, rank), chess.NewPosition(rookFile, rank)
}
