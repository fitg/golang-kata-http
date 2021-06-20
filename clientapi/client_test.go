package clientapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the client being tested and is
	// configured to use test server.
	client = CreateClient()
	url, _ := url.Parse(server.URL + "/")
	client.BaseURL = url

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestCreateClient(t *testing.T) {
	c := CreateClient()

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("CreateClient BaseURL is %v, want %v", got, want)
	}
	if got, want := c.RequestEncoding, defaultRequestEncoding; got != want {
		t.Errorf("CreateClient RequestEncoding is %v, want %v", got, want)
	}

	c2 := CreateClient()
	if c.HTTPClient == c2.HTTPClient {
		t.Error("CreateClient returned same http.Clients, but they should differ")
	}
}

func TestSendRequest(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v1/health", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"up"}`)
	})

	got, err := client.sendRequest("GET", fmt.Sprintf("%v%v", client.BaseURL, "/v1/health"), HealthOK, nil)
	if err != nil {
		t.Errorf("client.sendRequest returned error:\n%v", err)
	}

	want := []byte(`{"status":"up"}`)
	if !cmp.Equal(reflect.TypeOf(got).String(), reflect.TypeOf(want).String()) {
		t.Errorf("client.sendRequest returned %+v, want %+v", got, want)
	}
}
