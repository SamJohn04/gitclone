package start

import (
	"errors"
	"fmt"
	"os"

	"github.com/SamJohn04/gitclone/internal/files"
	"github.com/SamJohn04/gitclone/internal/utils"
)

func InitCommand() error {
	gitPath, err := utils.GitPath()
	if err != nil {
		return err
	}
	if utils.PathExists(gitPath) {
		return errors.New("not implemented")
	} else {
		err = os.Mkdir(gitPath, 0o755)
		if err != nil {
			return err
		}
	}

	err = files.WritaAllBaseFiles(gitPath, "main")
	if err != nil {
		return err
	}

	_, err = files.MakeBranchDir(gitPath)
	if err != nil {
		return err
	}

	// TODO add default hooks
	_, err = files.MakeHooksDir(gitPath)
	if err != nil {
		return err
	}

	err = files.WriteInfoDirAndContents(gitPath)
	if err != nil {
		return err
	}

	err = files.MakeObjectDirAndContents(gitPath)
	if err != nil {
		return err
	}

	err = files.MakeRefDirAndConents(gitPath)
	if err != nil {
		return err
	}

	fmt.Println("An empty repository initialized in", gitPath)
	return nil
}
