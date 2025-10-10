package cli

import "os"

type Input struct {
	Action string
	Paths  []string

	Write bool

	ShowHelp bool
	Unknown  bool
}

func ParseCommand() Input {
	userInput := Input{
		Action: "",
		Paths:  []string{},

		Write: false,

		ShowHelp: false,
		Unknown:  false,
	}
	args := os.Args[1:]

	if len(args) == 0 {
		userInput.ShowHelp = true
		return userInput
	}

	switch args[0] {
	case "init":
		userInput.Action = "init"
		return parseInit(args[1:], userInput)
	case "hash-object":
		userInput.Action = "hash-object"
		return parseHashObject(args[1:], userInput)
	default:
		userInput.Unknown = true
	}

	return userInput
}

func parseInit(args []string, userInput Input) Input {
	if len(args) == 0 {
		return userInput
	}

	// TODO make verbose
	if args[0][0] != '-' {
		userInput.Paths = append(userInput.Paths, args[0])
	}

	return userInput
}

func parseHashObject(args []string, userInput Input) Input {
	if len(args) == 0 {
		return userInput
	}

	// TODO make more verbose
	for _, arg := range args {
		if arg == "-w" {
			userInput.Write = true
		} else if arg[0] == '-' {
			userInput.Unknown = true
			return userInput
		} else {
			userInput.Paths = append(userInput.Paths, arg)
		}
	}

	return userInput
}
