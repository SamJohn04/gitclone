package files

import (
	"os"
	"path/filepath"
)

func MakeRefDir(dirPath string) error {
	refPath := filepath.Join(dirPath, "refs")
	return os.Mkdir(refPath, 0o755)
}

func MakeHeadsInRefDir(dirPath string) error {
	headsPath := filepath.Join(dirPath, "refs", "heads")
	return os.Mkdir(headsPath, 0o755)
}

func MakeTagsInRefDir(dirPath string) error {
	tagsPath := filepath.Join(dirPath, "refs", "tags")
	return os.Mkdir(tagsPath, 0o755)
}
