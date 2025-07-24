package handler

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/StevenCyb/GoCLI/pkg/cli"
)

const sourcePrefix = "# Generated "

func UpdateHandler(templateURL string, timeout time.Duration) func(*cli.Context) error {
	return func(_ *cli.Context) error {
		currentGitignore, err := os.ReadFile(".gitignore")
		if err != nil {
			return fmt.Errorf("failed to get current .gitignore: %w", err)
		}

		extractedNames := []string{"", ""}
		lines := string(currentGitignore)
		for _, line := range strings.Split(lines, "\n") {
			if strings.HasPrefix(line, sourcePrefix) {
				name := strings.TrimSpace(line[len(sourcePrefix):])
				if name == "" || !strings.HasSuffix(name, ".gitignore") {
					continue
				}
				extractedNames = append(extractedNames, strings.TrimSuffix(name, ".gitignore"))
			}
		}

		return BuildHandler(templateURL, timeout, extractedNames)(nil)
	}
}
