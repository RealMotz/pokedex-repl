package main

import (
	"errors"
	"github.com/RealMotz/pokedex-repl/internal/pokeapi"
)

func exitCommand(client pokeapi.Client) error {
	return errors.New("EOF")
}
