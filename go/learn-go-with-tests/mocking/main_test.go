package main

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	d = "sleep"
	w = "write"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, d)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, w)
	return
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spy := &SpyCountdownOperations{}

	Countdown(buffer, spy)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	t.Run("sleep before every print", func(t *testing.T) {
		spyPrint := &SpyCountdownOperations{}

		Countdown(spyPrint, spyPrint)

		want := []string{
			w,
			d,
			w,
			d,
			w,
			d,
			w,
		}

		if !reflect.DeepEqual(want, spyPrint.Calls) {
			t.Errorf("wanted calls %v got %v", want, spyPrint.Calls)
		}
	})
}
