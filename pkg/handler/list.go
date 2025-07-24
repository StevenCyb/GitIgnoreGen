package handler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/StevenCyb/GitIgnoreGen/pkg/git"
	"github.com/StevenCyb/GoCLI/pkg/cli"
)

func ListHandler(client git.IClient, timeout time.Duration) func(*cli.Context) error {
	return func(_ *cli.Context) error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		files, err := client.ListFiles(ctx)
		if err != nil {
			return fmt.Errorf("failed to list files: %w", err)
		}

		for _, file := range files {
			fmt.Println("- " + strings.TrimSuffix(file.Name, ".gitignore"))
		}

		return nil
	}
}
