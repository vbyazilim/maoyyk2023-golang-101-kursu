package stringutils_test

import (
	"fmt"
	"testing"
)

func BenchmarkSprintConcat(b *testing.B) {
	b.Run("sprint", benchSprint) // sub test gibi, sub benchmark!
	b.Run("concat", benchConcat)
}

func benchSprint(b *testing.B) {
	var s string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = fmt.Sprint("hello") // nolint:gosimple
	}

	gs = s // bunu yapmazsak allocation
}

func benchConcat(b *testing.B) {
	var s string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = "hello" + "world"
	}

	gs = s
}
