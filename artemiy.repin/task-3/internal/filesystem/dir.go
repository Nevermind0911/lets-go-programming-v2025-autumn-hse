package filesystem

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const defaultDirPermission fs.FileMode = 0o755

func EnsurePathExists(path string) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, defaultDirPermission); err != nil {
		return fmt.Errorf("не удалось создать директорию %s: %w", dir, err)
	}

	return nil
}
