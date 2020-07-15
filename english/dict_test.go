// Any copyright is dedicated to the Public Domain.
// https://creativecommons.org/publicdomain/zero/1.0/

package english

import (
	"reflect"
	"testing"
)

// Test for Dict function.
func TestDict(t *testing.T) {
	// Get dictionary result like speech, definition, example,
	// and synonym for 'homeless' word.
	got := Dict("homeless")

	// the result that I want.
	want := []Mean{
		{
			Speech: "adjective",
			Definitions: []Word{
				{
					Mean:    "(of a person) without a home, and therefore typically living on the streets.",
					Example: "the plight of young homeless people",
					Synonym: []string{
						"without a roof over one's head",
						"on the streets",
						"vagrant",
						"sleeping rough",
						"living rough",
					},
				},
			}},
	}

	// Compare slice between what I got and what I want.
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
