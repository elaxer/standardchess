package state

import "github.com/elaxer/chess"

var (
	Check     = chess.NewState("check", chess.StateTypeThreat)
	Checkmate = chess.NewState("checkmate", chess.StateTypeTerminal)

	Stalemate      = chess.NewState("stalemate", chess.StateTypeDraw)
	DrawFiftyMoves = chess.NewState("draw by fifty moves", chess.StateTypeDraw)
)
