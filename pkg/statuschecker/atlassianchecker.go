package statuschecker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AtlassianStatus struct {
	Page struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		URL       string `json:"url"`
		TimeZone  string `json:"time_zone"`
		UpdatedAt string `json:"updated_at"`
	} `json:"page"`
	Components []struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		Status         string `json:"status"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
		Position       int    `json:"position"`
		Description    string `json:"description"`
		Showcase       bool   `json:"showcase"`
		StartDate      string `json:"start_date"`
		GroupID        string `json:"group_id"`
		PageID         string `json:"page_id"`
		Group          bool   `json:"group"`
		OnlyShowIfDegr bool   `json:"only_show_if_degraded"`
	} `json:"components"`
}

func (a *AtlassianChecker) CheckStatus() (string, error) {
	url := "https://pv54g7ltsc24.statuspage.io/api/v2/summary.json"

	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch Atlassian status: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch Atlassian status: %s", resp.Status)
	}

	// Parse JSON response
	var status AtlassianStatus
	err = json.NewDecoder(resp.Body).Decode(&status)
	if err != nil {
		return "", fmt.Errorf("failed to decode Atlassian status response: %v", err)
	}

	// Extract the relevant status information
	statusInfo := fmt.Sprintf("Atlassian Status:\n- Name: %s\n- URL: %s\n- Time Zone: %s\n- Updated At: %s\n",
		status.Page.Name, status.Page.URL, status.Page.TimeZone, status.Page.UpdatedAt)

	return statusInfo, nil
}

func (a *AtlassianChecker) CanHandle(url string) bool {
	// Check if the URL is the Atlassian status page URL
	return url == "https://status.atlassian.com"
}
