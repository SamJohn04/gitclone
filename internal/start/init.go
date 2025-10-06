package start

import (
	"errors"
	"os"
	"path/filepath"
)

func InitCommand() error {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		return err
	}

	gitPath := filepath.Join(currentWorkingDirectory, ".git.clone")
	if _, err := os.Stat(gitPath); !errors.Is(err, os.ErrNotExist) {
		return errors.New("already exists")
	}

	err = os.Mkdir(gitPath, 0o755)
	if err != nil {
		return err
	}

	return nil
}
