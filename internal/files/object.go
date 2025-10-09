package files

import (
	"os"
	"path/filepath"
)

func MakeObjectDirAndContents(dirPath string) error {
	objectPath, err := MakeObjectDir(dirPath)
	if err != nil {
		return err
	}
	err = MakeInfoInObjectDir(objectPath)
	if err != nil {
		return err
	}
	return MakePackInObjectDir(objectPath)
}

func MakeObjectDir(dirPath string) (string, error) {
	objectPath := filepath.Join(dirPath, "objects")
	return objectPath, os.Mkdir(objectPath, 0o755)
}

func MakeInfoInObjectDir(objectPath string) error {
	infoPath := filepath.Join(objectPath, "info")
	return os.Mkdir(infoPath, 0o755)
}

func MakePackInObjectDir(objectPath string) error {
	packPath := filepath.Join(objectPath, "pack")
	return os.Mkdir(packPath, 0o755)
}
