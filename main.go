package main

import (
	"time"

	"github.com/RealMotz/pokedex-repl/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	startRepl(client)
}
