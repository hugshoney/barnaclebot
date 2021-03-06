// Any copyright is dedicated to the Public Domain.
// https://creativecommons.org/publicdomain/zero/1.0/

package english

import (
	"testing"
)

// Test for slang function.
func TestSlang(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// Call slang function to return top definition (def)
	// and example (eg).
	slang := Slang("thot")

	// Test to know if got definition is same with what I want.
	t.Run("Get definition of 'thot'", func(t *testing.T) {
		got := slang["definition"]
		want := "Pronounced \\ˈthȯt\\ and taken from THree - One - Two. The original version of THOT before someone came to think it meant something else. It was brought to you by THOTCON, a hacking [conference] based in Chicago [IL], USA which started in 2010.\n\nFor those that aren't [believers], check out the thotcon website or wikipedia."

		assertCorrectMessage(t, got, want)
	})

	// Test to know if got example is same with what I want.
	t.Run("Get example of 'thot'", func(t *testing.T) {
		got := slang["example"]
		want := "Damn, THOT CON is absolutely [bad-ass].\n\nThose THOT [IES] really are some of the best [hackers] in the world."

		assertCorrectMessage(t, got, want)
	})

	// Test to know how slang function handle not found definition
	// of slang word that I want to know.
	t.Run("Get no result for 'awokwok'", func(t *testing.T) {
		// Call slang function to find definition of 'awokwok'
		notSlang := "awokwok"
		slang := Slang(notSlang)
		if len(slang) != 0 {
			t.Errorf("%s is not slang", notSlang)
		}
	})
}
