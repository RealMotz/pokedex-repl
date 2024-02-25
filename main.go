package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(map[string]cliCommand) error
}

func main() {
	cmds := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: func(cmds map[string]cliCommand) error {
				fmt.Println("Welcome to the PokeDex!")
				fmt.Println("")
				fmt.Println("Usage: ")
				fmt.Println("")
				for _, cmd := range cmds {
					fmt.Println(cmd.name, ": ", cmd.description)
				}
				fmt.Println("")
				return nil
			},
		},
		"exit": {
			name:        "exit",
			description: "Exit the application",
			callback: func(cmds map[string]cliCommand) error {
				return errors.New("EOF")
			},
		},
	}

	fmt.Println("Welcome to the PokeDex!")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Pokedex > ")
	for scanner.Scan() {
		userInput := scanner.Text()
		if _, ok := cmds[userInput]; !ok {
			fmt.Println("command not found")
			fmt.Printf("Pokedex > ")
			continue
		}
		ok := cmds[userInput].callback(cmds)
		if ok != nil {
			return
		}
		fmt.Printf("Pokedex > ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
