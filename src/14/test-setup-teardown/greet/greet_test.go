package greet_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-setup-teardown/greet"
)

func TestMain(m *testing.M) {
	fmt.Println("do setup operations...")
	_ = os.Setenv("CUSTOM_HOST", "localhost")
	_ = os.Setenv("CUSTOM_PORT", "9000")

	result := m.Run()

	fmt.Println("do teardown operations...")
	_ = os.Unsetenv("CUSTOM_HOST")
	_ = os.Unsetenv("CUSTOM_PORT")

	os.Exit(result)
}

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
		// <setup code>
		fmt.Println("setup code from sub test initiated!")
		t.Run(name, func(t *testing.T) {
			if val, ok := os.LookupEnv("CUSTOM_PORT"); ok && val == "9000" {
				fmt.Println("using port 9000")
			}
			got := greet.SayHi(tc.input...)

			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
		// <tear-down code>
		fmt.Println("teardown code from sub test initiated!")
	}
}
