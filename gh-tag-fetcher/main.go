package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Release struct {
	TagName string `json:"tag_name"`
}

func getLatestReleaseTag(owner, repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to retrieve latest release tag. Status code: %d", resp.StatusCode)
	}

	var releaseData Release
	err = json.NewDecoder(resp.Body).Decode(&releaseData)
	if err != nil {
		return "", err
	}

	return releaseData.TagName, nil
}

func main() {
	// Replace with your GitHub repository owner and repository name
	owner := "fluxcd"
	repositories := []string{"flux2", "flagger"}

	for _, repo := range repositories {
		latestReleaseTag, err := getLatestReleaseTag(owner, repo)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("The latest release tag of %s/%s is: %s\n", owner, repo, latestReleaseTag)
	}
}
