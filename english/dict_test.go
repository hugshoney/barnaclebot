// Any copyright is dedicated to the Public Domain.
// https://creativecommons.org/publicdomain/zero/1.0/

package english

import (
	"reflect"
	"testing"
)

// Test for Mean function.
func TestMean(t *testing.T) {
	// Get dictionary result (speech, definition, and example)
	// for 'homeless' word.
	got := Mean("homeless")
	want := []map[string]string{
		{
			"speech":     "adjective",
			"definition": "(of a person) without a home, and therefore typically living on the streets.",
			"example":    "the plight of young homeless people",
		},
	}

	// Compare slice between what I got and what I want.
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
