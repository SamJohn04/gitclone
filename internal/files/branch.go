package files

import (
	"os"
	"path/filepath"

	"github.com/SamJohn04/gitclone/internal/constants"
)

func MakeBranchDir(dirPath string) (string, error) {
	branchPath := filepath.Join(dirPath, constants.BranchDirName)
	return branchPath, os.Mkdir(branchPath, 0o755)
}
