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
//  func Test404NotFound(t *testing.T) {
//
//    // arrange
//    json := `{error:{StatusCode:404,ErrorMessage:"Nothing was found"}}`
//    opts := func(ts *tests.MockHTTPServer) {
//      ts.ContentType = "application/json"
//      ts.StatusCode = 404
//    }
//    ts, _ := tests.NewMockHTTPServer(json, opts)
//
//    // act
//    fn := func() bool {
//      // sample code that would be part of your app/pkg normally
//      resp := http.Get("/status")
//      defer resp.Close()
//      if resp.StatusCode == 404 { return true }
//      return false
//    }
//    resultFromRemoteCode := fn()
//
//    // assert
//    assert.True(t, resultFromRemoteCode)
//
//  }
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
