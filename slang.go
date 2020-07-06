// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Structuring JSON for definition and example of slang word.
type Definition struct {
	Define  string `json:"definition"`
	Example string `json:"example"`
}

// Structuring JSON for list of definition and example result.
type Response struct {
	List []Definition `json:"list"`
}

// Get definition and example of slang word from UrbanDictionary API.
func slang(word string) (definition, example string) {
	url := fmt.Sprint("http://api.urbandictionary.com/v0/define?term=", word)

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var jsonResult Response
	json.Unmarshal(body, &jsonResult)

	// Give information if slang word not found in result.
	// I'm using len because I don't know to handle empty struct or
	// maybe it's just nil slice, I'm confused -_-.
	if len(jsonResult.List) == 0 {
		definition = fmt.Sprintf("Definition of \"%s\" is not found, try another day.", word)
		example = ""
	} else {
		definition = jsonResult.List[0].Define
		example = jsonResult.List[0].Example
	}

	return
}
