package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to the Pokedex!\n")
	for ;;{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned_text := cleanInput(text)
		if len(cleaned_text) == 0 {
			continue
		}
		CommandExec(text)
		fmt.Print("\n")
	}
}