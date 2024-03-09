package main

import (
	"errors"
	"fmt"

	"github.com/RealMotz/pokedex-repl/internal/pokeapi"
)

func inspectCommand(client pokeapi.Client, pokemonName string) error {
	if len(pokemonName) == 0 {
		return errors.New("pokemon name is required")
	}

	if poke, ok := pokedex[pokemonName]; ok {
		printPokemonInfo(poke)
		return nil
	}
	fmt.Println("You have not caught that pokemon")
	return nil
}

func printPokemonInfo(poke pokeapi.Pokemon) {
	fmt.Println("Name: " + poke.Name)
	fmt.Println("Height: " + fmt.Sprint(poke.Height))
	fmt.Println("Weight: " + fmt.Sprint(poke.Weight))
	fmt.Println("Stats:")
	for _, stat := range poke.Stats {
		fmt.Printf("  - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range poke.Types {
		fmt.Printf("  - %v\n", pokeType.Type.Name)
	}
}
