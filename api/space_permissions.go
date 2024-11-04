package api

import (
	"encoding/json"
	"fmt"
)

type SpacePermissionsResponse struct {
	Results []SpacePermission `json:"results"`
	Links   struct {
		Next string `json:"next"`
		Base string `json:"base"`
	} `json:"_links"`
}

type SpacePermission struct {
	Id        string `json:"id"`
	Principal struct {
		Type string `json:"type"`
		Id   string `json:"id"`
	} `json:"principal"`
	Operation struct {
		Key        string `json:"key"`
		TargetType string `json:"targetType"`
	}
}

func (apiClient *APIClient) GetSpacePermissionAssignments(spaceID string) ([]SpacePermission, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/spaces/%s/permissions", apiClient.Url, spaceID)
	headers := map[string]string{
		"Accept": "application/json",
	}
	resp, err := apiClient.Request("GET", url, headers, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var spacePermissionsResponse SpacePermissionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&spacePermissionsResponse); err != nil {
		return nil, err
	}

	return spacePermissionsResponse.Results, err
}
