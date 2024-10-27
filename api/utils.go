package api

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func GetENVValues(t *testing.T) (string, string, string, string) {
	err := godotenv.Load("../.env")
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
