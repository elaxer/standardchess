package mover_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/move"
	"github.com/elaxer/standardchess/internal/move/mover"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	whiteKingPosAfterShortCastling = chess.PositionFromString("g1")
	whiteRookPosAfterShortCastling = chess.PositionFromString("f1")

	whiteKingPosAfterLongCastling = chess.PositionFromString("c1")
	whiteRookPosAfterLongCastling = chess.PositionFromString("d1")

	blackKingPosAfterShortCastling = chess.PositionFromString("g8")
	blackRookPosAfterShortCastling = chess.PositionFromString("f8")

	blackKingPosAfterLongCastling = chess.PositionFromString("c8")
	blackRookPosAfterLongCastling = chess.PositionFromString("d8")
)

func TestMakeCastling(t *testing.T) {
	type args struct {
		castlingType move.Castling
		board        chess.Board
	}
	tests := []struct {
		name                           string
		args                           args
		kingPos, rookPos               chess.Position
		wantKingNewPos, wantRookNewPos chess.Position
	}{
		{
			"white_short",
			args{
				move.CastlingShort,
				standardtest.NewBoard(chess.SideWhite, nil),
			},
			chess.PositionFromString("e1"),
			chess.PositionFromString("h1"),
			whiteKingPosAfterShortCastling,
			whiteRookPosAfterShortCastling,
		},
		{
			"white_short_uncommon_position",
			args{
				move.CastlingShort,
				standardtest.NewBoard(chess.SideWhite, nil),
			},
			chess.PositionFromString("f1"),
			chess.PositionFromString("g1"),
			whiteKingPosAfterShortCastling,
			whiteRookPosAfterShortCastling,
		},
		{
			"white_long",
			args{
				move.CastlingLong,
				standardtest.NewBoard(chess.SideWhite, nil),
			},
			chess.PositionFromString("e1"),
			chess.PositionFromString("a1"),
			whiteKingPosAfterLongCastling,
			whiteRookPosAfterLongCastling,
		},
		{
			"white_long_uncommon_position",
			args{
				move.CastlingLong,
				standardtest.NewBoard(chess.SideWhite, nil),
			},
			chess.PositionFromString("g1"),
			chess.PositionFromString("a1"),
			whiteKingPosAfterLongCastling,
			whiteRookPosAfterLongCastling,
		},
		{
			"black_short",
			args{
				move.CastlingShort,
				standardtest.NewBoard(chess.SideBlack, nil),
			},
			chess.PositionFromString("e8"),
			chess.PositionFromString("h8"),
			blackKingPosAfterShortCastling,
			blackRookPosAfterShortCastling,
		},
		{
			"black_short_uncommon_position",
			args{
				move.CastlingShort,
				standardtest.NewBoard(chess.SideBlack, nil),
			},
			chess.PositionFromString("c8"),
			chess.PositionFromString("h8"),
			blackKingPosAfterShortCastling,
			blackRookPosAfterShortCastling,
		},
		{
			"black_long",
			args{
				move.CastlingLong,
				standardtest.NewBoard(chess.SideBlack, nil),
			},
			chess.PositionFromString("e8"),
			chess.PositionFromString("a8"),
			blackKingPosAfterLongCastling,
			blackRookPosAfterLongCastling,
		},
		{
			"black_long_uncommon_position",
			args{
				move.CastlingLong,
				standardtest.NewBoard(chess.SideBlack, nil),
			},
			chess.PositionFromString("b8"),
			chess.PositionFromString("a8"),
			blackKingPosAfterLongCastling,
			blackRookPosAfterLongCastling,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := tt.args.board

			king := piece.NewKing(board.Turn())
			rook := piece.NewRook(board.Turn())

			err := board.Squares().PlacePiece(king, tt.kingPos)
			require.NoError(t, err, "failed to place king")

			err = board.Squares().PlacePiece(rook, tt.rookPos)
			require.NoError(t, err, "failed to place rook")

			got, err := mover.MakeCastling(tt.args.castlingType, board)
			require.NoError(t, err)
			require.NotNil(t, got)

			assert.True(t, king.IsMoved(), "king should be marked as moved after castling")
			assert.True(t, rook.IsMoved(), "rook should be marked as moved after castling")

			kingNewPos := board.Squares().GetByPiece(king)
			assert.Equalf(t, tt.wantKingNewPos, kingNewPos, "king should be on %s, got %s", tt.wantKingNewPos, kingNewPos)

			rookNewPos := board.Squares().GetByPiece(rook)
			assert.Equal(t, tt.wantRookNewPos, rookNewPos, "rook should be on %s, got %s", tt.wantRookNewPos, rookNewPos)
		})
	}
}

func TestMakeCastling_Negative(t *testing.T) {
	_, err := mover.MakeCastling(
		move.CastlingShort,
		standardtest.NewBoard(chess.SideWhite, map[chess.Position]chess.Piece{
			chess.PositionFromString("e1"): standardtest.NewPiece("K"),
			chess.PositionFromString("h1"): standardtest.NewPiece("R"),
			chess.PositionFromString("f1"): standardtest.NewPiece("N"),
		}),
	)

	assert.Error(t, err)
}
