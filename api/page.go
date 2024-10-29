package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type GetPagesResponse struct {
	Results []Page `json:"results"`
}

func (apiClient *APIClient) GetPages() ([]Page, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/pages?body-format=storage", apiClient.Url)

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	resp, err := apiClient.Request("GET", url, headers, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pages GetPagesResponse
	err = json.Unmarshal(body, &pages)
	if err != nil {
		return nil, err
	}

	return pages.Results, nil
}

// Retrieves a Confluence page via it's ID
func (apiClient *APIClient) GetPageByID(pageID string) (*Page, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/pages/%s?body-format=storage", apiClient.Url, pageID)

	headers := map[string]string{
		"Accept": "application/json",
	}
	resp, err := apiClient.Request("GET", url, headers, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	resp, err := apiClient.Request("PUT", url, headers, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
