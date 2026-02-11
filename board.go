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
	"github.com/elaxer/standardchess/internal/move/normal"
	"github.com/elaxer/standardchess/internal/move/promotion"
	"github.com/elaxer/standardchess/internal/mover"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/rule"
	"github.com/elaxer/standardchess/metric"
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

	lastMovements := make([]map[string]string, 0, 2)
	if len(b.moveHistory) > 0 {
		switch move := b.moveHistory[len(b.moveHistory)-1].(type) {
		case *normal.MoveResult:
			lastMovements = append(lastMovements, map[string]string{
				"from": move.FromFull.String(),
				"to":   move.InputMove.To.String(),
			})
		case *promotion.MoveResult:
			lastMovements = append(lastMovements, map[string]string{
				"from": move.FromFull.String(),
				"to":   move.InputMove.To.String(),
			})
		case *enpassant.MoveResult:
			lastMovements = append(lastMovements, map[string]string{
				"from": move.FromFull.String(),
				"to":   move.InputMove.To.String(),
			})
		case *castling.MoveResult:
			lastMovements = append(
				lastMovements,
				map[string]string{
					"from": castling.KingInitPosition(move.Side()).String(),
					"to":   castling.KingCastledPosition(move.CastlingType, move.Side()).String(),
				},
				map[string]string{
					"from": castling.RookInitPosition(move.CastlingType, move.Side()).String(),
					"to":   castling.RookCastledPosition(move.CastlingType, move.Side()).String(),
				},
			)
		}
	}

	return json.Marshal(map[string]any{
		"turn":            b.turn,
		"is_check":        check.IsCheck(b),
		"state":           b.State(),
		"castlings":       metric.CastlingAbility(b).Value().(metric.Castlings)["practical"][b.turn],
		"captured_pieces": b.capturedPieces,
		"move_history":    b.moveHistory,
		"placement":       placements,
		"last_movements":  lastMovements,
	})
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
