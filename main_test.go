package goconfluenceclient_test

import (
	"testing"

	goconfluenceclient "github.com/crbroughton/go-confluence-client"
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
