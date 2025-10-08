package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func GitPath() (string, error) {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	gitPath := filepath.Join(currentWorkingDirectory, ".git.clone")
	return gitPath, nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
