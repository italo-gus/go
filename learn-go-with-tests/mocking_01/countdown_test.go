package main // Mocking

import (
	"bytes"
	"testing"
)

/*
func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	received := buffer.String()
	//expected := "3"

	expected := `3
2
1
Go!`

	if received != expected {
		t.Errorf("received %q expected %q", received, expected)
	}
}
*/

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		//t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
		t.Errorf("não há chamadas suficientes para o sleeper, quero 3 e recebi %d", spySleeper.Calls)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
