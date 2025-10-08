package files

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SamJohn04/gitclone/internal/utils"
)

func WriteHead(dirPath, branchName string) error {
	headPath := filepath.Join(dirPath, "HEAD")
	refText := fmt.Sprintf("ref: refs/heads/%s\n", branchName)
	return utils.WriteFile(headPath, refText)
}

func WriteConfig(dirPath string) error {
	configPath := filepath.Join(dirPath, "config")
	configFile, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer configFile.Close()

	fmt.Fprintln(configFile, "[core]")
	fmt.Fprintln(configFile, "\trepositoryformatversion = 0")
	fmt.Fprintln(configFile, "\tfilemode = true")
	fmt.Fprintln(configFile, "\tbare = false")
	fmt.Fprintln(configFile, "\tlogallrefupdates = true")

	return nil
}

func WriteDescription(dirPath string) error {
	descriptionPath := filepath.Join(dirPath, "description")
	return utils.WriteFile(descriptionPath,
		"Unnamed repository; edit this file 'description' to name the repository.\n")
}
