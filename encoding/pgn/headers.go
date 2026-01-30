package pgn

type Headers []Header

func (h Headers) Get(name string) (Header, bool) {
	for _, header := range h {
		if header.Name == name {
			return header, true
		}
	}

	return Header{}, false
}
