package languages

import (
	"fmt"
	"path/filepath"
)

func listProtoFileNamesInFolder(folder string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(folder, "*.proto"))
	if err != nil {
		return nil, fmt.Errorf("failed to glob proto files: %w", err)
	}
	for i, v := range files {
		files[i] = filepath.Base(v)
	}
	return files, nil
}
