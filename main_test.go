package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup() *http.ServeMux {
	store = NewMemoryStore()
	mux := http.NewServeMux()
	registerRoutes(mux)
	return mux
}

func TestGetTasksEmpty(t *testing.T) {
	mux := setup()
	req := httptest.NewRequest(http.MethodGet, "/api/tasks", nil)
	rw := httptest.NewRecorder()
	mux.ServeHTTP(rw, req)
	if rw.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", rw.Code)
	}
	var got []Task
	if err := json.NewDecoder(rw.Body).Decode(&got); err != nil {
		t.Fatalf("decoding: %v", err)
	}
	if len(got) != 0 {
		t.Fatalf("expected 0 tasks, got %d", len(got))
	}
}

func TestPostCreateTask(t *testing.T) {
	mux := setup()
	b := bytes.NewBufferString(`{"title":"hello"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/tasks", b)
	req.Header.Set("Content-Type", "application/json")
	rw := httptest.NewRecorder()
	mux.ServeHTTP(rw, req)
	if rw.Code != http.StatusCreated {
		t.Fatalf("expected 201 got %d", rw.Code)
	}
	var got Task
	if err := json.NewDecoder(rw.Body).Decode(&got); err != nil {
		t.Fatalf("decoding: %v", err)
	}
	if got.Title != "hello" {
		t.Fatalf("expected title hello got %s", got.Title)
	}
}
