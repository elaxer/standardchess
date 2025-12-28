package standardchess

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/state/rule"
)

var firstRowPieceNotations = []string{
	piece.NotationRook,
	piece.NotationKnight,
	piece.NotationBishop,
	piece.NotationQueen,
	piece.NotationKing,
	piece.NotationBishop,
	piece.NotationKnight,
	piece.NotationRook,
}

var stateRules = []rule.Rule{
	rule.Checkmate,
	rule.Stalemate,
	rule.Check,

	rule.FiftyMoves,
}

func New(turn chess.Side, placement map[chess.Position]chess.Piece) (chess.Board, error) {
	return NewSized(turn, placement, edgePosition)
}

func NewSized(turn chess.Side, placement map[chess.Position]chess.Piece, edgePosition chess.Position) (chess.Board, error) {
	squares, err := chess.SquaresFromPlacement(edgePosition, placement)
	if err != nil {
		return nil, err
	}

	return &board{
		turn:           turn,
		squares:        squares,
		moveHistory:    make([]chess.MoveResult, 0, 128),
		capturedPieces: make([]chess.Piece, 0, 30),

		stateRules: stateRules,
	}, nil
}

func NewFilled() chess.Board {
	board, err := New(chess.SideWhite, nil)
	must(err)

	for i, notation := range firstRowPieceNotations {
		file := chess.File(i + 1)

		wPiece, err := piece.New(notation, chess.SideWhite)
		must(err)

		must(board.Squares().PlacePiece(wPiece, chess.NewPosition(file, chess.RankMin)))
		must(board.Squares().PlacePiece(piece.NewPawn(chess.SideWhite), chess.NewPosition(file, chess.RankMin+1)))

		bPiece, err := piece.New(notation, chess.SideBlack)
		must(err)

		must(board.Squares().PlacePiece(bPiece, chess.NewPosition(file, edgePosition.Rank)))
		must(board.Squares().PlacePiece(piece.NewPawn(chess.SideBlack), chess.NewPosition(file, edgePosition.Rank-1)))
	}

	return board
}

func NewFromMoves(moves []chess.Move) (chess.Board, error) {
	board := NewFilled()
	for i, move := range moves {
		if _, err := board.MakeMove(move); err != nil {
			return nil, fmt.Errorf("%s#%d: %w", move, i+1, err)
		}
	}

	return board, nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
