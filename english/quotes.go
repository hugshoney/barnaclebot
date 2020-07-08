// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package english

import (
	"encoding/json"
	"math/rand"
	"time"
)

func Quotes() map[string]string {
	data := GetData("https://type.fit/api/quotes", "")

	var jsonResult []map[string]string
	json.Unmarshal(data, &jsonResult)

	// Create random seed using unix time.
	rand.Seed(time.Now().Unix())
	// Select random quotes text from list all quotes.
	quotes := jsonResult[rand.Intn(len(jsonResult))]

	return quotes
}
