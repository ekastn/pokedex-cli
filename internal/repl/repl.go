package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func SartRepl() {
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

		command, ok := commands[cleaned[0]]

		if !ok {
			fmt.Println("Invalid command. Type 'help' to see the available commands.")
			continue
		}

		err := command.callback()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
