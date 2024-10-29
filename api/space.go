package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SpaceResponse struct {
	Results []Space `json:"results"`
	Links   struct {
		Next string `json:"next"`
		Base string `json:"base"`
	} `json:"_links"`
}

type Space struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	AuthorID    string `json:"authorId"`
	CreatedAt   string `json:"createdAt"`
	HomepageID  string `json:"homepageId"`
	Description struct {
		Plain map[string]interface{} `json:"plain"`
		View  map[string]interface{} `json:"view"`
	} `json:"description"`
	Icon struct {
		Path            string `json:"path"`
		APIDownloadLink string `json:"apiDownloadLink"`
	} `json:"icon"`
	Links struct {
		WebUI string `json:"webui"`
	} `json:"_links"`
}

func (apiClient *APIClient) GetSpaces() ([]Space, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/spaces", apiClient.Url)
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	resp, err := apiClient.Request("GET", url, headers, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var spaceResponse SpaceResponse
	if err := json.NewDecoder(resp.Body).Decode(&spaceResponse); err != nil {
		return nil, err
	}

	return spaceResponse.Results, nil
}

func (apiClient *APIClient) FindSpaceByKey(spaceKey string) (*Space, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/spaces", apiClient.Url)
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	resp, err := apiClient.Request("GET", url, headers, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var spaceResponse SpaceResponse
	if err := json.NewDecoder(resp.Body).Decode(&spaceResponse); err != nil {
		return nil, err
	}

	for _, space := range spaceResponse.Results {
		if space.Key == spaceKey {
			return &space, nil
		}
	}

	return nil, errors.New("space not found")
}
