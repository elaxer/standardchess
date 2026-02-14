// Package result contains abstract move result structure
package result

import (
	"errors"
	"fmt"

	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/state"
)

var (
	ErrValidation = errors.New("abstract result validation error")

	errValidationEmptyNewState = fmt.Errorf("%w: empty new state", ErrValidation)
)

type Abstract struct {
	MoveSide chess.Color
	NewState chess.State
	IsCheck  bool
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
		return errValidationEmptyNewState
	}

	return nil
}

func (r *Abstract) Suffix() string {
	if r.NewState == state.Checkmate {
		return "#"
	}
	if r.IsCheck {
		return "+"
	}

	return ""
}
