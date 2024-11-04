package api_test

import (
	"testing"

	"github.com/crbroughton/go-confluence-client/api"
	"github.com/stretchr/testify/assert"
)

func TestGetSpacePermissions(t *testing.T) {
	baseURL, email, apiToken, _, spaceID := api.GetENVValues(t)
	client := api.NewClient(baseURL, email, apiToken)

	spacePermissions, err := client.GetSpacePermissionAssignments(spaceID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	assert.Equal(t, 25, len(spacePermissions))
}
