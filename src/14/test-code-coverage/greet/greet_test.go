package greet_test

import (
	"fmt"
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-code-coverage/greet"
)

func TestSayHi(t *testing.T) {
	want := "hi vigo!"
	got := greet.SayHi("vigo")

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestSayHiWithNoArgs(t *testing.T) {
	want := "hi everybody!"
	got := greet.SayHi()

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestSayHiWithArgs(t *testing.T) {
	want := "hi vigo!\nhi turbo!\nhi max!"
	got := greet.SayHi("vigo", "turbo", "max")

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

// This is example of single argument usage.
func ExampleSayHi() {
	fmt.Println(greet.SayHi("vigo"))
	// Output: hi vigo!
}

// This is example of no argument usage.
func ExampleSayHi_withNoArg() {
	fmt.Println(greet.SayHi())
	// Output: hi everybody!
}

// This is example of with many arguments usage.
func ExampleSayHi_withArgs() {
	fmt.Println(greet.SayHi("vigo", "turbo", "max"))
	// Output: hi vigo!
	// hi turbo!
	// hi max!
}
