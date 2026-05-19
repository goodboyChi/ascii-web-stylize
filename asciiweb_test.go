package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"web/asciiArt"
)

func TestAscii(t *testing.T) {

result := asciiArt.Ascii("A", "standard")
 if result == "" {
	t.Error("expected output, got")
 }

}

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	homeHandler(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}
}
	


