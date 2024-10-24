package goconfluenceclient_test

import (
	"os"
	"testing"

	goconfluenceclient "github.com/crbroughton/go-confluence-client"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

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
	err := godotenv.Load(".env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}
	baseURL := os.Getenv("CONFLUENCE_BASE_URL")
	email := os.Getenv("CONFLUENCE_EMAIL")
	apiToken := os.Getenv("CONFLUENCE_API_TOKEN")
	pageID := os.Getenv("CONFLUENCE_PAGE_ID")

	if baseURL == "" || email == "" || apiToken == "" || pageID == "" {
		t.Skip("Skipping integration test; Please provide the required environment variables")
	}

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
