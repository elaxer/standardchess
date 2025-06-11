package pgn

import (
	"regexp"

	"github.com/elaxer/chess/encoding/pgn"
)

var (
	RegexpHeaders = regexp.MustCompile(`\[(?P<name>[\w]+)\s+"(?P<value>[^"]*)"\]`)
	RegexpMoves   = regexp.MustCompile("([NBKRQ]?[a-h]?[1-8]?x?[a-h][1-8](?:=[NBRQ])?)|([0Oo]-[0Oo](-[0Oo])?)")
)

func NewDecoder() *pgn.Decoder {
	return pgn.NewDecoder(RegexpHeaders, RegexpMoves)
}
