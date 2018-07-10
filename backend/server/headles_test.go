package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ComputePractice2018/photohosting/backend/data"
)

var testProfile = "{\"name\":\"Sergey\",\"nickname\":\"Serg0vich\",\"Email\":\"ssergeev@mail.com\",\"Description\":\"О себе\",\"Birthday\":\"28.11.1981\",\"Photos\":\"Фотографии\"}"

func TestCrudHandlers(t *testing.T) {
	cl := data.NewProfileList()
	router := NewRouter(cl)

	req, err := http.NewRequest("GET", "/api/photohosting/profiles", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Error("Expected 200 code")
	}
	if resp.Header.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Error("Expected json MIME-type")
	}

	testData := strings.NewReader(testProfile)
	req, err = http.NewRequest("POST", "/api/photohosting/profiles", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 code (gotten: %d)", resp.StatusCode)
	}
	if resp.Header.Get("Location") != "/api/photohosting/profiles/1" {
		t.Error("Expected another location")
	}

	if len(cl.GetProfile()) != 1 {
		t.Error("Expected new value")
	}

	testData = strings.NewReader(testProfile)
	req, err = http.NewRequest("PUT", "/api/photohosting/profiles/1", testData)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected 201 code (gotten: %d)", resp.StatusCode)
	}
	if resp.Header.Get("Location") != "/api/photohosting/profiles/1" {
		t.Error("Expected another location")
	}

	if len(cl.GetProfiles()) != 1 {
		t.Error("Expected old value")
	}

	req, err = http.NewRequest("DELETE", "/api/photohosting/profiles/1", nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	resp = w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204 code (gotten: %d)", resp.StatusCode)
	}
	if len(cl.Getprofiles()) != 0 {
		t.Error("Expected old value")
	}
}
