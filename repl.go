package main

import (
	"errors"
	"fmt"
	"os"
	"pokedexcligo/internal"
	"strings"
)

// import "regexp"

type t_cliCommand struct {
	name        string
	description string
	callback    func(args []string) error
	args_n      int
}

var cliCommandsMap map[string]t_cliCommand

func init() {
	cliCommandsMap = map[string]t_cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display map locations",
			callback:    commandMapCurrent,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous map locations",
			callback:    commandMapBack,
		},
		"mapn": {
			name:        "mapn",
			description: "Display next map locations",
			callback:    commandMapNext,
		},
		"explore": {
			name:        "explore",
			description: "Display all pokemons available in a location",
			callback:    commandPokemonArea,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect pokemons added to pokedex",
			callback:    commandInspect,
		},
		"pokedex" : {
			name: "pokedex",
			description: "Display all pokemons added to pokedex",
			callback: commandPokedex,
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

func CommandExec(command string, args []string) error {
	cliCommand, ok := cliCommandsMap[command]
	if ok {
		cliCommand.callback(args)
	} else {
		fmt.Printf("Unknown command\n")
	}
	return nil
}

func commandExit(args []string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(args []string) error {
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
	for _, loc := range locations {
		fmt.Printf("%v\n", loc.Name)
	}
	return nil
}

func commandMapCurrent(args []string) error {
	return commandMapExec("current")
}

func commandMapNext(args []string) error {
	return commandMapExec("next")
}

func commandMapBack(args []string) error {
	return commandMapExec("back")
}

func commandPokemonArea(args []string) error {
	if len(args) < 1 {
		return errors.New("missing Location Area")
	}
	area := args[0]
	pokemons := internal.GetPokemonInArea(area)
	if len(pokemons) == 0 {
		fmt.Printf("No pokemon found in area or area doesn't exist.\n")
	} else {
		fmt.Printf("Found Pokemon:\n")
		for _, name := range pokemons {
			fmt.Printf("- %v\n", name)
		}
	}
	return nil
}

func commandCatch(args []string) error {
	if len(args) < 1 {
		return errors.New("Missing Pokemon name.")
	}
	pokemon_name := strings.ToLower(args[0])
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon_name)
	success, err := internal.GetPokemon(pokemon_name)
	if err {
		return errors.New("Something went wrong.")
	} else {
		if success {
			fmt.Printf("%v was caught!\n", pokemon_name)
		} else {
			fmt.Printf("%v escaped!\n", pokemon_name)
		}
	}
	return nil
}

func commandInspect(args []string) error {
	if len(args) < 1 {
		return errors.New("Missing Pokemon name.")
	}
	if stat, ok := internal.GetPokemonFromPokedex(args[0]); ok {
		// print stats
		fmt.Printf("Name: %v\n", stat.Name)
		fmt.Printf("Height: %v\n", stat.Height)
		fmt.Printf("Weight: %v\n", stat.Weight)
		fmt.Printf("Stats:\n")
		for _, unique_stat := range stat.Stats {
			fmt.Printf("  -%v: %v\n", unique_stat.Stat.Name, unique_stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, types := range stat.Types {
			fmt.Printf("  - %v\n", types.Type.Name)
		}
		return nil
	}
	return errors.New("you have not caught that pokemon")
}

func commandPokedex(args []string) error {
	pokemons := internal.GetCaughPokemon()
	if len(pokemons) == 0 {
		fmt.Printf("Your Pokedex is empty!\n")
	} else {
		fmt.Printf("Your Pokedex:\n")
		for _, name := range pokemons {
			fmt.Printf(" - %v\n", name)
		}
	}
	return nil
}
