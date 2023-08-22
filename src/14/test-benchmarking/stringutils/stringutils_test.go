package stringutils_test

import (
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-benchmarking/stringutils"
)

func TestReverse(t *testing.T) {
	want := "mişey"
	got := stringutils.Reverse("yeşim")

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestReverseVigo(t *testing.T) {
	want := "ogiv"
	got := stringutils.ReverseVigo("vigo")

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

var gs string

func BenchmarkReverse(b *testing.B) {
	var s string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = stringutils.Reverse("aklındaysa kapında!")
	}
	gs = s
}

func BenchmarkReverseVigo(b *testing.B) {
	var s string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = stringutils.ReverseVigo("aklındaysa kapında!")
	}
	gs = s
}
