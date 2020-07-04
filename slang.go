package main

import (
	"encoding/json"
    "fmt"
	"io/ioutil"
	"net/http"
)

type Slang struct {
	Definition string `json:"definition"`
	Example string `json:"example"`
}

type SlangResult struct {
    List []Slang `json:"list"`
}

func GetSlang(word string) SlangResult {
    url := fmt.Sprint("http://api.urbandictionary.com/v0/define?term=",  word)
    
    res, err := http.Get(url)
    if err != nil {
        panic(err.Error())
    }
    
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err.Error())
    }
    
    var jsonResult SlangResult
    json.Unmarshal(body, &jsonResult)
    
    return jsonResult
}

func SlangDef(word string) string {
    result := GetSlang(word)
    top := result.List[0].Definition
    
    return top
}

func SlangEg(word string) string {
    result := GetSlang(word)
    top := result.List[0].Example
    
    return top
}
