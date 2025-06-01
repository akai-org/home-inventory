package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockDB implements db.DB interface for testing
type mockDB struct {
	pingErr error
}

func (m *mockDB) Ping(ctx context.Context) error {
	return m.pingErr
}

func TestHealthCheck_Healthy(t *testing.T) {
	app := &App{
		Database: &mockDB{pingErr: nil},
	}

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	app.HealthCheck(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if body["status"] != "ok" {
		t.Errorf("expected status 'ok', got '%s'", body["status"])
	}
}

func TestHealthCheck_Unhealthy(t *testing.T) {
	app := &App{
		Database: &mockDB{pingErr: errors.New("db down")},
	}

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	app.HealthCheck(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("expected status 503, got %d", resp.StatusCode)
	}

	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if body["status"] != "unhealthy" {
		t.Errorf("expected status 'unhealthy', got '%s'", body["status"])
	}
}
