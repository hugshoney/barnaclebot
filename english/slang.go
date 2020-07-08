// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package english

import (
	"encoding/json"
	"fmt"
)

// Structuring JSON for definition and example of slang word.
type SlangWord struct {
	Define  string `json:"definition"`
	Example string `json:"example"`
}

// Structuring JSON for list of definition and example result.
type SlangResponse struct {
	List []SlangWord `json:"list"`
}

// Get definition and example of slang word from UrbanDictionary API.
func Slang(word string) (definition, example string) {
	data := GetData("http://api.urbandictionary.com/v0/define?term=", word)

	var jsonResult SlangResponse
	json.Unmarshal(data, &jsonResult)

	// Give information if slang word not found in result.
	// I'm using len because I don't know to handle empty struct or
	// maybe it's just nil slice, I'm confused -_-.
    // TODO i'ts really bad handling, refactor please.
	if len(jsonResult.List) == 0 {
		definition = fmt.Sprintf("Definition of \"%s\" is not found, try another day.", word)
		example = ""
	} else {
		definition = jsonResult.List[0].Define
		example = jsonResult.List[0].Example
	}

	return
}
