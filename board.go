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
	"errors"
	"slices"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/check"
	"github.com/elaxer/standardchess/internal/move/castling"
	"github.com/elaxer/standardchess/internal/move/enpassant"
	"github.com/elaxer/standardchess/internal/mover"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/rule"
)

var EdgePosition = chess.NewPosition(chess.FileH, chess.Rank8)

var (
	ErrCannotMoveInTerminalState = errors.New("cannot make a move when the board is in a terminal state")
	ErrNoMovesToUndo             = errors.New("there are no moves to undo")
)

type moveObserver interface {
	AfterBoardMakeMove(board chess.Board)
	AfterBoardUndoMove()
}

type board struct {
	turn           chess.Color
	squares        *chess.Squares
	moveHistory    []chess.Move
	capturedPieces []chess.Piece
	stateRules     []rule.Rule

	moves []chess.Position
	state chess.State

	observers []moveObserver
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

func (b *board) MoveHistory() []chess.Move {
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

	enPassantPosition := enpassant.EnPassantTargetSquare(b)
	if enpassant.ValidateMove(from, enPassantPosition, b) == nil {
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

func (b *board) MakeMove(move string) (chess.Move, error) {
	if b.State().IsTerminal() {
		return nil, ErrCannotMoveInTerminalState
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

	for _, observer := range b.observers {
		observer.AfterBoardMakeMove(b)
	}

	moveResult.SetBoardNewState(b.State())

	return moveResult, nil
}

func (b *board) UndoLastMove() (chess.Move, error) {
	movesCount := len(b.moveHistory)
	if movesCount == 0 {
		return nil, ErrNoMovesToUndo
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

	for _, observer := range b.observers {
		observer.AfterBoardUndoMove()
	}

	return lastMove, nil
}

func (b *board) MarshalJSON() ([]byte, error) {
	type Piece struct {
		Color      chess.Color      `json:"color"`
		Notation   string           `json:"notation"`
		IsMoved    bool             `json:"is_moved"`
		LegalMoves []chess.Position `json:"legal_moves"`
	}
	type Placement struct {
		Piece    *Piece         `json:"piece"`
		Position chess.Position `json:"position"`
	}

	placements := make([]*Placement, 0, 32)
	for position, piece := range b.squares.Iter() {
		if piece == nil {
			continue
		}

		placement := &Placement{
			&Piece{
				piece.Color(),
				piece.Notation(),
				piece.IsMoved(),
				make([]chess.Position, 0),
			},
			position,
		}
		if piece.Color() == b.turn {
			placement.Piece.LegalMoves = b.LegalMoves(piece)
		}

		placements = append(placements, placement)
	}

	return json.Marshal(map[string]any{
		"turn":     b.turn,
		"is_check": check.IsCheck(b),
		"state":    b.State(),
		"castlings": map[string]bool{
			"O-O":   castling.ValidateMoveWithObstacle(castling.TypeShort, b.turn, b) == nil,
			"O-O-O": castling.ValidateMoveWithObstacle(castling.TypeLong, b.turn, b) == nil,
		},
		"captured_pieces": b.capturedPieces,
		"move_history":    b.moveHistory,
		"placement":       placements,
	})
}
