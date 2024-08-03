package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
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
		"explore": {
			name:        "explore (location areas)",
			description: "Find a pokemon in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch (pokemon name)",
			description: "Attempt to catch a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect (pokemon name)",
			description: "Inspect a caught pokemon",
			callback:    callbackInspect,
		},
	}
}

func callbackHelp(cfg *config, args ...string) error {
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

func callbackExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}

func callbackMap(cfg *config, args ...string) error {
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

func callbackMapb(cfg *config, args ...string) error {
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

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a location area name")
	}

	locationArea, err := cfg.PokeapiClinet.GetLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, area := range locationArea.PokemonEncounters {
		fmt.Println(area.Pokemon.Name)
	}

	return nil
}

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a pokemon name to catch")
	}

	pokemon, err := cfg.PokeapiClinet.GetPokemon(args[0])
	if err != nil {
		return err
	}

	threshhold := 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	if randNum > threshhold {
		fmt.Printf("%v excaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%v was caught\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	}

	return nil
}

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("You must provide a pokemon name to inspect")
	}

	pokemon, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("- %v\n", t.Type.Name)
	}

	return nil
}
