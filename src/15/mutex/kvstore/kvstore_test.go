package kvstore_test

import (
	"testing"

	"github.com/vbyazilim/maoyyk2023-golang-101-kursu/src/15/mutex/kvstore"
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
	got, _ := s.Get("foo")
	<-done

	if got != want {
		t.Errorf("want: %v; got: %v", want, got)
	}
}
