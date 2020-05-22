package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ersonp/go-resume/common"
	"github.com/ersonp/go-resume/middlewares"
)

// Test that a GET request to the home page returns 200
func TestHomePage(t *testing.T) {
	r := middlewares.GetRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/", ShowHomePage)

	req, _ := http.NewRequest("GET", "/index", nil)
	common.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home</title>") > 0

		return statusOK && pageOK
	})
}
