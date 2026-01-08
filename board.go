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

var edgePosition = chess.NewPosition(chess.FileH, chess.Rank8)

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
	turn           chess.Side
	squares        *chess.Squares
	moveHistory    []chess.MoveResult
	capturedPieces []chess.Piece
	stateRules     []rule.Rule

	moves []chess.Position
	state chess.State
}

func NewBoardEmpty(turn chess.Side, placement map[chess.Position]chess.Piece, edgePosition chess.Position) (chess.Board, error) {
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

func NewBoard() chess.Board {
	board, err := NewBoardEmpty(chess.SideWhite, nil, edgePosition)
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

func NewBoardFromMoves(moves []chess.Move) (chess.Board, error) {
	board := NewBoard()
	for i, move := range moves {
		if _, err := board.MakeMove(move); err != nil {
			return nil, fmt.Errorf("%s#%d: %w", move, i+1, err)
		}
	}

	return board, nil
}

func (b *board) Squares() *chess.Squares {
	return b.squares
}

func (b *board) Turn() chess.Side {
	return b.turn
}

func (b *board) State() chess.State {
	if b.state != nil {
		return b.state
	}

	for _, rule := range b.stateRules {
		if state := rule(b, b.turn); state != nil {
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

	if p.Side() != b.Turn() {
		return pseudoMoves
	}

	legalMoves := make([]chess.Position, 0, cap(pseudoMoves))
	for _, to := range pseudoMoves {
		//nolint:errcheck,gosec
		b.squares.MovePieceTemporarily(from, to, func() {
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

		placement := &Placement{Piece: piece, LegalMoves: make([]chess.Position, 0, 27), Position: pos}
		if piece.Side() == b.turn {
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
		"placements":      placements,
	})
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
