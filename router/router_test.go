package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ARtorias742/matcher"
)

func TestRouter_PathParams(t *testing.T) {
	router := NewRouter()
	router.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		params := matcher.GetParams(r.Context())
		w.Write([]byte("User ID: " + params["id"]))
	})

	req, _ := http.NewRequest("GET", "/users/123", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
	if rr.Body.String() != "User ID: 123" {
		t.Errorf("expected body 'User ID: 123', got '%s'", rr.Body.String())
	}
}

func TestRouter_Middleware(t *testing.T) {
	router := NewRouter()
	mw := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("MW "))
			next(w, r)
		}
	}
	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}).Use(mw)

	req, _ := http.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Body.String() != "MW Hello" {
		t.Errorf("expected 'MW Hello', got '%s'", rr.Body.String())
	}
}

func TestRouter_HostRouting(t *testing.T) {
	router := NewRouter()
	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("example.com"))
	}).Host("example.com")

	req, _ := http.NewRequest("GET", "/test", nil)
	req.Host = "example.com"
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Body.String() != "example.com" {
		t.Errorf("expected 'example.com', got '%s'", rr.Body.String())
	}
}

func TestRouter_NamedRoute(t *testing.T) {
	router := NewRouter()
	router.Get("/home", func(w http.ResponseWriter, r *http.Request) {}).Name("home")
	if route := router.FindByName("home"); route == nil {
		t.Error("expected to find route 'home', got nil")
	}
}
