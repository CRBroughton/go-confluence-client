package api_test

import (
	"testing"

	"github.com/crbroughton/go-confluence-client/api"
	"github.com/stretchr/testify/assert"
)

func TestGetSpaceProperties(t *testing.T) {
	baseURL, email, apiToken, _, spaceID := api.GetENVValues(t)
	client := api.NewClient(baseURL, email, apiToken)

	spaceProperties, err := client.GetSpaceProperties(spaceID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectation := []api.SpaceProperties{}
	assert.Equal(t, expectation, spaceProperties)
}
