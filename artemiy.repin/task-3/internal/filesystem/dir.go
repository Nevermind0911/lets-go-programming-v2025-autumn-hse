package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

func EnsurePathExists(path string) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("не удалось создать директорию %s: %w", dir, err)
	}

	return nil
}
