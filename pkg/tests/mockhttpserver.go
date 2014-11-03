package mocks

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

// NewMockHTTPServer returns a mock instance of MockHTTPServer.
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
