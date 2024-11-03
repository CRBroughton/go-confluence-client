package api_test

import (
	"strings"
	"testing"
	"time"

	"github.com/crbroughton/go-confluence-client/api"
	"github.com/stretchr/testify/assert"
)

func TestGetPages(t *testing.T) {
	baseURL, email, apiToken, _, _ := api.GetENVValues(t)
	client := api.NewClient(baseURL, email, apiToken)

	pages, err := client.GetPages()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t, "test page", pages[5].Title)
	assert.Equal(t, "<p>test content</p>", pages[5].Body.Storage.Value)
}

func TestGetPageByID(t *testing.T) {
	baseURL, email, apiToken, pageID, _ := api.GetENVValues(t)
	client := api.NewClient(baseURL, email, apiToken)

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
	baseURL, email, apiToken, pageID, _ := api.GetENVValues(t)
	client := api.NewClient(baseURL, email, apiToken)

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
	baseURL, email, apiToken, pageID, _ := api.GetENVValues(t)
	client := api.NewClient(baseURL, email, apiToken)

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

func TestCreatesANewPageAndDeletesThePage(t *testing.T) {
	baseURL, email, apiToken, _, spaceID := api.GetENVValues(t)
	client := api.NewClient(baseURL, email, apiToken)
	var createdPageID string

	t.Run("Creating the page", func(t *testing.T) {
		page, err := client.CreatePage("new page title", spaceID, "body here")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		createdPageID = page.ID
		assert.Equal(t, "new page title", page.Title)
		assert.Equal(t, "body here", page.Body.Storage.Value)
	})
	// Wait to ensure the page creation is fully processed
	time.Sleep(5 * time.Second)
	t.Run("Deletes the created page", func(t *testing.T) {
		responseCode, err := client.DeletePage(createdPageID)

		if err != nil {
			if !strings.Contains(err.Error(), "0") {
				t.Fatalf("got the wrong error code, got %v", err)
			}
		}

		assert.Equal(t, 0, responseCode)
	})
}
