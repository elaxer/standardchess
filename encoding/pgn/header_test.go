package pgn

import "testing"

func TestHeader_String(t *testing.T) {
	type fields struct {
		Name  string
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"valid",
			fields{"Event", "FIDE World Championship"},
			`[Event "FIDE World Championship"]`,
		},
		{
			"empty_value",
			fields{"Event", ""},
			`[Event ""]`,
		},
		{
			"empty_name",
			fields{"", "FIDE World Championship"},
			`[ "FIDE World Championship"]`,
		},
		{
			"empty_both",
			fields{"", ""},
			`[ ""]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHeader(tt.fields.Name, tt.fields.Value)
			if got := h.String(); got != tt.want {
				t.Errorf("Header.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
