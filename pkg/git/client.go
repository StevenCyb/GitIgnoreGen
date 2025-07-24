package git

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// expectedURLParts is the number of parts we expect in the GitHub subdir URL.
const expectedURLParts = 5

// Client represents a GitHub client for accessing repository contents.
type Client struct {
	Owner      string
	Repo       string
	Branch     string
	SubdirPath string
	httpClient *http.Client
}

// New creates a new GitHub client with the default HTTP client.
func New(baseURL string) (*Client, error) {
	return NewWithClient(http.DefaultClient, baseURL)
}

// NewWithClient creates a new GitHub client with a custom HTTP client.
func NewWithClient(httpClient *http.Client, baseURL string) (*Client, error) {
	re := regexp.MustCompile(`https://github.com/([^/]+)/([^/]+)/tree/([^/]+)/(.*)`)
	matches := re.FindStringSubmatch(baseURL)
	if len(matches) != expectedURLParts {
		return nil, fmt.Errorf("invalid GitHub subdir URL: %s", baseURL)
	}

	return &Client{
		Owner:      matches[1],
		Repo:       matches[2],
		Branch:     matches[3],
		SubdirPath: matches[4],
		httpClient: httpClient,
	}, nil
}

// ListFiles lists files in the specified GitHub repository subdirectory.
func (c *Client) ListFiles(ctx context.Context) ([]FileMetadata, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s?ref=%s",
		c.Owner, c.Repo, c.SubdirPath, c.Branch)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	var filesMetadata []FileMetadata

	if err = json.NewDecoder(resp.Body).Decode(&filesMetadata); err != nil {
		return nil, err
	}

	return filesMetadata, nil
}

// Download downloads the content of a file from the GitHub repository.
func (c *Client) Download(ctx context.Context, fileMetadata FileMetadata) (*string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fileMetadata.DownloadURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub raw file error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	content := string(body)

	return &content, nil
}
