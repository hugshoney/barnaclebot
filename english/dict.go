// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package english

import "encoding/json"

// Struct for word meaning and example.
type Word struct {
	Mean    string `json:"definition"`
	Example string `json:"Example"`
}

// Struct for items in result.
type Dict struct {
	Speech      string `json:"partOfSpeech"`
	Definitions []Word `json:"definitions"`
}

// Struct to hold response when calling API.
type DictResponse struct {
	Meanings []Dict `json:"meanings"`
}

// Get word meaning, example, and part of speech from word.
func Mean(word string) []Dict {
	data := GetData("https://api.dictionaryapi.dev/api/v2/entries/en/", word)

	var jsonResult []DictResponse
	json.Unmarshal(data, &jsonResult)

	// Return only Dict as result.
	result := jsonResult[0].Meanings

	return result
}
