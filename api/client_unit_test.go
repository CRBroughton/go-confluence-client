package api_test

import (
	"testing"

	"github.com/crbroughton/go-confluence-client/api"
	"github.com/stretchr/testify/assert"
)

func TestTheNewClientMethodReturnsAValidAPIClient(t *testing.T) {
	client := api.NewClient("https://your.domain.com/", "crbroughton@posteo.uk", "12345678")

	expectation := &api.APIClient{
		Url:      "https://your.domain.com",
		Email:    "crbroughton@posteo.uk",
		ApiToken: "12345678",
	}

	assert.Equal(t, expectation, client)
}
