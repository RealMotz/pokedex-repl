package main

import (
	"errors"
	"fmt"

	"github.com/RealMotz/pokedex-repl/internal/pokeapi"
)

func pokedexCommand(client pokeapi.Client, param string) error {
  if len(param) > 0 {
    return errors.New("too many arguments")
  }
  if len(pokedex) == 0 {
    fmt.Println("Your pokedex is empty...")
    return nil
  }
  fmt.Println("Your Pokedex:")
  for _, poke := range pokedex {
    println("- " + poke.Name)
  }
	return nil
}
