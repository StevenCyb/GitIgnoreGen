package git

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRoundTripper struct {
	resp *http.Response
	err  error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.resp, m.err
}

func newMockClient(resp *http.Response, err error) *http.Client {
	return &http.Client{Transport: &mockRoundTripper{resp: resp, err: err}}
}

func TestNewWithClient_InvalidURL(t *testing.T) {
	t.Parallel()

	c, err := NewWithClient(http.DefaultClient, "https://github.com/toptal/gitignore.io/master/Localizations")
	assert.Nil(t, c)
	assert.Error(t, err)
}

func TestListFiles_Success(t *testing.T) {
	t.Parallel()

	jsonResp := `[{"name":"foo.json","type":"file","download_url":"url1"},{"name":"bar.txt","type":"file","download_url":"url2"}]`
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonResp)),
	}
	client, _ := NewWithClient(newMockClient(resp, nil), "https://github.com/toptal/gitignore.io/tree/master/Localizations")
	files, err := client.ListFiles(context.Background())
	assert.NoError(t, err)
	assert.Len(t, files, 2)
	assert.Equal(t, "foo.json", files[0].Name)
	assert.Equal(t, "bar.txt", files[1].Name)
}

func TestListFiles_HTTPError(t *testing.T) {
	t.Parallel()

	resp := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(bytes.NewBufferString("")),
	}
	client, _ := NewWithClient(newMockClient(resp, nil), "https://github.com/toptal/gitignore.io/tree/master/Localizations")
	files, err := client.ListFiles(context.Background())
	assert.Error(t, err)
	assert.Nil(t, files)
}

func TestListFiles_RequestError(t *testing.T) {
	t.Parallel()

	client, _ := NewWithClient(newMockClient(nil, errors.New("fail")), "https://github.com/toptal/gitignore.io/tree/master/Localizations")
	files, err := client.ListFiles(context.Background())
	assert.Error(t, err)
	assert.Nil(t, files)
}

func TestDownloadJSON_Success(t *testing.T) {
	t.Parallel()

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("{\"key\":\"value\"}")),
	}
	client, _ := NewWithClient(newMockClient(resp, nil), "https://github.com/toptal/gitignore.io/tree/master/Localizations")
	meta := FileMetadata{Name: "foo.json", Type: "file", DownloadURL: "url1"}
	content, err := client.Download(context.Background(), meta)
	assert.NoError(t, err)
	assert.Equal(t, "{\"key\":\"value\"}", *content)
}

func TestDownloadJSON_HTTPError(t *testing.T) {
	t.Parallel()

	resp := &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(bytes.NewBufferString("")),
	}
	client, _ := NewWithClient(newMockClient(resp, nil), "https://github.com/toptal/gitignore.io/tree/master/Localizations")
	meta := FileMetadata{Name: "foo.json", Type: "file", DownloadURL: "url1"}
	content, err := client.Download(context.Background(), meta)
	assert.Error(t, err)
	assert.Nil(t, content)
}

func TestDownloadJSON_RequestError(t *testing.T) {
	t.Parallel()

	client, _ := NewWithClient(newMockClient(nil, errors.New("fail")), "https://github.com/toptal/gitignore.io/tree/master/Localizations")
	meta := FileMetadata{Name: "foo.json", Type: "file", DownloadURL: "url1"}
	content, err := client.Download(context.Background(), meta)
	assert.Error(t, err)
	assert.Nil(t, content)
}
