package files

import (
	"os"
	"path/filepath"

	"github.com/SamJohn04/gitclone/internal/constants"
)

func MakeHooksDir(dirPath string) (string, error) {
	hooksPath := filepath.Join(dirPath, constants.HookDirName)
	return hooksPath, os.Mkdir(hooksPath, 0o755)
}
