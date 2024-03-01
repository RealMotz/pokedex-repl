package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokeLocation struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

func getLocation(url string) PokeLocation {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}

	location := PokeLocation{}
	decodingErr := json.Unmarshal(body, &location)
	if decodingErr != nil {
		fmt.Println(err)
	}
	return location
}

var location PokeLocation

func printMaps(locations PokeLocation) {
  fmt.Println("============================")
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
  fmt.Println("============================")
}

func mapCommand() error {
	if location.Next == nil {
		location = getLocation("https://pokeapi.co/api/v2/location/")
	} else {
		location = getLocation(*location.Next)
	}
  printMaps(location)

	return nil
}

func mapBackCommand() error {
	if location.Previous == nil {
    fmt.Println("Nothing to display")
    return nil
	}
	location = getLocation(*location.Previous)
  printMaps(location)

	return nil
}
