package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TesCreateObject(t *testing.T) {
	// db.Get()
	req, err := http.NewRequest("GET", "/api/CreateObject?content=1&folder=image/&key=a/image/1.jpg",
		nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-WX-OPENID", "111111")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateObject)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Hello, world!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSearch(t *testing.T) {

// db.Get()
	req, err := http.NewRequest("GET", "/api/ObjectList?query=1&start=1&size=10",
		nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-WX-OPENID", "111111")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ObjectList)

	handler.ServeHTTP(rr, req)
}
