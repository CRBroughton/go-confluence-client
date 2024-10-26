package api

import (
	"strings"
)

type APIClient struct {
	Url      string
	Email    string
	ApiToken string
}

func NewClient(url, email, apiToken string) *APIClient {
	return &APIClient{
		Url:      strings.TrimSuffix(url, "/"),
		Email:    email,
		ApiToken: apiToken,
	}
}
