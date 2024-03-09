package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RealMotz/pokedex-repl/internal/pokeapi"
)

func startRepl(client pokeapi.Client) {
	fmt.Println("Welcome to the PokeDex!")
	fmt.Printf("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
	cmds := commands()

	for scanner.Scan() {
		userInput := strings.Split(scanner.Text(), " ")
		command := clean(userInput[0])
		args := ""
		if len(userInput) == 2 {
			args = clean(userInput[1])
		}
		if len(userInput) > 2 {
			fmt.Println("too many arguments")
			fmt.Printf("Pokedex > ")
			continue
		}

		if _, ok := cmds[command]; !ok {
			fmt.Println("not a valid command")
			fmt.Println("type 'help' to see all valid commands")
			fmt.Printf("Pokedex > ")
			continue
		}
		err := cmds[command].callback(client, args)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Pokedex > ")
			continue
		}
		fmt.Printf("Pokedex > ")
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(pokeapi.Client, string) error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations",
			callback:    mapCommand,
		},
		"bmap": {
			name:        "bmap",
			description: "Displays the previous 20 locations",
			callback:    mapBackCommand,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "Shows the pokemon encounters in a given area",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attempts to capture a pokemon",
			callback:    catchCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows all captured pokemon",
			callback:    pokedexCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the application",
			callback:    exitCommand,
		},
	}
}

func clean(input string) string {
	return strings.Trim(strings.ToLower(input), " ")
}
