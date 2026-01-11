package standardchess

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/state"
)

var (
	// StateClear means a clear default state of the chess board.
	StateClear = chess.StateClear

	// StateCheck means there is a check on the board.
	StateCheck = state.Check

	// StateCheckmate means there is a checkmate on the board.
	StateCheckmate = state.Checkmate
	// StateStalemate means there is a stalemate on the board.
	StateStalemate = state.Stalemate

	// StateFiftyMoves is used if fifty consecutive moves occur without a pawn move or any capture.
	StateFiftyMoves = state.FiftyMoves
	// StateThreefoldRepetition means a case when the same position occurs three times
	// (same player to move and same rights).
	StateThreefoldRepetition = state.ThreefoldRepetition
	// StateInsufficientMaterial means a draw when neither side has enough material to checkmate
	// (e.g., king vs king, king and bishop vs king).
	StateInsufficientMaterial = state.InsufficientMaterial
)
