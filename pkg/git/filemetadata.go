package git

// FileMetadata represents metadata for a file in a Git repository.
type FileMetadata struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	DownloadURL string `json:"download_url"`
}
