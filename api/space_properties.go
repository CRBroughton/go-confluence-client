package api

import (
	"encoding/json"
	"fmt"
)

type GetSpacePropertiesResponse struct {
	Results []SpaceProperties `json:"results"`
	Links   struct {
		Next string `json:"next"`
		Base string `json:"base"`
	} `json:"_links"`
}

type SpaceProperties struct {
	Id        string                 `json:"id"`
	Key       string                 `json:"key"`
	CreatedAt string                 `json:"createdAt"`
	CreatedBy string                 `json:"createdBy"`
	Version   SpacePropertiesVersion `json:"version"`
}

type SpacePropertiesVersion struct {
	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	Message   string `json:"message"`
	Number    int    `json:"number"`
}

func (apiClient *APIClient) GetSpaceProperties(spaceID string) ([]SpaceProperties, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/spaces/%s/properties", apiClient.Url, spaceID)
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	resp, err := apiClient.Request("GET", url, headers, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var spacePropertiesResponse GetSpacePropertiesResponse
	if err := json.NewDecoder(resp.Body).Decode(&spacePropertiesResponse); err != nil {
		return nil, err
	}

	return spacePropertiesResponse.Results, nil
}
