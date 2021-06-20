package clientapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHealthService_Check(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v1/health", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"up"}`)
	})

	got, err := client.Health.Check()
	if err != nil {
		t.Errorf("Health.Check returned error:\n%v", err)
	}

	want := &Status{Status: "up"}
	if !cmp.Equal(got, want) {
		t.Errorf("Health.Check returned %+v, want %+v", got, want)
	}
}
