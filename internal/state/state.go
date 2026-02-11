// Package state contains a set of variables representing the state of the boards.
package state

import "github.com/elaxer/chess"

var (
	Checkmate = chess.NewState("checkmate", true)
	Stalemate = chess.NewState("stalemate", true)

	FiftyMoves           = chess.NewState("fifty moves rule", true)
	FivefoldRepetition   = chess.NewState("fivefold repetition", true)
	InsufficientMaterial = chess.NewState("insufficient material", true)
)
