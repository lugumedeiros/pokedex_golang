package main

import (
	"fmt"
	"os"
	"strings"
)

// import "regexp"

type t_cliCommand struct {
	name string
	description string
	callback func() error
}
var cliCommandsMap map[string]t_cliCommand

func init() {
	cliCommandsMap = map[string]t_cliCommand{
		"exit" : {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help" : {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
}

func cleanInput(text string) []string {
	str_slice := strings.Fields(text)
	var new_slice []string

	for i := range str_slice {
		lower := strings.ToLower(str_slice[i])
		new_slice = append(new_slice, lower)
	}
	return new_slice
}

func CommandExec(command string) error {
	cliCommand, ok := cliCommandsMap[command]
	if ok {
		cliCommand.callback()
	} else {
		fmt.Printf("Unknown command")
	}
	return nil
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("Usage:\n\n")
	for _, command_t := range cliCommandsMap {
		fmt.Printf("%v: %v\n", command_t.name, command_t.description)
	}
	return nil
}
