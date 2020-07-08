// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package english

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Struct for word meaning and example.
type Word struct {
	Mean    string `json:"definition"`
	Example string `json:"Example"`
}

// Struct for items in result.
type DictResult struct {
	Speech      string `json:"partOfSpeech"`
	Definitions []Word `json:"definitions"`
}

// Struct to hold response when calling API.
type DictResponse struct {
	Meanings []DictResult `json:"meanings"`
}

// Function to get word meaning, example, and part of speech.
func Mean(word string) []map[string]string {
	// TODO it's need to written in somewhere else.
	// It's already written so many times, I need function to call this thing.
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word
	response, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	var jsonResult []DictResponse
	json.Unmarshal(body, &jsonResult)

	// Variable to hold list result of definition,
	// example, and speech. Probably I will add synonyms.
	var result []map[string]string
	// If definition of word found in Dictionary,
	// let's process it to be nice slice of map.
	// TODO it's not really nice slice, btw. Refactor please.
	if len(jsonResult) != 0 {
		var speech, definition, example string
		// Loops list items in dictionary result. I know
		// it's looks stupid, but it's work!.
		for _, item := range jsonResult[0].Meanings {
			speech = item.Speech
			// Sometimes some speech have more than one
			// of definition and example, I don't know
			// why this thing can happen, but this is
			// what I got from API, so I just process it.
			for _, word := range item.Definitions {
				definition = word.Mean

				// If example usage of word exist, let's
				// add it to map.
				// This code looks real bad, it's need rewriten.
				if word.Example != "" {
					example = word.Example

					result = append(result, map[string]string{
						"speech":     speech,
						"definition": definition,
						"example":    example,
					})
				} else {
					result = append(result, map[string]string{
						"speech":     speech,
						"definition": definition,
					})
				}
			}
		}
	}

	return result
}
