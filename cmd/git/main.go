/*
Gitclone clones the command git.
*/
package main

import (
	"fmt"
	"os"

	"github.com/SamJohn04/gitclone/internal/cli"
	"github.com/SamJohn04/gitclone/internal/start"
)

func main() {
	userInput := cli.ParseCommand()
	if userInput.ShowHelp {
		showHelp()
		os.Exit(1)
	}
	if userInput.Unknown {
		os.Exit(1)
	}

	switch userInput.Action {
	case "init":
		// TODO implement variety and verbosity
		if start.InitCommand() {
			os.Exit(1)
		}
	}
}

func showHelp() {
	fmt.Println("usage: gitclone <command> [<args>]")
	fmt.Println("commands currently implemented:")
	fmt.Println("\tinit")
}
