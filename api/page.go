package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
	Number  int    `json:"number"`
	Message string `json:"message"`
}

// Retrieves a Confluence page via it's ID
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

func (apiClient *APIClient) UpdatePageByID(pageID, title, pageBody string, versionNumber int, versionMessage string) (*Page, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/pages/%s?body-format=storage", apiClient.Url, pageID)

	pageData := Page{
		ID:     pageID,
		Status: "current",
		Title:  title,
		Body: PageBody{
			Storage: PageStorage{
				Value:          pageBody,
				Representation: "storage",
			},
		},
		Version: PageVersion{
			Number:  versionNumber,
			Message: versionMessage,
		},
	}

	jsonData, err := json.Marshal(pageData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(apiClient.Email, apiClient.ApiToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update the confluence page: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var updatedPage Page
	err = json.Unmarshal(body, &updatedPage)
	if err != nil {
		return nil, err
	}

	return &updatedPage, nil
}
