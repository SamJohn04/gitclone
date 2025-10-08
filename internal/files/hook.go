package files

import (
	"os"
	"path/filepath"
)

func MakeHooksDir(dirPath string) error {
	hooksPath := filepath.Join(dirPath, "hooks")
	return os.Mkdir(hooksPath, 0o755)
}
