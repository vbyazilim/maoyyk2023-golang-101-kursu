package kvstore_test

import (
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/14/test-datarace/kvstore"
)

func TestDataRace(t *testing.T) {
	st := make(map[string]string)
	done := make(chan struct{})

	s := kvstore.New(st)
	_ = s.Set("foo", "bar")

	go func() {
		_ = s.Set("foo", "data race...")
		done <- struct{}{}
	}()

	want := "bar"
	got, _ := s.Get("foo") // always returns "bar"
	<-done                 // after line 19, blocking ends... map changes but doesn't affect got variable!

	// fmt.Println(s.Get("foo")) data race... <nil>

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}
