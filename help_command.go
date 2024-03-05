package main

import (
	"errors"
	"fmt"

	"github.com/RealMotz/pokedex-repl/internal/pokeapi"
)

func helpCommand(client pokeapi.Client, param string) error {
  if len(param) > 0 {
    return errors.New("too many arguments")
  }
	fmt.Println("Welcome to the PokeDex!")
	fmt.Println()
	fmt.Println("Usage: ")
	fmt.Println()
	for _, cmd := range commands() {
		fmt.Println(cmd.name, ": ", cmd.description)
	}
	fmt.Println()
	return nil
}
