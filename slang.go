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

func slang(word string) Response {
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

	return jsonResult
}

func slangDef(word string) string {
	result := slang(word)
	top := result.List[0].Define

	return top
}

func slangEg(word string) string {
	result := slang(word)
	top := result.List[0].Example

	return top
}
