package api_test

import (
	"testing"

	"github.com/crbroughton/go-confluence-client/api"
	"github.com/stretchr/testify/assert"
)

func TestGetSpaceByID(t *testing.T) {
	baseURL, email, apiToken, _ := api.GetENVValues(t)
	client := api.NewClient(baseURL, email, apiToken)

	space, err := client.FindSpaceByKey("TS")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t, "65853", space.ID)
	assert.Equal(t, "Test Space", space.Name)
}
