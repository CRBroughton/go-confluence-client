package api_test

import (
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

func setupMockHTTPServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate headers dynamically in each request
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept header to be 'application/json', got '%s'", r.Header.Get("Accept"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type header to be 'application/json', got '%s'", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
	}))
}
func TestAPIClientRequestSetsRequiredHeaders(t *testing.T) {
	mockServer := setupMockHTTPServer(t)
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

	resp, err := apiClient.Request("GET", mockServer.URL, headers, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func TestAPIClientRequestMissingRequiredHeaders(t *testing.T) {
	mockServer := setupMockHTTPServer(t)
	defer mockServer.Close()

	apiClient := &api.APIClient{
		HttpClient: http.DefaultClient,
		Email:      "test@example.com",
		ApiToken:   "dummyToken",
	}

	resp, err := apiClient.Request("GET", mockServer.URL, nil, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()
}
