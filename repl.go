package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	fmt.Println("Welcome to the PokeDex!")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Pokedex > ")
	cmds := commands()

	for scanner.Scan() {
		userInput := clean(scanner.Text())
		if _, ok := cmds[userInput]; !ok {
			fmt.Println("command not found")
			fmt.Printf("Pokedex > ")
			continue
		}
		ok := cmds[userInput].callback()
		if ok != nil {
			return
		}
		fmt.Printf("Pokedex > ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
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
