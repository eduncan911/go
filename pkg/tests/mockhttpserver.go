package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// MockHTTPServer represents a mock net/http.HTTPServer
type MockHTTPServer struct {
	ContentType string
	StatusCode  int
}

// NewMockHTTPServer returns a mock instance of MockHTTPServer.  It wraps the funkiness around
// mocking up Go's concrete instance of internals for Http methods.
//
// 	package tests
//
// 	import (
// 		"net/http"
// 		"strings"
// 		"testing"
//
// 		"github.com/eduncan911/gomspec/assert"
// 	)
//
// 	// sample code that would be part of your app/pkg normally
// 	// and you will want to test
// 	var myFunc = func(url string) bool {
// 		resp, _ := http.Get(url)
// 		defer resp.Body.Close()
// 		if resp.StatusCode == 404 {
// 			return true
// 		}
// 		return false
// 	}
//
// 	func TestMockHTTPServerSample(t *testing.T) {
//
// 		// arrange
// 		json := `{error:{StatusCode:404,ErrorMessage:"Nothing was found"}}`
// 		options := func(ts *MockHTTPServer) {
// 			ts.ContentType = "application/json"
// 			ts.StatusCode = 404
// 		}
// 		ts, _ := NewMockHTTPServer(json, options)
// 		url := strings.Join([]string{ts.URL, "/status"}, "") // the key is using the ts.URL as your base
//
// 		// act
// 		resultFromRemoteCode := myFunc(url)
//
// 		// assert
// 		assert.True(t, resultFromRemoteCode)
// 	}
func NewMockHTTPServer(resultToReturn string, options ...func(*MockHTTPServer)) (*httptest.Server, error) {

	// defaults
	srv := MockHTTPServer{
		ContentType: "text/plain",
		StatusCode:  200,
	}

	// set the options
	for _, option := range options {
		option(&srv)
	}

	// setup the mocked Go httptest server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", srv.ContentType)
		w.WriteHeader(srv.StatusCode)
		fmt.Fprint(w, resultToReturn)
	}))

	return ts, nil
}
