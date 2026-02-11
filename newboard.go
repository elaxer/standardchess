package standardchess

import (
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/repetition"
	"github.com/elaxer/standardchess/internal/rule"
)

var firstRowPieceNotations = [...]string{
	piece.NotationRook,
	piece.NotationKnight,
	piece.NotationBishop,
	piece.NotationQueen,
	piece.NotationKing,
	piece.NotationBishop,
	piece.NotationKnight,
	piece.NotationRook,
}

func NewBoard() chess.Board {
	board, err := NewBoardEmpty(chess.ColorWhite, nil, EdgePosition)
	must(err)

	squares := board.Squares()
	for i, notation := range firstRowPieceNotations {
		//nolint:gosec
		file := chess.File(i + 1)

		wPiece, err := piece.New(notation, chess.ColorWhite)
		must(err)

		must(squares.PlacePiece(wPiece, chess.NewPosition(file, chess.RankMin)))
		must(
			squares.PlacePiece(
				piece.NewPawn(chess.ColorWhite),
				chess.NewPosition(file, chess.RankMin+1),
			),
		)

		bPiece, err := piece.New(notation, chess.ColorBlack)
		must(err)

		must(squares.PlacePiece(bPiece, chess.NewPosition(file, EdgePosition.Rank)))
		must(
			squares.PlacePiece(
				piece.NewPawn(chess.ColorBlack),
				chess.NewPosition(file, EdgePosition.Rank-1),
			),
		)
	}

	return board
}

func NewBoardFromMoves(moves []string) (chess.Board, error) {
	board := NewBoard()
	for i, move := range moves {
		if _, err := board.MakeMove(move); err != nil {
			return nil, fmt.Errorf("%s#%d: %w", move, i+1, err)
		}
	}

	return board, nil
}

func NewBoardEmpty(
	turn chess.Color,
	placement map[chess.Position]chess.Piece,
	edgePosition chess.Position,
) (chess.Board, error) {
	squares, err := chess.SquaresFromPlacement(edgePosition, placement)
	if err != nil {
		return nil, err
	}

	repetition := repetition.NewRepetition()

	var stateRules = []rule.Rule{
		rule.Checkmate,
		rule.Stalemate,
		rule.NewFivefoldRepetition(repetition).Rule,
		rule.NewThreefoldRepetition(repetition).Rule,
		rule.FiftyMoves,
	}

	return &board{
		turn:           turn,
		squares:        squares,
		moveHistory:    make([]chess.Move, 0, 128),
		moves:          make([]chess.Position, 0, 64),
		capturedPieces: make([]chess.Piece, 0, 30),

		stateRules: stateRules,

		observers: []moveObserver{repetition},
	}, nil
}
