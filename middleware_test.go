package traefikuseragent_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	mw "github.com/GiGInnovationLabs/traefikuseragent"
)

func TestGeoIPBasic(t *testing.T) {
	mwCfg := mw.CreateConfig()

	called := false
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) { called = true })

	instance, err := mw.New(context.TODO(), next, mwCfg, "traefikuseragent")
	if err != nil {
		t.Fatalf("Error creating %v", err)
	}

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)

	instance.ServeHTTP(recorder, req)
	if recorder.Result().StatusCode != http.StatusOK {
		t.Fatalf("Invalid return code")
	}
	if called != true {
		t.Fatalf("next handler was not called")
	}
}
