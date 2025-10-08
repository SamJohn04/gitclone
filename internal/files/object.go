package files

import (
	"os"
	"path/filepath"
)

func MakeObjectDir(dirPath string) error {
	objectPath := filepath.Join(dirPath, "objects")
	return os.Mkdir(objectPath, 0o755)
}

func MakeInfoInObjectDir(dirPath string) error {
	infoPath := filepath.Join(dirPath, "objects", "info")
	return os.Mkdir(infoPath, 0o755)
}

func MakePackInObjectDir(dirPath string) error {
	packPath := filepath.Join(dirPath, "objects", "pack")
	return os.Mkdir(packPath, 0o755)
}
