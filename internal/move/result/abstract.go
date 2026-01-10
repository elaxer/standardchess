// Package result contains abstract move result structure
package result

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/state"
)

var ErrValidation = errors.New("abstract result validation error")

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
		return fmt.Errorf("%w: empty new state", ErrValidation)
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
