package greet_test

import (
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-parallel/greet"
)

// TestSayHi will not complete until all parallel tests started by Run have completed.
// As a result, no other parallel tests can run in parallel to these parallel tests.
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
		tc := tc // capture range variable to ensure that
		// tc gets bound to the correct instance.

		t.Run(name, func(t *testing.T) {
			t.Parallel() // run in parallel

			got := greet.SayHi(tc.input...)
			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
	}
}
