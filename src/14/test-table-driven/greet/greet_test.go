package greet_test

import (
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-table-driven/greet"
)

func TestSayHi(t *testing.T) {
	type test struct {
		testName string
		input    []string
		want     string
	}

	tests := []test{
		{testName: "run with single arg", input: []string{"vigo"}, want: "hi vigo!"},
		{testName: "run with multiple args", input: []string{"vigo", "turbo"}, want: "hi vigo!\nhi turbo!"},
		{testName: "run with no arg", input: []string{}, want: "hi everybody!"},
	}

	// with anonymous struct
	// tests := []struct {
	// 	testName string
	// 	input    []string
	// 	want     string
	// }{
	// 	{testName: "run with single arg", input: []string{"vigo"}, want: "hi vigo!"},
	// 	{testName: "run with multiple args", input: []string{"vigo", "turbo"}, want: "hi vigo!\nhi turbo!"},
	// 	{testName: "run with no arg", input: []string{}, want: "hi everybody!"},
	// }

	for _, tc := range tests {
		got := greet.SayHi(tc.input...)

		if got != tc.want {
			t.Errorf("[%s]: want: %v; got: %v", tc.testName, tc.want, got)
		}
	}
}
