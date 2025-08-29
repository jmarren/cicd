package internal

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tests = map[string]string{
	"/":   "hi there",
	"/hi": "howdy",
}

func testRoute(t *testing.T, path string, want string) {

	res, err := http.Get(path)
	if err != nil {
		t.Fatalf("failed to GET / with error: %s\n", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("failed to read res.Body with err: %s\n", err)
	}
	if string(body) != want {
		t.Fatalf("\n|wanted: %s\n|got: %s", want, body)
	}

	fmt.Println(string(body))

}

func TestServer(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer ts.Close()

	t.Logf("getting url: %s\n", ts.URL)

	for path, want := range tests {
		testRoute(t, ts.URL+path, want)
	}

}
