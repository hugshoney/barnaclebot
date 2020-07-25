// Any copyright is dedicated to the Public Domain.
// https://creativecommons.org/publicdomain/zero/1.0/

package english

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// Test for GetData
func TestGetData(t *testing.T) {
	// Create data for body.
	data := []byte(`{"It's just a test"}`)

	// Create server for testing.
	serverTest := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(data)
			w.WriteHeader(http.StatusOK)
		}))
	// Close server in the end of function.
	defer serverTest.Close()
	// Use testing server url as paramater for GetData.
	got := GetData(serverTest.URL, "")

	// Check if is equal or not with data.
	if !reflect.DeepEqual(got, data) {
		t.Errorf("got %v want %v", got, data)
	}
}
