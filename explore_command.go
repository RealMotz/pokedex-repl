package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/RealMotz/pokedex-repl/internal/pokeapi"
)

var pokeArea pokeapi.PokeArea

func exploreCommand(client pokeapi.Client, area string) error {
	if len(area) == 0 {
		return errors.New("Provide a valid area name")
	}

	url := Concatenate(locationAreaUrl, area)
	if val, ok := cache.Get(url); ok {
		json.Unmarshal(val, &pokeArea)
    printPokemon(pokeArea)
		return nil
	}

	pokeArea, err := client.GetArea(url)
	if err != nil {
		return err
	}

	val, err := json.Marshal(pokeArea)
	if err != nil {
		return err
	}

	cache.Add(url, val)
  printPokemon(pokeArea)	
	return nil
}

func printPokemon(area pokeapi.PokeArea) {
	fmt.Printf("Exploring %v... \n", area.Location.Name)
	fmt.Println("Pokemon found:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Println("+ " + encounter.Pokemon.Name)
	}
}
