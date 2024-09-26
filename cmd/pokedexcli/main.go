package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ekastn/pokedex-cli/internal/pokeapi"
)

type config struct {
	PokeapiClinet        pokeapi.Client
	PrevLocationAreasUrl *string
	NextLocationAreasUrl *string
	caughtPokemon        map[string]pokeapi.Pokemon
}

func main() {
	dat, err := os.ReadFile("assets/pokedex.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dat))

	startRepl(&config{
		PokeapiClinet: pokeapi.NewClient(),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	})
}
