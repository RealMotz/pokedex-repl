package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/RealMotz/pokedex-repl/internal/pokeapi"
	"github.com/RealMotz/pokedex-repl/internal/pokecache"
)

var location pokeapi.PokeLocation
var startingUrl string = "https://pokeapi.co/api/v2/location/?offset=0&limit=20"
var next *string = &startingUrl
var prev *string = nil
var cache pokecache.PokeCache = pokecache.NewCache(time.Second * 86400)

func printMaps(locations pokeapi.PokeLocation) {
	fmt.Println("============================")
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println("============================")
}

func mapCommand(client pokeapi.Client) error {
	if next == nil {
		next = &startingUrl
	}
	if val, ok := cache.Get(*next); ok {
		json.Unmarshal(val, &location)
		next = location.Next
		prev = location.Previous
		printMaps(location)
		return nil
	}

	location = client.GetLocation(*next)
	val, err := json.Marshal(location)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cache.Add(*next, val)
	prev = location.Previous 
	next = location.Next
	printMaps(location)
	return nil
}

func mapBackCommand(client pokeapi.Client) error {
	if prev == nil {
    next = &startingUrl
		fmt.Println("Nothing to display")
		return nil
	}
	val, ok := cache.Get(*prev)
	if !ok {
		fmt.Println("No records founds in cache")
		return nil
	}

	json.Unmarshal(val, &location)
	next = location.Next
	prev = location.Previous
	printMaps(location)
	return nil
}
