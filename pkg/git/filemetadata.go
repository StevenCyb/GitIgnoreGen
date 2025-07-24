package git

import "strings"

// FileMetadata represents metadata for a file in a Git repository.
type FileMetadata struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	DownloadURL string `json:"download_url"`
}

// FileMetadataList is a slice of FileMetadata.
type FileMetadataList []FileMetadata

// FilterExtension filters the FileMetadataList by a given file extension.
func (f *FileMetadataList) FilterExtension(extension string) FileMetadataList {
	var filtered FileMetadataList
	for _, file := range *f {
		if strings.HasSuffix(file.Name, extension) {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

// GetByName retrieves a FileMetadata by its name from the FileMetadataList.
func (f *FileMetadataList) GetByName(name string) *FileMetadata {
	for _, file := range *f {
		if file.Name == name {
			return &file
		}
	}
	return nil
}
