// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package english

import (
	"io/ioutil"
	"net/http"
)

func GetData(url, word string) []byte {
	// If API need input word as argument.
	if word != "" {
		// Concatination word with url as
		// full address.
		url += word
	}
	// Take response from API url.
	response, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	// Get body of data from API response.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	// Return body of data from API.
	return body
}
