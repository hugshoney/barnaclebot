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

	definition = jsonResult.List[0].Define
	example = jsonResult.List[0].Example
	return
}
