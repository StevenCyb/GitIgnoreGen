package git

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
}

func (m *ClientMock) ListFiles(ctx context.Context) ([]FileMetadata, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]FileMetadata), args.Error(1)
}

func (m *ClientMock) Download(ctx context.Context, file FileMetadata) (*string, error) {
	args := m.Called(ctx, file)
	if content, ok := args.Get(0).(*string); ok {
		return content, args.Error(1)
	}
	return nil, args.Error(1)
}
