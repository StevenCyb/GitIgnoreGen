package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterExtension(t *testing.T) {
	files := FileMetadataList{
		{Name: "foo.json", Type: "file", DownloadURL: "url1"},
		{Name: "bar.txt", Type: "file", DownloadURL: "url2"},
		{Name: "baz.json", Type: "file", DownloadURL: "url3"},
		{Name: "qux.md", Type: "file", DownloadURL: "url4"},
	}
	jsonFiles := files.FilterExtension(".json")
	assert.Len(t, jsonFiles, 2)
	assert.Equal(t, "foo.json", jsonFiles[0].Name)
	assert.Equal(t, "baz.json", jsonFiles[1].Name)
}

func TestGetByName(t *testing.T) {
	files := FileMetadataList{
		{Name: "foo.json", Type: "file", DownloadURL: "url1"},
		{Name: "bar.txt", Type: "file", DownloadURL: "url2"},
	}
	f := files.GetByName("foo.json")
	assert.NotNil(t, f)
	assert.Equal(t, "foo.json", f.Name)
	assert.Nil(t, files.GetByName("notfound.json"))
}

func TestGetByName_NotFound(t *testing.T) {
	files := FileMetadataList{
		{Name: "foo.json", Type: "file", DownloadURL: "url1"},
		{Name: "bar.txt", Type: "file", DownloadURL: "url2"},
	}
	f := files.GetByName("notfound.json")
	assert.Nil(t, f)
}
