package api

import (
	"encoding/json"
	"fmt"
)

type SpacePermissionsAssignmentsResponse struct {
	Results []SpacePermissionAssignment `json:"results"`
	Links   struct {
		Next string `json:"next"`
		Base string `json:"base"`
	} `json:"_links"`
}

type SpacePermissionAssignment struct {
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

func (apiClient *APIClient) GetSpacePermissionAssignments(spaceID string) ([]SpacePermissionAssignment, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/spaces/%s/permissions", apiClient.Url, spaceID)
	headers := map[string]string{
		"Accept": "application/json",
	}
	resp, err := apiClient.Request("GET", url, headers, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var spacePermissionAssignmentResponse SpacePermissionsAssignmentsResponse
	if err := json.NewDecoder(resp.Body).Decode(&spacePermissionAssignmentResponse); err != nil {
		return nil, err
	}

	return spacePermissionAssignmentResponse.Results, nil
}

type SpacePermissionsResponse struct {
	Results []SpacePermission `json:"results"`
	Links   struct {
		Next string `json:"next"`
		Base string `json:"base"`
	} `json:"_links"`
}

type SpacePermission struct {
	Id                    string   `json:"id"`
	DisplayName           string   `json:"displayName"`
	Description           string   `json:"description"`
	RequiredPermissionIds []string `json:"requiredPermissionIds"`
}

func (apiClient *APIClient) GetAvailableSpacePermissions() ([]SpacePermission, error) {
	url := fmt.Sprintf("%s/wiki/api/v2/space-permissions", apiClient.Url)
	headers := map[string]string{
		"Accept": "application/json",
	}
	resp, err := apiClient.Request("GET", url, headers, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var availableSpacePermissionsResponse SpacePermissionsResponse

	if err := json.NewDecoder(resp.Body).Decode(&availableSpacePermissionsResponse); err != nil {
		return nil, err
	}

	return availableSpacePermissionsResponse.Results, nil
}
