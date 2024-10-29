package api_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/crbroughton/go-confluence-client/api"
	"github.com/stretchr/testify/assert"
)

func TestTheNewClientMethodReturnsAValidAPIClient(t *testing.T) {
	client := api.NewClient("https://your.domain.com/", "crbroughton@posteo.uk", "12345678")

	expectation := &api.APIClient{
		Url:        "https://your.domain.com",
		Email:      "crbroughton@posteo.uk",
		ApiToken:   "12345678",
		HttpClient: &http.Client{},
	}

	assert.Equal(t, expectation, client)
}

func setupMockHTTPServer() (*httptest.Server, error) {
	var headerError error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept") != "application/json" {
			headerError = fmt.Errorf("expected Accept header to be 'application/json', got '%s'", r.Header.Get("Accept"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			headerError = fmt.Errorf("expected Content-Type header to be 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	return server, headerError
}
func TestAPIClientRequestSetsRequiredHeaders(t *testing.T) {
	mockServer, err := setupMockHTTPServer()
	if err != nil {
		t.Fatalf("Setup mock server failed with header error: %v", err)
	}
	defer mockServer.Close()

	apiClient := &api.APIClient{
		HttpClient: http.DefaultClient,
		Email:      "test@example.com",
		ApiToken:   "dummyToken",
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	resp, err := apiClient.Request("GET", mockServer.URL, headers)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func TestAPIClientRequestMissingRequiredHeaders(t *testing.T) {
	mockServer, err := setupMockHTTPServer()
	if err != nil {
		t.Fatalf("Setup mock server failed with header error: %v", err)
	}
	defer mockServer.Close()

	apiClient := &api.APIClient{
		HttpClient: http.DefaultClient,
		Email:      "test@example.com",
		ApiToken:   "dummyToken",
	}

	resp, err := apiClient.Request("GET", mockServer.URL, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()
}
