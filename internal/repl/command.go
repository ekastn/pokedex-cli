package repl

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbacExit,
		},
	}
}

func callbackHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()

    fmt.Println("Usage:")
    commands := getCommands()

    for _, cmd := range commands {
        fmt.Printf("- %v: %v\n", cmd.name, cmd.description)
    }

	fmt.Println()

	return nil
}

func callbacExit() error {
	os.Exit(0)
	return nil
}
