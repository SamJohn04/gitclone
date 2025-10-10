package files

import (
	"os"
	"path/filepath"

	"github.com/SamJohn04/gitclone/internal/constants"
)

func MakeRefDirAndConents(dirPath string) error {
	refPath, err := MakeRefDir(dirPath)
	if err != nil {
		return err
	}
	err = MakeHeadsInRefDir(refPath)
	if err != nil {
		return err
	}
	return MakeTagsInRefDir(refPath)
}

func MakeRefDir(dirPath string) (string, error) {
	refPath := filepath.Join(dirPath, constants.RefDirName)
	return refPath, os.Mkdir(refPath, 0o755)
}

func MakeHeadsInRefDir(refPath string) error {
	headsPath := filepath.Join(refPath, constants.HeadsInRefName)
	return os.Mkdir(headsPath, 0o755)
}

func MakeTagsInRefDir(refPath string) error {
	tagsPath := filepath.Join(refPath, constants.TagsInRefName)
	return os.Mkdir(tagsPath, 0o755)
}
