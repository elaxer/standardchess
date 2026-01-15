package rule_test

import (
	"testing"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/rule"
	"github.com/elaxer/standardchess/internal/standardtest"
	"github.com/elaxer/standardchess/internal/state"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	type args struct {
		board chess.Board
	}
	tests := []struct {
		name string
		args args
		want chess.State
	}{
		{
			"check",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("a8"): piece.NewRook(chess.ColorBlack),
				}),
			},
			state.Check,
		},
		{
			"check_bishop",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("e1"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("h8"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("b4"): piece.NewBishop(chess.ColorWhite),
				}),
			},
			state.Check,
		},
		{
			"check_pawns",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("d4"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("c3"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("e3"): piece.NewBishop(chess.ColorWhite),
				}),
			},
			state.Check,
		},
		{
			"no_check",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("d4"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("a1"): piece.NewRook(chess.ColorBlack),
				}),
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rule.Check(tt.args.board)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCheckmate(t *testing.T) {
	type args struct {
		board chess.Board
	}
	tests := []struct {
		name string
		args args
		want chess.State
	}{
		{
			"checkmate",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("a8"): piece.NewRook(chess.ColorBlack),
					chess.PositionFromString("b8"): piece.NewRook(chess.ColorBlack),
				}),
			},
			state.Checkmate,
		},
		{
			// no checkmate because the black king can capture the threatening rook
			"no_checkmate",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a1"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("h8"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("a2"): piece.NewRook(chess.ColorBlack),
					chess.PositionFromString("b8"): piece.NewRook(chess.ColorBlack),
				}),
			},

			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rule.Checkmate(tt.args.board)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStalemate(t *testing.T) {
	type args struct {
		board chess.Board
	}
	tests := []struct {
		name string
		args args
		want chess.State
	}{
		{
			"stalemate",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorWhite, map[chess.Position]chess.Piece{
					chess.PositionFromString("a8"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("b6"): piece.NewKing(chess.ColorBlack),
					chess.PositionFromString("c7"): piece.NewQueen(chess.ColorBlack),
				}),
			},
			state.Stalemate,
		},
		{
			"no_stalemate",
			args{
				standardtest.NewBoardEmpty8x8(chess.ColorBlack, map[chess.Position]chess.Piece{
					chess.PositionFromString("a8"): piece.NewKing(chess.ColorWhite),
					chess.PositionFromString("c7"): piece.NewQueen(chess.ColorWhite),
					chess.PositionFromString("b6"): piece.NewKing(chess.ColorBlack),
				}),
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rule.Stalemate(tt.args.board)
			assert.Equal(t, tt.want, got)
		})
	}
}
