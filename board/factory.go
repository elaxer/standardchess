package board

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/chess/position"
	"github.com/elaxer/standardchess/piece"
	"github.com/elaxer/standardchess/state/rule"
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

type factory struct{}

func NewFactory() chess.BoardFactory {
	return &factory{}
}

func (f *factory) Create(turn chess.Side, placement map[position.Position]chess.Piece) (chess.Board, error) {
	squares, err := chess.SquaresFromPlacement(edgePosition, placement)
	if err != nil {
		return nil, err
	}

	return &board{
		turn:           turn,
		squares:        squares,
		movesHistory:   make([]chess.MoveResult, 0, 128),
		capturedPieces: make([]chess.Piece, 0, 30),

		stateRules: stateRules,
	}, nil
}

func (f *factory) CreateFilled() chess.Board {
	board, _ := f.Create(chess.SideWhite, nil)
	for i, notation := range firstRowPieceNotations {
		file := position.File(i + 1)

		board.Squares().PlacePiece(piece.New(notation, chess.SideWhite), position.New(file, position.RankMin))
		board.Squares().PlacePiece(piece.NewPawn(chess.SideWhite), position.New(file, position.RankMin+1))

		board.Squares().PlacePiece(piece.New(notation, chess.SideBlack), position.New(file, edgePosition.Rank))
		board.Squares().PlacePiece(piece.NewPawn(chess.SideBlack), position.New(file, edgePosition.Rank-1))
	}

	return board
}

func (f *factory) CreateFromMoves(moves []chess.Move) (chess.Board, error) {
	board := f.CreateFilled()
	for i, move := range moves {
		if _, err := board.MakeMove(move); err != nil {
			return nil, fmt.Errorf("%s#%d: %w", move, i+1, err)
		}
	}

	return board, nil
}
