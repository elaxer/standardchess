// Package standardchess contains logic for working with a chessboard.
// The package has functions for creating boards in different ways:
// empty boards, boards with a starting position, boards based on a list of moves.
// The created boards have logic for executing or canceling moves, manipulating the board position,
// as well as methods for obtaining various information about the board, such as its current state,
// available moves, move history, and so on.
//
// The package also contains code for creating chess pieces: rook, knight, bishop, queen, king, pawns.
package standardchess

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/move/enpassant"
	"github.com/elaxer/standardchess/internal/mover"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/state/rule"
	"github.com/elaxer/standardchess/metric"
)

var EdgePosition = chess.NewPosition(chess.FileH, chess.Rank8)

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

var stateRules = []rule.Rule{
	rule.Checkmate,
	rule.Stalemate,
	rule.Check,

	rule.FiftyMoves,
}

type board struct {
	turn           chess.Color
	squares        *chess.Squares
	moveHistory    []chess.MoveResult
	capturedPieces []chess.Piece
	stateRules     []rule.Rule

	moves []chess.Position
	state chess.State
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

func NewBoardFromMoves(moves []chess.Move) (chess.Board, error) {
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

	return &board{
		turn:           turn,
		squares:        squares,
		moveHistory:    make([]chess.MoveResult, 0, 128),
		moves:          make([]chess.Position, 0, 128),
		capturedPieces: make([]chess.Piece, 0, 30),

		stateRules: stateRules,
	}, nil
}

func (b *board) Squares() *chess.Squares {
	return b.squares
}

func (b *board) Turn() chess.Color {
	return b.turn
}

func (b *board) State() chess.State {
	if b.state != nil {
		return b.state
	}

	for _, rule := range b.stateRules {
		if state := rule(b); state != nil {
			b.state = state

			return b.state
		}
	}

	b.state = chess.StateClear

	return b.state
}

func (b *board) CapturedPieces() []chess.Piece {
	return b.capturedPieces
}

func (b *board) MoveHistory() []chess.MoveResult {
	return b.moveHistory
}

func (b *board) Moves() []chess.Position {
	if len(b.moves) > 0 {
		return b.moves
	}

	uniqueMoves := make(map[chess.Position]bool, 32)

	for piece := range b.squares.GetAllPieces(b.turn) {
		for _, move := range b.LegalMoves(piece) {
			uniqueMoves[move] = true
		}
	}

	for move := range uniqueMoves {
		b.moves = append(b.moves, move)
	}

	return b.moves
}

func (b *board) LegalMoves(p chess.Piece) []chess.Position {
	from := b.squares.GetByPiece(p)
	if from.IsEmpty() {
		return make([]chess.Position, 0)
	}

	pseudoMoves := p.PseudoMoves(from, b.squares)

	if p.Color() != b.Turn() {
		return pseudoMoves
	}

	legalMoves := make([]chess.Position, 0, cap(pseudoMoves))
	for _, to := range pseudoMoves {
		_ = b.squares.MovePieceTemporarily(from, to, func() {
			_, kingPosition := b.squares.FindPiece(piece.NotationKing, b.turn)
			if !b.IsSquareAttacked(kingPosition) {
				legalMoves = append(legalMoves, to)
			}
		})
	}

	enPassantPosition := enpassant.EnPassantPosition(b)
	if err := enpassant.ValidateMove(from, enPassantPosition, b); err == nil {
		legalMoves = append(legalMoves, enPassantPosition)
	}

	return legalMoves
}

func (b *board) IsSquareAttacked(position chess.Position) bool {
	for piece := range b.squares.GetAllPieces(!b.turn) {
		from := b.squares.GetByPiece(piece)
		if slices.Contains(piece.PseudoMoves(from, b.squares), position) {
			return true
		}
	}

	return false
}

func (b *board) MakeMove(move chess.Move) (chess.MoveResult, error) {
	if b.State().Type().IsTerminal() {
		return nil, nil
	}

	moveResult, err := mover.MakeMove(move, b)
	if err != nil {
		return nil, err
	}

	b.moveHistory = append(b.moveHistory, moveResult)
	b.turn = !b.turn
	if moveResult.CapturedPiece() != nil {
		b.capturedPieces = append(b.capturedPieces, moveResult.CapturedPiece())
	}

	b.moves = b.moves[:0]
	b.state = nil

	moveResult.SetBoardNewState(b.State())

	return moveResult, nil
}

func (b *board) UndoLastMove() (chess.MoveResult, error) {
	movesCount := len(b.moveHistory)
	if movesCount == 0 {
		return nil, nil
	}

	lastMove := b.moveHistory[movesCount-1]
	b.moveHistory = b.moveHistory[:movesCount-1]

	if err := mover.UndoMove(lastMove, b); err != nil {
		return nil, err
	}

	b.turn = !b.turn
	if lastMove.CapturedPiece() != nil {
		_ = slices.Delete(b.capturedPieces, len(b.capturedPieces)-1, len(b.capturedPieces))
	}

	b.moves = b.moves[:0]
	b.state = nil

	return lastMove, nil
}

func (b *board) MarshalJSON() ([]byte, error) {
	type Placement struct {
		Piece      chess.Piece      `json:"piece"`
		Position   chess.Position   `json:"position"`
		LegalMoves []chess.Position `json:"legal_moves"`
	}

	placements := make([]*Placement, 0, 32)
	for pos, piece := range b.squares.Iter() {
		if piece == nil {
			continue
		}

		placement := &Placement{
			Piece:      piece,
			LegalMoves: make([]chess.Position, 0, 27),
			Position:   pos,
		}
		if piece.Color() == b.turn {
			placement.LegalMoves = b.LegalMoves(piece)
		}

		placements = append(placements, placement)
	}

	return json.Marshal(map[string]any{
		"turn":            b.turn,
		"state":           b.State(),
		"castlings":       metric.CastlingAbility(b).Value().(metric.Castlings)["practical"][b.turn],
		"captured_pieces": b.capturedPieces,
		"move_history":    b.moveHistory,
		"placement":       placements,
	})
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
