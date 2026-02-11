// Package repetition provides utilities for tracking repeated board
// positions.
//
// It maintains a history of position hashes and allows querying how
// many times the current position has occurred. This can be used to
// implement threefold or fivefold repetition rules.
package repetition

import "github.com/elaxer/chess"

type Repetition struct {
	hashes []string
}

func NewRepetition() *Repetition {
	return &Repetition{make([]string, 0)}
}

func (r *Repetition) CurrentRepetitions() int {
	if len(r.hashes) == 0 {
		return 0
	}

	lastHash := r.hashes[len(r.hashes)-1]
	repetitions := 0

	for _, h := range r.hashes {
		if h == lastHash {
			repetitions++
		}
	}

	return repetitions
}

func (r *Repetition) AfterBoardMakeMove(board chess.Board) {
	r.hashes = append(r.hashes, hash(board))
}

func (r *Repetition) AfterBoardUndoMove() {
	if len(r.hashes) > 0 {
		r.hashes = r.hashes[:len(r.hashes)-1]
	}
}
