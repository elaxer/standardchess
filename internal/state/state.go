// Package state contains a set of variables representing the state of the boards.
package state

import "github.com/elaxer/chess"

var (
	Check = chess.NewState("check", chess.StateTypeThreat)

	Checkmate = chess.NewState("checkmate", chess.StateTypeTerminal)
	Stalemate = chess.NewState("stalemate", chess.StateTypeTerminal)

	FiftyMoves           = chess.NewState("fifty moves rule", chess.StateTypeClear)
	ThreefoldRepetition  = chess.NewState("threefold repetition", chess.StateTypeClear)
	InsufficientMaterial = chess.NewState("insufficient material", chess.StateTypeTerminal)
)
