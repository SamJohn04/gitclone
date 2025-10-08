package start

import (
	"fmt"
	"os"

	"github.com/SamJohn04/gitclone/internal/files"
	"github.com/SamJohn04/gitclone/internal/utils"
)

func InitCommand() {
	gitPath, err := utils.GitPath()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if utils.PathExists(gitPath) {
		fmt.Println("Reinitialization not yet implemented")
		return
	} else {
		err = os.Mkdir(gitPath, 0o755)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	// !ignore return err
	files.WriteHead(gitPath, "main")
	files.WriteConfig(gitPath)
	files.WriteDescription(gitPath)

	// !ignore return err
	files.MakeBranchDir(gitPath)

	// !ignore return err
	// TODO add default hooks
	files.MakeHooksDir(gitPath)

	// !ignore return err
	files.MakeInfoDir(gitPath)
	files.WriteExcludeInInfo(gitPath)

	// !ignore return err
	files.MakeObjectDir(gitPath)
	files.MakeInfoInObjectDir(gitPath)
	files.MakePackInObjectDir(gitPath)

	// !ignore return err
	files.MakeRefDir(gitPath)
	files.MakeHeadsInRefDir(gitPath)
	files.MakeTagsInRefDir(gitPath)

	fmt.Println("An empty repository initialized in", gitPath)
}
