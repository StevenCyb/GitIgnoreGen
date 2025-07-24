package handler

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/StevenCyb/GitIgnoreGen/pkg/git"
	"github.com/StevenCyb/GoCLI/pkg/cli"
)

func BuildHandler(templateURL string, timeout time.Duration, args []string) func(*cli.Context) error {
	return func(_ *cli.Context) error {
		gitignoreCases := args[2:]

		client, err := git.New(templateURL)
		if err != nil {
			return fmt.Errorf("failed to create Git client: %w", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		files, err := client.ListFiles(ctx)
		if err != nil {
			return fmt.Errorf("failed to list files: %w", err)
		}

		gitignoreFile := ""
		for _, gitignoreCase := range gitignoreCases {
			found := false
			gitignoreCase = strings.ToLower(gitignoreCase)
			for _, file := range files {
				if strings.ToLower(file.Name) == gitignoreCase+".gitignore" {
					found = true
					gitignoreFile += fmt.Sprintf("# Generated %s\n", file.Name)
					content, err := client.Download(ctx, file)
					if err != nil {
						return fmt.Errorf("failed to download file %s: %w", file.Name, err)
					}
					gitignoreFile += *content + "\n"
					break
				}
			}

			if !found {
				return fmt.Errorf("no .gitignore file found for case: %s", gitignoreCase)
			}
		}

		gitignoreFile = strings.TrimSuffix(gitignoreFile, "\n")
		err = os.WriteFile(".gitignore", []byte(gitignoreFile), 0644)
		if err != nil {
			return fmt.Errorf("failed to write .gitignore file: %w", err)
		}

		return nil
	}
}
