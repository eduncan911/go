package tests

import (
	"net/http"
	"strings"
	"testing"

	"github.com/eduncan911/gomspec/assert"
)

// sample code that would be part of your app/pkg normally
// and you will want to test
var myFunc = func(url string) bool {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return true
	}
	return false
}

func TestMockHTTPServerSample(t *testing.T) {

	// arrange
	json := `{error:{StatusCode:404,ErrorMessage:"Nothing was found"}}`
	options := func(ts *MockHTTPServer) {
		ts.ContentType = "application/json"
		ts.StatusCode = 404
	}
	ts, _ := NewMockHTTPServer(json, options)
	url := strings.Join([]string{ts.URL, "/status"}, "") // the key is using the ts.URL as your base

	// act
	resultFromRemoteCode := myFunc(url)

	// assert
	assert.True(t, resultFromRemoteCode)
}

func TestMockHTTPServerDefaults(t *testing.T) {

	// arrange
	json := `{}`
	ts, _ := NewMockHTTPServer(json)

	// act
	r, _ := http.Get(ts.URL)

	// assert
	assert.Equal(t, 200, r.StatusCode, "The default status code should have been 200")
	assert.Equal(t, "text/plain", r.Header.Get("content-type"), "The default content-type should have been text/plain")
}
