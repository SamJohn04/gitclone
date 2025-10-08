/*
Gitclone clones the command git.
*/
package main

import (
	"fmt"
	"os"

	"github.com/SamJohn04/gitclone/internal/start"
)

func main() {
	if len(os.Args) == 1 {
		showHelp()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "init":
		start.InitCommand()
	}
}

func showHelp() {
	fmt.Println("usage: gitclone <command> [<args>]")
}
