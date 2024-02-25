package main

import (
	"fmt"
)

func helpCommand() error {
	fmt.Println("Welcome to the PokeDex!")
	fmt.Println("")
	fmt.Println("Usage: ")
	fmt.Println("")
	for _, cmd := range commands() {
		fmt.Println(cmd.name, ": ", cmd.description)
	}
	fmt.Println("")
	return nil
}
