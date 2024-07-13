package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eka-septian/pokedex-cli/internal/repl"
)

func main() {
	dat, err := os.ReadFile("assets/pokedex.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dat))

	repl.SartRepl()
}
