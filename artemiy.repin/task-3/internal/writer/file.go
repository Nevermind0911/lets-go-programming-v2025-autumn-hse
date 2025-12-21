package writer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Nevermind0911/task-3/internal/filesystem"
)

func SaveJSON(path string, data interface{}) error {
	if err := filesystem.EnsurePathExists(path); err != nil {
		return fmt.Errorf("подготовка директории: %w", err)
	}

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка генерации JSON: %w", err)
	}

	if err := os.WriteFile(path, bytes, 0o600); err != nil {
		return fmt.Errorf("ошибка записи файла: %w", err)
	}

	return nil
}
