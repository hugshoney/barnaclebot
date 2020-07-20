// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package english

import (
	"encoding/json"
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
func Slang(word string) map[string]string {
	data := GetData("http://api.urbandictionary.com/v0/define?term=", word)

	var jsonResult SlangResponse
	json.Unmarshal(data, &jsonResult)

	// Make empty map for initial result.
	var result = make(map[string]string)
	// Add definition and example to map if there is result from API.
	if len(jsonResult.List) != 0 {
		result["definition"] = jsonResult.List[0].Define
		result["example"] = jsonResult.List[0].Example
	}

	return result
}
