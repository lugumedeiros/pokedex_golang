package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"pokedexcligo/internal"
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
		"map" : {
			name: "map",
			description: "Display map locations",
			callback: commandMapCurrent,
		},
		"mapb" : {
			name: "mapb",
			description: "Display previous map locations",
			callback: commandMapBack,
		},
		"mapn" : {
			name: "mapn",
			description: "Display next map locations",
			callback: commandMapNext,
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
		fmt.Printf("Unknown command\n")
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

func commandMapExec(direction string) error {
	var locations []internal.Location
	switch direction {
	case "current":
		locations = internal.GetLocation()
	case "next":
		locations = internal.GetLocationNext()
	case "back":
		locations = internal.GetLocationBack()
	}
	if locations == nil {
		return errors.New("Something went wrong.\n")
	}
	for _, loc := range locations{
		fmt.Printf("%v\n", loc.Name)
	}
	return nil
}

func commandMapCurrent() error {
	return commandMapExec("current")
}

func commandMapNext() error {
	return commandMapExec("next")
}

func commandMapBack() error {
	return commandMapExec("back")
}
