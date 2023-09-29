package filesystem

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func PathExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func ScanForFiles(path string, fileTypes []string) ([]fs.DirEntry, error) {
	relevantFiles := []fs.DirEntry{}
	files, err := os.ReadDir(path)

	if err != nil {
		return relevantFiles, fmt.Errorf("could not read directory: %w", err)
	}

	for _, file := range files {
		name := strings.ToLower(file.Name())

		for _, fileType := range fileTypes {
			if strings.HasSuffix(name, strings.ToLower(fileType)) {
				relevantFiles = append(relevantFiles, file)
			}
		}
	}

	return relevantFiles, nil
}
