package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteInfoDirAndContents(dirPath string) error {
	infoPath, err := MakeInfoDir(dirPath)
	if err != nil {
		return err
	}
	return WriteExcludeInInfo(infoPath)
}

func MakeInfoDir(dirPath string) (string, error) {
	infoPath := filepath.Join(dirPath, "info")
	return infoPath, os.Mkdir(infoPath, 0o755)
}

func WriteExcludeInInfo(infoPath string) error {
	excludePath := filepath.Join(infoPath, "exclude")
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
