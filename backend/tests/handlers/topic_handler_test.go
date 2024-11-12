package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock implementation of TopicService for testing
type MockTopicService struct {
	FetchTopicsFunc func(ctx context.Context) ([]string, error)
}

func (m *MockTopicService) FetchTopics(ctx context.Context) ([]string, error) {
	return m.FetchTopicsFunc(ctx)
}

// TopicHandler is the handler we are testing
type TopicHandler struct {
	Service *MockTopicService
}

// GetTopics uses the Service to fetch topics and responds with JSON
func (h *TopicHandler) GetTopics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	topics, err := h.Service.FetchTopics(ctx)
	if err != nil {
		http.Error(w, `{"error":"Failed to fetch topics"}`, http.StatusInternalServerError)
		return
	}
	response := map[string]interface{}{"topics": topics}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response) // Use topics variable in response
}

func TestGetTopics_Success(t *testing.T) {
	// Mock service returns a list of topics
	mockService := &MockTopicService{
		FetchTopicsFunc: func(ctx context.Context) ([]string, error) {
			return []string{"Climate Change", "Renewable Energy"}, nil
		},
	}
	handler := &TopicHandler{Service: mockService}

	// Create a new HTTP request and response recorder
	req, _ := http.NewRequest("GET", "/topics", nil)
	rr := httptest.NewRecorder()

	// Call the handler's GetTopics method
	handler.GetTopics(rr, req)

	// Check if the response status code is 200
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	// Check if the response body contains the expected topics JSON
	expected := `{"topics":["Climate Change","Renewable Energy"]}`
	if rr.Body.String() != expected+"\n" { // json.Encode adds a newline
		t.Errorf("expected body %s, got %s", expected, rr.Body.String())
	}
}

func TestGetTopics_Failure(t *testing.T) {
	// Mock service returns an error
	mockService := &MockTopicService{
		FetchTopicsFunc: func(ctx context.Context) ([]string, error) {
			return nil, errors.New("fetch error")
		},
	}
	handler := &TopicHandler{Service: mockService}

	// Create a new HTTP request and response recorder
	req, _ := http.NewRequest("GET", "/topics", nil)
	rr := httptest.NewRecorder()

	// Call the handler's GetTopics method
	handler.GetTopics(rr, req)

	// Check if the response status code is 500
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, status)
	}

	// Check if the response body contains the expected error message JSON
	expected := `{"error":"Failed to fetch topics"}`
	if rr.Body.String() != expected+"\n" {
		t.Errorf("expected body %s, got %s", expected, rr.Body.String())
	}
}
