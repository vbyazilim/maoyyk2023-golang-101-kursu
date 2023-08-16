package greet_test

import (
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-table-driven-sub-tests/greet"
)

func TestSayHi(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  string
	}{
		"run with single arg":    {input: []string{"vigo"}, want: "hi vigo!"},
		"run with multiple args": {input: []string{"vigo", "turbo"}, want: "hi vigo!\nhi turbo!"},
		"run with no arg":        {input: []string{}, want: "hi everybody!"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet.SayHi(tc.input...)

			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
	}
}
