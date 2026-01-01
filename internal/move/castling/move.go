package castling

import (
	"regexp"

	"github.com/elaxer/rgx"
)

const (
	TypeShort CastlingType = true
	TypeLong  CastlingType = false
)

var regexpCastling = regexp.MustCompile("^[0Oo]-[0Oo](?P<long>-[0Oo])?[#+]?$")

type CastlingType bool

func TypeFromString(str string) (CastlingType, error) {
	result, err := rgx.Group(regexpCastling, str)
	if err != nil {
		return false, err
	}

	return CastlingType(result["long"] == ""), nil
}

func (m CastlingType) Validate() error {
	return nil
}

func (m CastlingType) IsShort() bool {
	return m == TypeShort
}

func (m CastlingType) IsLong() bool {
	return m == TypeLong
}

func (m CastlingType) String() string {
	return map[CastlingType]string{
		TypeShort: "O-O",
		TypeLong:  "O-O-O",
	}[m]
}
