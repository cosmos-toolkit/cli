package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	templatesRepo = "https://api.github.com/repos/cosmos-toolkit/templates/contents"
	packagesRepo  = "https://api.github.com/repos/cosmos-toolkit/packages/contents"
)

type contentItem struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// ListTemplates returns template names from cosmos-toolkit/templates.
func ListTemplates() ([]string, error) {
	return listDirs(templatesRepo)
}

// ListPackages returns package names from cosmos-toolkit/packages.
func ListPackages() ([]string, error) {
	return listDirs(packagesRepo)
}

func listDirs(apiURL string) ([]string, error) {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return []string{}, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var items []contentItem
	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var dirs []string
	for _, item := range items {
		if item.Type == "dir" && item.Name != "." && item.Name != ".." {
			dirs = append(dirs, item.Name)
		}
	}
	return dirs, nil
}
