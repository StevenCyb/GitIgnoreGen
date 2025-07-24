package handler

import (
	"errors"
	"os"
	"testing"

	"github.com/StevenCyb/GitIgnoreGen/pkg/git"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBuildHandler_Success(t *testing.T) {
	t.Parallel()

	client := &git.ClientMock{}
	files := []git.FileMetadata{
		{Name: "foo.gitignore", Type: "file", DownloadURL: "url1"},
		{Name: "bar.gitignore", Type: "file", DownloadURL: "url2"},
	}
	client.On("ListFiles", mock.Anything).Return(files, nil)
	client.On("Download", mock.Anything, files[0]).Return(ptr("foo content"), nil)
	client.On("Download", mock.Anything, files[1]).Return(ptr("bar content"), nil)

	args := []string{"build", "--out", "foo", "bar"}
	err := BuildHandler(client, 5, args)(nil)
	defer os.Remove(".gitignore")
	assert.NoError(t, err)
}

func TestBuildHandler_ListFilesError(t *testing.T) {
	t.Parallel()

	client := &git.ClientMock{}
	client.On("ListFiles", mock.Anything).Return(nil, errors.New("API error"))
	args := []string{"build", "--out", "foo"}
	err := BuildHandler(client, 5, args)(nil)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "API error")
}

func TestBuildHandler_DownloadError(t *testing.T) {
	t.Parallel()

	client := &git.ClientMock{}
	files := []git.FileMetadata{
		{Name: "foo.gitignore", Type: "file", DownloadURL: "url1"},
	}
	client.On("ListFiles", mock.Anything).Return(files, nil)
	client.On("Download", mock.Anything, files[0]).Return(nil, errors.New("download error"))
	args := []string{"build", "--out", "foo"}
	err := BuildHandler(client, 5, args)(nil)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "download error")
}

func TestBuildHandler_FileNotFound(t *testing.T) {
	t.Parallel()

	client := &git.ClientMock{}
	files := []git.FileMetadata{
		{Name: "foo.gitignore", Type: "file", DownloadURL: "url1"},
	}
	client.On("ListFiles", mock.Anything).Return(files, nil)
	args := []string{"build", "--out", "bar"}
	err := BuildHandler(client, 5, args)(nil)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "no .gitignore file found for case: bar")
}

func ptr(s string) *string {
	return &s
}
