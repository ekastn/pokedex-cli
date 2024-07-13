package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(* config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "lists some location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "lists some location areas",
			callback:    callbackMapb,
		},
	}
}

func callbackHelp(cfg *config) error {
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

func callbackExit(cfg *config) error {
	os.Exit(0)
	return nil
}

func callbackMap(cfg *config) error {
	res, err := cfg.PokeapiClinet.ListLocationAreas(cfg.NextLocationAreasUrl)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Println(area.Name)
	}

	cfg.NextLocationAreasUrl = res.Next
	cfg.PrevLocationAreasUrl = res.Previous

	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.PrevLocationAreasUrl == nil {
		return errors.New("You are on the first page")
	}

	res, err := cfg.PokeapiClinet.ListLocationAreas(cfg.PrevLocationAreasUrl)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Println(area.Name)
	}

	cfg.NextLocationAreasUrl = res.Next
	cfg.PrevLocationAreasUrl = res.Previous

	return nil
}
