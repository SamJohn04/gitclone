package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func MakeInfoDir(dirPath string) error {
	infoPath := filepath.Join(dirPath, "info")
	return os.Mkdir(infoPath, 0o755)
}

func WriteExcludeInInfo(dirPath string) error {
	excludePath := filepath.Join(dirPath, "info", "exclude")
	excludeFile, err := os.OpenFile(excludePath, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer excludeFile.Close()

	fmt.Fprintln(excludeFile, "# git ls-files --others --exclude-from=.git/info/exclude")
	fmt.Fprintln(excludeFile, "# Lines that start with '#' are comments.")
	fmt.Fprintln(excludeFile, "# For a project mostly in C, the following would be a good set of")
	fmt.Fprintln(excludeFile, "# exclude patterns (uncomment them if you want to use them):")
	fmt.Fprintln(excludeFile, "# *.[oa]")
	fmt.Fprintln(excludeFile, "# *~")

	return nil
}
