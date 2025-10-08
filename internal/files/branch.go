package files

import (
	"os"
	"path/filepath"
)

func MakeBranchDir(dirPath string) error {
	branchPath := filepath.Join(dirPath, "branches")
	return os.Mkdir(branchPath, 0o755)
}
