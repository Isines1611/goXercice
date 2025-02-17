// Lite version of Gitter CLI

package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type GitHubFile struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	SHA         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
}

const apiURL = ""

func DownloadFiles(dir string) {
	parsedUrl, _ := url.Parse(apiURL)
	parts := strings.Split(strings.Trim(parsedUrl.Path, "/"), "/")

	owner := parts[0]
	repo := parts[1]
	path := ""

	if len(parts) > 2 {
		path = strings.Join(parts[4:], "/")
	} else {
		path = ""
	}

	var dirPaths []string
	var files []GitHubFile

	dirPaths = append(dirPaths, path)

	mapFiles(owner, repo, dir, &dirPaths, &files)
}

func mapFiles(owner, repo, baseDir string, dirPaths *[]string, ptr *[]GitHubFile) {
	dirPath := (*dirPaths)[len(*dirPaths)-1]
	*dirPaths = (*dirPaths)[:len(*dirPaths)-1]

	files, err := getContent(owner, repo, dirPath)
	if err != nil {
		fmt.Println("❌ Error getting the content:", err)
	}

	for _, file := range files {

		if file.Type == "file" {
			*ptr = append(*ptr, file)
		} else {
			*dirPaths = append(*dirPaths, file.Path)
		}
	}

	if len(*dirPaths) <= 0 {
		downloadAll(baseDir, *ptr)
	} else {
		mapFiles(owner, repo, baseDir, dirPaths, ptr)
	}
}

func getContent(owner, repo, path string) ([]GitHubFile, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", owner, repo, path)

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	// Parse JSON response
	var files []GitHubFile // Use slice to match array response
	if err := json.Unmarshal(body, &files); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	return files, nil
}

func downloadAll(baseDir string, files []GitHubFile) {
	if baseDir == "" {
		baseDir, _ = os.Getwd()
	}

	for _, file := range files {
		dir := filepath.Join(baseDir, filepath.Dir(file.Path))
		err := downloadFile(file.DownloadURL, dir, file.Name)
		if err != nil {
			return
		}
	}
}

func downloadFile(url, dir, filename string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filePath := filepath.Join(dir, filename)

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch file: %w", err)
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	// Create file
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// copy response body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}
