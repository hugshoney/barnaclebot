// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package english

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func Quotes() map[string]string {
	url := "https://type.fit/api/quotes"
	response, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	var jsonResult []map[string]string
	json.Unmarshal([]byte(body), &jsonResult)

	// Create random seed using unix time.
	rand.Seed(time.Now().Unix())
	// Select random quotes text from list all quotes.
	quotes := jsonResult[rand.Intn(len(jsonResult))]

	return quotes
}
