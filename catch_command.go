package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/RealMotz/pokedex-repl/internal/pokeapi"
)

var pokemon pokeapi.Pokemon
var pokemonUrl string = "https://pokeapi.co/api/v2/pokemon/"
var pokedex map[string]pokeapi.Pokemon = make(map[string]pokeapi.Pokemon)
var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func catchCommand(client pokeapi.Client, pokemonName string) error {
	if len(pokemonName) == 0 {
		return errors.New("pokemon name is required")
	}

	url := Concatenate(pokemonUrl, pokemonName)
	pokemon, err := client.GetPokemon(url)
	if err != nil {
		return err
	}

	catchChance := pokemon.BaseExperience / 2
	boost := []int{1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 4}
	catchAttempt := boost[rand.Intn(len(boost))] * catchChance
	roll := r.Intn(pokemon.BaseExperience)
  fmt.Println(catchAttempt)
  fmt.Println(roll)

	catched := catchAttempt >= roll
	fmt.Println("Throwing a Pokeball at " + pokemon.Name)
	if !catched {
		fmt.Println(pokemon.Name + " escaped!")
    return nil
	}

	fmt.Println(pokemon.Name + " was caught!")
	pokedex[pokemonName] = pokemon
	return nil
}
