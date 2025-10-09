package files

import (
	"os"
	"path/filepath"
)

func MakeHooksDir(dirPath string) (string, error) {
	hooksPath := filepath.Join(dirPath, "hooks")
	return hooksPath, os.Mkdir(hooksPath, 0o755)
}
