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

type Definition struct {
	Define  string `json:"definition"`
	Example string `json:"example"`
}

type Response struct {
	List []Definition `json:"list"`
}

// Get definition and example of slang word.
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

	if len(jsonResult.List) == 0 {
		definition = fmt.Sprintf("%s is not found, try another day.", word)
		example = ""
	} else {
		definition = jsonResult.List[0].Define
		example = jsonResult.List[0].Example
	}

	return
}
