package files

import (
	"os"
	"path/filepath"
)

func MakeBranchDir(dirPath string) (string, error) {
	branchPath := filepath.Join(dirPath, "branches")
	return branchPath, os.Mkdir(branchPath, 0o755)
}
