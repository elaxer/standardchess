package pgn

import "fmt"

// Header represents a PGN header with a name and value.
// It is used to store metadata about the chess game, such as event, site, date, etc.
// The header is formatted as "[name "value"]" in the PGN string.
type Header struct {
	// Name is the name of the header, such as "Event", "Site", "Date", etc.
	Name string
	// Value is the value of the header, which can be a string containing information about the game.
	Value string
}

func NewHeader(name, value string) Header {
	return Header{name, value}
}

// String formats the header as a PGN string.
// It returns a string in the format "[name "value"]".
func (h Header) String() string {
	return fmt.Sprintf("[%s \"%s\"]", h.Name, h.Value)
}
