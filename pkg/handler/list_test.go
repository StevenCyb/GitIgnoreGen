package handler

import (
	"errors"
	"testing"

	"github.com/StevenCyb/GitIgnoreGen/pkg/git"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListHandler(t *testing.T) {
	t.Parallel()

	client := &git.ClientMock{}
	expectedFiles := []git.FileMetadata{
		{Name: "foo.gitignore", Type: "file", DownloadURL: "url1"},
		{Name: "bar.gitignore", Type: "file", DownloadURL: "url2"},
	}
	client.On("ListFiles", mock.Anything).Return(expectedFiles, nil)
	err := ListHandler(client, 5)(nil)
	assert.NoError(t, err)
}

func TestListHandler_APIError(t *testing.T) {
	t.Parallel()

	var expectedErr error = errors.New("API error")
	client := &git.ClientMock{}
	client.On("ListFiles", mock.Anything).Return(nil, expectedErr)
	err := ListHandler(client, 5)(nil)
	assert.Error(t, err)
	assert.ErrorContains(t, err, expectedErr.Error())
}
