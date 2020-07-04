package main

import (
	"testing"
)

func TestSlang(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	def, eg := slang("thot")

	t.Run("Get definition of 'thot'", func(t *testing.T) {
		got := def
		want := "Pronounced \\ˈthȯt\\ and taken from THree - One - Two. The original version of THOT before someone came to think it meant something else. It was brought to you by THOTCON, a hacking [conference] based in Chicago [IL], USA which started in 2010.\n\nFor those that aren't [believers], check out the thotcon website or wikipedia."

		assertCorrectMessage(t, got, want)
	})

	t.Run("Get example of 'thot'", func(t *testing.T) {
		got := eg
		want := "Damn, THOT CON is absolutely [bad-ass].\n\nThose THOT [IES] really are some of the best [hackers] in the world."

		assertCorrectMessage(t, got, want)
	})
}
