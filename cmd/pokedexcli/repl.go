package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, ok := commands[cleaned[0]]

		if !ok {
			fmt.Println("Invalid command. Type 'help' to see the available commands.")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}
