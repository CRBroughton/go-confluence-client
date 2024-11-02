package api_test

import (
	"testing"

	"github.com/crbroughton/go-confluence-client/api"
	"github.com/stretchr/testify/assert"
)

func TestGetSpaces(t *testing.T) {
	baseURL, email, apiToken, _, _ := api.GetENVValues(t)
	client := api.NewClient(baseURL, email, apiToken)

	spaces, err := client.GetSpaces()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t, "163844", spaces[2].ID)
	assert.Equal(t, "My first space", spaces[2].Name)

	assert.Equal(t, "65853", spaces[1].ID)
	assert.Equal(t, "Test Space", spaces[1].Name)

}

func TestGetSpaceByKey(t *testing.T) {
	baseURL, email, apiToken, _, _ := api.GetENVValues(t)
	client := api.NewClient(baseURL, email, apiToken)

	space, err := client.FindSpaceByKey("TS")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t, "65853", space.ID)
	assert.Equal(t, "Test Space", space.Name)
}

func TestGetSpaceByID(t *testing.T) {
	baseURL, email, apiToken, _, spaceID := api.GetENVValues(t)
	client := api.NewClient(baseURL, email, apiToken)

	space, err := client.GetSpaceByID(spaceID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t, "65853", space.ID)
}
