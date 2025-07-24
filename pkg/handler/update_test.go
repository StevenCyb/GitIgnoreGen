package handler

import (
	"errors"
	"os"
	"testing"

	"github.com/StevenCyb/GitIgnoreGen/pkg/git"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateHandler_Success(t *testing.T) {
	t.Parallel()

	// Prepare .gitignore file
	content := "# Generated foo.gitignore\n# Generated bar.gitignore"
	os.WriteFile(".gitignore", []byte(content), 0644)
	defer os.Remove(".gitignore")

	client := &git.ClientMock{}
	files := []git.FileMetadata{
		{Name: "foo.gitignore", Type: "file", DownloadURL: "url1"},
		{Name: "bar.gitignore", Type: "file", DownloadURL: "url2"},
	}
	client.On("ListFiles", mock.Anything).Return(files, nil)
	client.On("Download", mock.Anything, files[0]).Return(ptr("foo content"), nil)
	client.On("Download", mock.Anything, files[1]).Return(ptr("bar content"), nil)

	err := UpdateHandler(client, 5)(nil)
	defer os.Remove(".gitignore")
	assert.NoError(t, err)
}

func TestUpdateHandler_FileReadError(t *testing.T) {
	t.Parallel()

	os.Remove(".gitignore") // Ensure file does not exist
	client := &git.ClientMock{}
	err := UpdateHandler(client, 5)(nil)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "failed to get current .gitignore")
}

func TestUpdateHandler_BuildHandlerError(t *testing.T) {
	t.Parallel()

	content := "# Generated foo.gitignore"
	os.WriteFile(".gitignore", []byte(content), 0644)
	defer os.Remove(".gitignore")

	client := &git.ClientMock{}
	client.On("ListFiles", mock.Anything).Return(nil, errors.New("API error"))
	err := UpdateHandler(client, 5)(nil)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "API error")
}
