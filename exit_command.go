package main

import (
	"errors"
	"os"

	"github.com/RealMotz/pokedex-repl/internal/pokeapi"
)

func exitCommand(client pokeapi.Client, param string) error {
  if len(param) > 0 {
    return errors.New("too many arguments")
  }
  defer os.Exit(0)
  return nil
}
