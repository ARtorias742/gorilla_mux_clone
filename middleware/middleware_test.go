package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}

	mwHandler := Logging(handler)

	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()
	mwHandler.ServeHTTP(rr, req)

	if rr.Body.String() != "OK" {
		t.Errorf("expected 'OK', got '%s'", rr.Body.String())
	}

}
