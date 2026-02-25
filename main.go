package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;;{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned_text := cleanInput(text)
		if len(cleaned_text) == 0 {
			continue
		}
		fmt.Printf("Your command was: %v\n", cleaned_text[0])
	}
}