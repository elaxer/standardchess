package rule

import (
	"github.com/elaxer/chess"
	"github.com/elaxer/standardchess/internal/repetition"
	"github.com/elaxer/standardchess/internal/state"
)

type repetitionRule struct {
	*repetition.Repetition
	repetitionsNum int
	stateValue     chess.State
}

func NewThreefoldRepetition(repetition *repetition.Repetition) *repetitionRule {
	return &repetitionRule{repetition, 3, state.ThreefoldRepetition}
}

func NewFivefoldRepetition(repetition *repetition.Repetition) *repetitionRule {
	return &repetitionRule{repetition, 5, state.FivefoldRepetition}
}

func (r *repetitionRule) Rule(board chess.Board) chess.State {
	if r.CurrentRepetitions() < r.repetitionsNum {
		return nil
	}

	return r.stateValue
}
