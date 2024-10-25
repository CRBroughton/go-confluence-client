package goconfluenceclient_test

import (
	"os"
	"testing"

	goconfluenceclient "github.com/crbroughton/go-confluence-client"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func getENVValues(t *testing.T) (string, string, string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}
	baseURL := os.Getenv("CONFLUENCE_BASE_URL")
	email := os.Getenv("CONFLUENCE_EMAIL")
	apiToken := os.Getenv("CONFLUENCE_API_TOKEN")
	pageID := os.Getenv("CONFLUENCE_PAGE_ID")

	if baseURL == "" || email == "" || apiToken == "" || pageID == "" {
		t.Skip("Skipping integration tests; Please provide the required environment variables")
	}

	return baseURL, email, apiToken, pageID
}

func TestTheNewClientMethodReturnsAValidAPIClient(t *testing.T) {
	client := goconfluenceclient.NewClient("https://your.domain.com/", "crbroughton@posteo.uk", "12345678")

	expectation := &goconfluenceclient.APIClient{
		Url:      "https://your.domain.com",
		Email:    "crbroughton@posteo.uk",
		ApiToken: "12345678",
	}

	assert.Equal(t, expectation, client)
}

func TestGetPageByID(t *testing.T) {
	baseURL, email, apiToken, pageID := getENVValues(t)
	client := goconfluenceclient.NewClient(baseURL, email, apiToken)

	page, err := client.GetPageByID(pageID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if page.ID != pageID {
		t.Errorf("expected ID '%s', got %s", pageID, page.ID)
	}

	assert.Equal(t, "test page", page.Title)
	assert.Equal(t, "<p>test content</p>", page.Body.Storage.Value)

}

func TestUpdatePageByID(t *testing.T) {
	baseURL, email, apiToken, pageID := getENVValues(t)
	client := goconfluenceclient.NewClient(baseURL, email, apiToken)

	page, err := client.GetPageByID(pageID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	pageBody := "<h2>page-content</h2>"

	page, err = client.UpdatePageByID(page.ID, "new-title", pageBody, page.Version.Number+1, "version message")
	if err != nil {
		t.Error("failed to update the page:", err.Error())
	}

	assert.Equal(t, "new-title", page.Title)
	assert.Equal(t, "<h2>page-content</h2>", page.Body.Storage.Value)

	// Cleanup
	ResetState(t)
}

func ResetState(t *testing.T) {
	baseURL, email, apiToken, pageID := getENVValues(t)
	client := goconfluenceclient.NewClient(baseURL, email, apiToken)

	page, err := client.GetPageByID(pageID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	pageBody := "<p>test content</p>"

	page, err = client.UpdatePageByID(page.ID, "test page", pageBody, page.Version.Number+1, "version message")
	if err != nil {
		t.Error("failed to update the page:", err.Error(), page.ID)
	}
}
