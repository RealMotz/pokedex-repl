package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/RealMotz/pokedex-repl/internal/pokeapi"
	"github.com/RealMotz/pokedex-repl/internal/pokecache"
)

var location pokeapi.PokeLocation
var locationAreaUrl string = "https://pokeapi.co/api/v2/location-area/"
var pagination string = "?offset=0&limit=20"
var next *string = nil
var prev *string = nil
var cache pokecache.PokeCache = pokecache.NewCache(time.Second * 86400)

func printMaps(locations pokeapi.PokeLocation) {
	fmt.Println("============================")
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println("============================")
}

func Concatenate(params ...string) string {
	url := ""
	for i := 0; i < len(params); i++ {
		url += params[i]
	}

	return url
}

func mapCommand(client pokeapi.Client, param string) error {
	if len(param) > 0 {
		return errors.New("too many arguments")
	}

	if next == nil {
		url := Concatenate(locationAreaUrl, pagination)
		next = &url
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

func mapBackCommand(client pokeapi.Client, param string) error {
	if len(param) > 0 {
		return errors.New("too many arguments")
	}

	if prev == nil {
		url := Concatenate(locationAreaUrl, pagination)
		next = &url
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
