package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type APIClient struct {
	Url        string
	Email      string
	ApiToken   string
	HttpClient *http.Client
}

func NewClient(url, email, apiToken string) *APIClient {
	return &APIClient{
		Url:        strings.TrimSuffix(url, "/"),
		Email:      email,
		ApiToken:   apiToken,
		HttpClient: &http.Client{},
	}
}

func (apiClient *APIClient) Request(
	requestType, url string,
	headers map[string]string,
	body io.Reader,
) (*http.Response, error) {
	req, err := http.NewRequest(requestType, url, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(apiClient.Email, apiClient.ApiToken)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := apiClient.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return resp, nil
}
