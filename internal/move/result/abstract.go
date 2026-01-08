package result

import (
	"errors"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/state/state"
)

type Abstract struct {
	MoveSide chess.Color
	NewState chess.State
}

func (r *Abstract) Side() chess.Color {
	return r.MoveSide
}

func (r *Abstract) SetBoardNewState(state chess.State) {
	r.NewState = state
}

func (r *Abstract) BoardNewState() chess.State {
	return r.NewState
}

func (r *Abstract) Validate() error {
	if r.NewState == nil {
		// todo
		return errors.New("dfoios")
	}

	return nil
}

func (r *Abstract) Suffix() string {
	switch r.NewState {
	case state.Check:
		return "+"
	case state.Checkmate:
		return "#"
	}

	return ""
}
