// Any copyright is dedicated to the Public Domain.
// https://creativecommons.org/publicdomain/zero/1.0/

package english

import "testing"

func TestQuotes(t *testing.T) {
	// Call quotes function and get random quotes.
	result := Quotes()
	// Check if quotes text exist or not in result.
	if _, exist := result["text"]; !exist {
		t.Errorf("key for quotes text is not found in result.")
	}
}
