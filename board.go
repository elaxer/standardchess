package standardchess

import (
	"encoding/json"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/metric"
	"github.com/elaxer/standardchess/internal/move/mover"
	"github.com/elaxer/standardchess/internal/piece"
	"github.com/elaxer/standardchess/internal/state/rule"
)

var edgePosition = chess.NewPosition(chess.FileH, chess.Rank8)

type board struct {
	turn           chess.Side
	squares        *chess.Squares
	moveHistory    []chess.MoveResult
	capturedPieces []chess.Piece

	stateRules []rule.Rule
}

func (b *board) Squares() *chess.Squares {
	return b.squares
}

func (b *board) Turn() chess.Side {
	return b.turn
}

func (b *board) State(side chess.Side) chess.State {
	for _, rule := range b.stateRules {
		if state := rule(b, side); state != nil {
			return state
		}
	}

	return chess.StateClear
}

func (b *board) CapturedPieces() []chess.Piece {
	return b.capturedPieces
}

func (b *board) MoveHistory() []chess.MoveResult {
	return b.moveHistory
}

func (b *board) Moves(side chess.Side) chess.PositionSet {
	moves := mapset.NewSetWithSize[chess.Position](32)
	for _, piece := range b.squares.GetAllPieces(side) {
		moves = moves.Union(b.LegalMoves(piece))
	}

	return moves
}

func (b *board) LegalMoves(p chess.Piece) chess.PositionSet {
	from := b.squares.GetByPiece(p)
	if from.IsEmpty() {
		return nil
	}

	pseudoMoves := p.PseudoMoves(from, b.squares)

	if p.Side() != b.Turn() {
		return pseudoMoves
	}

	legalMoves := mapset.NewSetWithSize[chess.Position](pseudoMoves.Cardinality())
	for to := range pseudoMoves.Iter() {
		b.squares.MovePieceTemporarily(from, to, func() {
			_, kingPosition := b.squares.FindPiece(piece.NotationKing, b.turn)
			if !b.Moves(!b.turn).ContainsOne(kingPosition) {
				legalMoves.Add(to)
			}
		})
	}

	return legalMoves
}

func (b *board) MakeMove(move chess.Move) (chess.MoveResult, error) {
	moveResult, err := mover.MakeMove(move, b)
	if err != nil {
		return nil, err
	}

	b.moveHistory = append(b.moveHistory, moveResult)
	b.turn = !b.turn
	if moveResult.CapturedPiece() != nil {
		b.capturedPieces = append(b.capturedPieces, moveResult.CapturedPiece())
	}

	return moveResult, nil
}

func (b *board) UndoLastMove() (chess.MoveResult, error) {
	return nil, nil
}

func (b *board) MarshalJSON() ([]byte, error) {
	type Placement struct {
		Piece      chess.Piece       `json:"piece"`
		Position   chess.Position    `json:"position"`
		LegalMoves chess.PositionSet `json:"legal_moves"`
	}

	placements := make([]*Placement, 0, 32)
	for pos, piece := range b.squares.Iter() {
		if piece == nil {
			continue
		}

		placement := &Placement{Piece: piece, LegalMoves: mapset.NewSet[chess.Position](), Position: pos}
		if piece.Side() == b.turn {
			placement.LegalMoves = b.LegalMoves(piece)
		}

		placements = append(placements, placement)
	}

	return json.Marshal(map[string]any{
		"turn":           b.turn,
		"state":          b.State(b.turn),
		"castlings":      metric.CastlingAbility(b).Value().(metric.Castlings)["practical"][b.turn],
		"capturedPieces": b.capturedPieces,
		"move_history":   b.moveHistory,
		"placements":     placements,
	})
}
