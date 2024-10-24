package goconfluenceclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type Page struct {
	ID      string      `json:"id"`
	Status  string      `json:"status"`
	Title   string      `json:"title"`
	Body    PageBody    `json:"body"`
	Version PageVersion `json:"version"`
}

type PageBody struct {
	Storage PageStorage `json:"storage"`
}

type PageStorage struct {
	Representation string `json:"representation"`
	Value          string `json:"value"`
}

type PageVersion struct {
	Number int `json:"number"`
}

func (apiClient *APIClient) GetPageByID(pageID string) (*Page, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/pages/%s?body-format=storage", apiClient.Url, pageID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(apiClient.Email, apiClient.ApiToken)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get the confluence page: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var page Page
	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, err
	}

	return &page, nil
}
