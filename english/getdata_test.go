// Any copyright is dedicated to the Public Domain.
// https://creativecommons.org/publicdomain/zero/1.0/

package english

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetData(t *testing.T) {
	jsonData := []byte(`{"It's just a test"}`)

	serverAPI := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(jsonData)
			w.WriteHeader(http.StatusOK)
		}))
	defer serverAPI.Close()
	got := GetData(serverAPI.URL, "")

	if !reflect.DeepEqual(got, jsonData) {
		t.Errorf("got %v want %v", got, jsonData)
	}
}
